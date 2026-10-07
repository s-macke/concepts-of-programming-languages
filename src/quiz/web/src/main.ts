import { hostGroup, showGroup } from "./group";
import { api, ApiError } from "./api";
import { el, renderInputs, renderPrompt } from "./render";
import type { SessionInfo } from "./types";
import "./style.css";

const app = document.querySelector<HTMLDivElement>("#app")!;
let routeGeneration = 0;

/** Replaces the whole view. The app is small enough not to need diffing. */
function show(...nodes: (Node | string)[]): void {
  app.replaceChildren(...nodes);
}

function showError(message: string): void {
  const banner = el("div", { className: "error", textContent: message });
  app.prepend(banner);
}

/** Start page: the quizzes on the server plus an upload field for testing. */
async function showStartPage(): Promise<void> {
  const generation = routeGeneration;
  const hash = location.hash;
  // A response from an earlier navigation must not replace the current screen.
  const isCurrent = () => generation === routeGeneration && hash === location.hash;
  show(el("p", { className: "loading", textContent: "Loading quizzes…" }));

  let quizzes;
  try {
    quizzes = await api.listQuizzes();
  } catch (err) {
    if (!isCurrent()) return;
    show(el("h1", { textContent: "Quiz" }));
    showError(`Cannot load quizzes: ${(err as Error).message}`);
    return;
  }

  if (!isCurrent()) return;

  const list = el("ul", { className: "quiz-list" });
  for (const quiz of quizzes) {
    const button = el("button", { className: "quiz-choice", type: "button" },
      el("span", { className: "quiz-title", textContent: quiz.title }),
      el("span", { className: "quiz-meta", textContent: `${quiz.numQuestions} questions` }));
    button.addEventListener("click", () => void start(() => api.startSession(quiz.id)));

    const host = el("button", { type: "button", className: "primary", textContent: "Host group quiz" });
    host.onclick = async () => { host.disabled = true; try { await hostGroup(quiz.id); } catch (err) { showError((err as Error).message); } finally { host.disabled = false; } };
    const item = el("li", {}, button, host);
    if (quiz.description) item.append(el("p", { className: "quiz-description", textContent: quiz.description }));
    list.append(item);
  }
  if (quizzes.length === 0) {
    list.append(el("li", { className: "empty", textContent: "No quizzes on the server yet." }));
  }

  const groupUpload = el("input", { type: "checkbox" });
  const uploadMode = el("label", {}, groupUpload, " Host uploaded quiz as a group");
  const file = el("input", { type: "file", accept: ".yaml,.yml", id: "upload" });
  file.addEventListener("change", () => {
    const chosen = file.files?.[0];
    if (chosen) {
      if (groupUpload.checked) void hostGroup(chosen).catch(err => showError((err as Error).message));
      else void start(() => api.upload(chosen));
    }
  });

  show(
    el("h1", { textContent: "Quiz" }),
    list,
    el("section", { className: "upload" },
      el("h2", { textContent: "Test your own quiz" }),
      el("p", { className: "hint", textContent: "Upload a YAML quiz file. It stays in memory and is never stored on the server." }),
      uploadMode, file),
  );
}

/** Runs a session starter and moves on to the first question. */
async function start(begin: () => Promise<SessionInfo>): Promise<void> {
  show(el("p", { className: "loading", textContent: "Starting…" }));
  try {
    const info = await begin();
    await showQuestion(info, 1);
  } catch (err) {
    await showStartPage();
    const message = err instanceof ApiError ? err.message : (err as Error).message;
    showError(`Cannot start the quiz: ${message}`);
  }
}

async function showQuestion(info: SessionInfo, n: number): Promise<void> {
  const question = await api.question(info.sessionId, n);
  const inputs = renderInputs(question);

  const feedback = el("div", { className: "feedback" });
  const submit = el("button", { className: "primary", type: "button", textContent: "Check answer" });
  const next = el("button", { className: "primary", type: "button", hidden: true,
    textContent: n < info.numQuestions ? "Next question" : "See result" });

  submit.addEventListener("click", async () => {
    const given = inputs.read();
    if (given.length === 0) {
      feedback.replaceChildren(el("p", { className: "hint", textContent: "Please give an answer first." }));
      return;
    }
    submit.disabled = true;
    try {
      const result = await api.answer(info.sessionId, n, given);
      inputs.disable();
      inputs.mark(result.correctAnswer);
      submit.hidden = true;
      next.hidden = false;

      const verdict = el("p", {
        className: result.correct ? "verdict correct" : "verdict wrong",
        textContent: result.correct ? `Correct (+${result.points})` : "Not correct",
      });
      feedback.replaceChildren(verdict);
      if (!result.correct && question.type === "text") {
        feedback.append(el("p", { className: "hint", textContent: `Expected: ${result.correctAnswer.join(" / ")}` }));
      }
      if (result.explanation) {
        feedback.append(el("p", { className: "explanation", textContent: result.explanation }));
      }
      next.focus();
    } catch (err) {
      submit.disabled = false;
      feedback.replaceChildren(el("p", { className: "error", textContent: (err as Error).message }));
    }
  });

  next.addEventListener("click", () => {
    if (n < info.numQuestions) void showQuestion(info, n + 1);
    else void showResult(info);
  });

  show(
    el("header", { className: "quiz-header" },
      el("h1", { textContent: info.title }),
      el("p", { className: "progress", textContent: `Question ${n} of ${info.numQuestions}` })),
    el("article", { className: "question" }, renderPrompt(question), inputs.node, feedback,
      el("div", { className: "actions" }, submit, next)),
  );
}

async function showResult(info: SessionInfo): Promise<void> {
  const summary = await api.result(info.sessionId);
  const again = el("button", { className: "primary", type: "button", textContent: "Back to the quiz list" });
  again.addEventListener("click", () => void showStartPage());

  show(
    el("header", { className: "quiz-header" }, el("h1", { textContent: info.title })),
    el("section", { className: "result" },
      el("p", { className: "score", textContent: `${summary.points} of ${summary.totalPoints} points` }),
      el("p", { textContent: `${summary.correct} of ${summary.numQuestions} questions answered correctly.` }),
      el("div", { className: "actions" }, again)),
  );
}

let stopGroup: (() => void) | undefined;
function route(): void {
  routeGeneration++;
  stopGroup?.(); stopGroup = undefined;
  const match = location.hash.match(/^#\/group\/([a-f0-9]{32})$/);
  if (match) stopGroup = showGroup(app, match[1]);
  else void showStartPage();
}
window.addEventListener("hashchange", route);
route();
