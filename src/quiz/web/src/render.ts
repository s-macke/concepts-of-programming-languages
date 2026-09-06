import type { Question, QuizImage } from "./types";

/** Small helper to build an element with attributes and children in one call. */
export function el<K extends keyof HTMLElementTagNameMap>(
  tag: K,
  props: Partial<HTMLElementTagNameMap[K]> = {},
  ...children: (Node | string)[]
): HTMLElementTagNameMap[K] {
  const node = Object.assign(document.createElement(tag), props);
  for (const child of children) node.append(child);
  return node;
}

/** Images travel inside the quiz file as base64, so they become data URLs. */
export function renderImage(image: QuizImage): HTMLImageElement {
  return el("img", {
    className: "quiz-image",
    src: `data:${image.mime ?? "image/png"};base64,${image.data}`,
    alt: image.alt ?? "",
  });
}

export function renderCode(source: string, lang?: string): HTMLElement {
  const code = el("code", { textContent: source });
  if (lang) code.classList.add(`language-${lang}`);
  return el("pre", { className: "quiz-code" }, code);
}

/**
 * Builds the input controls for a question and returns a function that reads
 * the current answer. Which control is used follows from the question type:
 * radio buttons, checkboxes or a text field.
 */
export function renderInputs(question: Question): {
  node: HTMLElement;
  read: () => string[];
  disable: () => void;
  mark: (correct: string[]) => void;
} {
  if (question.type === "text") {
    const input = el("input", {
      className: "text-answer",
      type: "text",
      autocomplete: "off",
      placeholder: "Your answer",
    });
    return {
      node: el("div", { className: "answers" }, input),
      read: () => (input.value.trim() === "" ? [] : [input.value]),
      disable: () => (input.disabled = true),
      mark: () => {},
    };
  }

  const multiple = question.type === "multiple";
  const inputs: HTMLInputElement[] = [];
  const labels = new Map<string, HTMLLabelElement>();

  const list = el("div", { className: "answers" });
  for (const option of question.options ?? []) {
    const input = el("input", {
      type: multiple ? "checkbox" : "radio",
      name: "answer",
      value: option.id,
    });
    inputs.push(input);

    const label = el("label", { className: "option" }, input, el("span", { textContent: option.text }));
    if (option.image) label.append(renderImage(option.image));
    labels.set(option.id, label);
    list.append(label);
  }

  return {
    node: list,
    read: () => inputs.filter((i) => i.checked).map((i) => i.value),
    disable: () => inputs.forEach((i) => (i.disabled = true)),
    mark: (correct) => {
      for (const input of inputs) {
        const label = labels.get(input.value)!;
        if (correct.includes(input.value)) label.classList.add("correct");
        else if (input.checked) label.classList.add("wrong");
      }
    },
  };
}

/** Renders the immutable part of a question: text, code snippet and image. */
export function renderPrompt(question: Question): HTMLElement {
  const prompt = el("div", { className: "prompt" }, el("p", { className: "question-text", textContent: question.text }));
  if (question.code) prompt.append(renderCode(question.code.source, question.code.lang));
  if (question.image) prompt.append(renderImage(question.image));
  return prompt;
}
