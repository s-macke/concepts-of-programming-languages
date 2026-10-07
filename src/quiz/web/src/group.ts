import QRCode from "qrcode";
import { el, renderInputs, renderPrompt } from "./render";
import type { Question, AnswerResult } from "./types";

interface Stats { number: number; text: string; submitted: number; unanswered: number; correct: number; incorrect: number; selections?: Record<string, number> }
interface State {
  title: string; phase: string; revision: number; number: number; total: number;
  participants: number; submitted: number; answered: boolean; question?: Question;
  correctAnswer?: string[]; explanation?: string; personal?: AnswerResult; stats?: Stats;
  summary?: { active: number; completed: number; submissions: number; average: number; questions: Stats[] };
}
interface Credential { token: string; host: boolean; base?: string }
class HTTPError extends Error { constructor(readonly status: number, message: string) { super(message); } }
async function request<T>(path: string, token = "", body?: unknown): Promise<T> {
  const response = await fetch(`/api/groups${path}`, {
    method: body === undefined ? "GET" : "POST",
    headers: { ...(token ? { Authorization: `Bearer ${token}` } : {}), ...(body instanceof FormData || body === undefined ? {} : { "Content-Type": "application/json" }) },
    body: body === undefined ? undefined : body instanceof FormData ? body : JSON.stringify(body),
    cache: "no-store",
  });
  const result = await response.json();
  if (!response.ok) throw new HTTPError(response.status, result.error ?? "Request failed");
  return result as T;
}
const key = (id: string) => `quiz-group:${id}`;
export async function hostGroup(quiz: string | File): Promise<void> {
  let body: unknown = { quizId: quiz };
  if (quiz instanceof File) { const form = new FormData(); form.append("file", quiz); body = form; }
  const room = await request<{ id: string; token: string; publicURL: string }>(quiz instanceof File ? "/upload" : "", "", body);
  sessionStorage.setItem(key(room.id), JSON.stringify({ token: room.token, host: true, base: room.publicURL }));
  location.hash = `/group/${room.id}`;
}

export function showGroup(app: HTMLElement, id: string): () => void {
  let stopped = false;
  let timer: ReturnType<typeof setTimeout> | undefined;
  let credential: Credential | undefined;
  let joining = false;
  try { credential = JSON.parse(sessionStorage.getItem(key(id)) ?? "null") ?? undefined; } catch { /* A damaged token can be replaced by joining again. */ }
  let state: State;
  let lastView = "";
  let progress: HTMLElement;
  const connection = el("p", { className: "hint", role: "status" });
  const content = el("div");
  app.replaceChildren(content, connection);
  const error = (err: unknown) => { connection.textContent = (err as Error).message; };
  const button = (label: string, work: () => Promise<void>) => {
    const b = el("button", { type: "button", className: "primary", textContent: label });
    b.onclick = async () => { b.disabled = true; try { await work(); } catch (err) { error(err); } finally { b.disabled = false; } };
    return b;
  };
  const home = () => { location.hash = ""; };
  function terminal(message: string) {
    stopped = true; clearTimeout(timer); sessionStorage.removeItem(key(id));
    content.replaceChildren(el("h1", { textContent: "Group quiz" }), el("p", { textContent: message }), button("Back to quizzes", async () => home()));
  }
  async function action(name: string) {
    await request(`/${id}/actions`, credential!.token, { action: name, revision: state.revision });
    if (name === "delete") { terminal("This room and its results have been deleted."); return; }
    await refresh();
  }
  function statistics(s: Stats, q?: Question): HTMLElement {
    const section = el("section", { className: "group-stats" },
      el("h2", { textContent: "Group results" }),
      el("p", { textContent: `${s.submitted} submitted · ${s.unanswered} unanswered · ${s.correct} correct · ${s.incorrect} incorrect` }));
    for (const option of q?.options ?? []) {
      const count = s.selections?.[option.id] ?? 0;
      const percent = s.submitted ? 100 * count / s.submitted : 0;
      const bar = el("progress", { max: 100, value: percent });
      bar.setAttribute("aria-label", option.text || option.id);
      section.append(el("div", { className: "stat-row" }, el("span", { textContent: `${option.text || option.id}: ${count} (${percent.toFixed(0)}%)` }), bar));
    }
    if (q?.type === "multiple") section.append(el("p", { className: "hint", textContent: "Percentages are of submissions; multiple selections can total more than 100%." }));
    return section;
  }
  function progressText(): string {
    return state.phase === "lobby" || state.phase === "finished"
      ? `${state.participants} participants`
      : `${state.participants} participants · ${state.submitted} submitted`;
  }
  function render() {
    // Participant counts and submission totals update without replacing answer inputs.
    if (progress) progress.textContent = progressText();
    const view = `${state.revision}:${state.answered}:${state.phase === "reveal" ? state.participants : ""}`;
    if (lastView === view) return;
    lastView = view;
    progress = el("p", { className: "progress", textContent: progressText() });
    const header = el("header", { className: "quiz-header" }, el("h1", { textContent: state.title }),
      el("p", { className: "hint", textContent: credential!.host ? "Host view · Anonymous group quiz" : "Anonymous group quiz" }), progress);
    const body = el("section");
    content.replaceChildren(header, body);
    if (state.phase === "lobby") {
      body.append(el("h2", { textContent: credential!.host ? "Scan to join" : "You're in!" }));
      if (credential!.host) {
        const url = `${credential!.base || location.origin + location.pathname}#/group/${id}`;
        const canvas = el("canvas", { className: "join-qr" });
        canvas.setAttribute("aria-label", "QR code for joining this quiz");
        body.append(canvas);
        void QRCode.toCanvas(canvas, url, { width: 300, margin: 4, errorCorrectionLevel: "M" }).catch(error);
        const link = el("input", { className: "join-link", value: url, readOnly: true });
        link.setAttribute("aria-label", "Join link");
        body.append(link, button("Copy join link", async () => {
          try { await navigator.clipboard.writeText(url); connection.textContent = "Join link copied."; }
          catch { link.select(); connection.textContent = "Select and copy the join link above."; }
        }), el("p", { className: "hint", textContent: "Phones must be able to reach this address. For localhost hosting, configure -public-url with your reachable server address." }),
        button("Start quiz", () => action("start")));
      } else body.append(el("p", { textContent: "Waiting for the host to start. No name or account is needed." }));
    } else if (state.phase === "finished") {
      const summary = state.summary!;
      body.append(el("h2", { textContent: "Group summary" }),
        el("p", { className: "score", textContent: `${summary.average.toFixed(1)}% average score` }),
        el("p", { className: "hint", textContent: `Across ${summary.active} participants who submitted at least once; unanswered questions earn zero. All ${state.total} questions count toward the score.` }),
        el("p", { textContent: `${summary.completed} completed the entire quiz · ${summary.submissions} total submissions` }));
      for (const stat of summary.questions) body.append(el("h2", { textContent: `${stat.number}. ${stat.text}` }),
        el("p", { textContent: `${stat.submitted ? (100 * stat.correct / stat.submitted).toFixed(0) : "0"}% correct among ${stat.submitted} submissions · ${stat.unanswered} unanswered` }));
      body.append(button("Back to quizzes", async () => home()));
    } else if (state.question) {
      body.append(el("p", { className: "progress", textContent: `Question ${state.number} of ${state.total}` }), renderPrompt(state.question));
      if (state.phase === "question") {
        if (credential!.host) {
          // Show choices on the projector, but never submit as the host.
          const inputs = renderInputs(state.question); inputs.disable(); body.append(inputs.node, button("Reveal results", () => action("reveal")));
        } else if (state.answered) body.append(el("p", { textContent: "Answer submitted. Waiting for the host to reveal results…" }));
        else {
          const inputs = renderInputs(state.question);
          body.append(inputs.node, button("Submit answer", async () => {
            const answer = inputs.read();
            if (!answer.length) throw new Error("Please give an answer first.");
            await request(`/${id}/answers/${state.number}`, credential!.token, { answer });
            await refresh();
          }));
        }
      } else {
        const inputs = renderInputs(state.question); inputs.disable(); inputs.mark(state.correctAnswer ?? []);
        body.append(inputs.node);
        if (state.question.type === "text") body.append(el("p", { textContent: `Expected: ${state.correctAnswer?.join(" / ")}` }));
        if (state.explanation) body.append(el("p", { className: "explanation", textContent: state.explanation }));
        if (!credential!.host) body.append(el("p", { className: "verdict", textContent: state.personal ? state.personal.correct ? `Your answer was correct (+${state.personal.points}).` : "Your answer was not correct." : "You did not answer this question." }));
        if (state.stats) body.append(statistics(state.stats, state.question));
        if (credential!.host) body.append(state.number < state.total ? button("Next question", () => action("next")) : button("Finish quiz", () => action("finish")));
        else body.append(el("p", { className: "hint", textContent: "Waiting for the host…" }));
      }
    }
    if (credential!.host) {
      const actions = el("div", { className: "actions" });
      if (state.phase !== "finished") actions.append(button("Finish early", () => action("finish")));
      actions.append(button("Delete room", () => action("delete")));
      body.append(actions);
      if (state.phase !== "lobby" && state.phase !== "finished") {
        const url = `${credential!.base || location.origin + location.pathname}#/group/${id}`;
        body.append(el("p", { className: "hint" }, "Late arrivals can join: ", el("a", { href: url, textContent: url, target: "_blank", rel: "noopener" })));
      }
    }
  }
  let fetching = false;
  async function refresh() {
    if (stopped || fetching) return;
    fetching = true;
    try {
      const next = await request<State>(`/${id}`, credential!.token);
      if (stopped) return;
      state = next; connection.textContent = ""; render();
    } catch (err) {
      if (stopped) return;
      if (err instanceof HTTPError && (err.status === 404 || err.status === 403)) terminal(err.message);
      else connection.textContent = "Connection interrupted. Retrying… " + (err as Error).message;
    } finally { fetching = false; }
  }
  async function poll() { await refresh(); if (!stopped) timer = setTimeout(() => void poll(), 2000); }
  function loadGroup() {
    content.replaceChildren(el("p", { className: "loading", textContent: "Loading group…" }));
    void poll();
  }
  if (credential) loadGroup();
  else content.replaceChildren(el("h1", { textContent: "Join anonymous group quiz" }),
    el("p", { textContent: "No name or account. Results are shown only as group totals. Keep this tab open to retain your place." }),
    button("Join", async () => {
      if (stopped || credential || joining) return;
      joining = true;
      try {
        const result = await request<{ token: string }>(`/${id}/join`, "", {});
        if (stopped) return;
        credential = { token: result.token, host: false };
        sessionStorage.setItem(key(id), JSON.stringify(credential));
        // Joining is complete even if the first state fetch is slow or fails.
        // Remove Join now; polling retries with the same participant credential.
        loadGroup();
      } finally {
        joining = false;
      }
    }));
  return () => { stopped = true; clearTimeout(timer); };
}
