import type {
  AnswerResult,
  Question,
  QuizListEntry,
  SessionInfo,
  Summary,
} from "./types";

/** Thrown for any non-2xx response, carrying the server's error message. */
export class ApiError extends Error {
  constructor(readonly status: number, message: string) {
    super(message);
  }
}

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const response = await fetch(path, init);
  if (!response.ok) {
    let message = response.statusText;
    try {
      const body = (await response.json()) as { error?: string };
      if (body.error) message = body.error;
    } catch {
      // Not a JSON error body; keep the status text.
    }
    throw new ApiError(response.status, message);
  }
  return (await response.json()) as T;
}

export const api = {
  listQuizzes: () => request<QuizListEntry[]>("/api/quizzes"),

  startSession: (quizId: string) =>
    request<SessionInfo>("/api/sessions", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ quizId }),
    }),

  upload: (file: File) => {
    const form = new FormData();
    form.append("file", file);
    return request<SessionInfo>("/api/upload", { method: "POST", body: form });
  },

  /** Question numbers are one based in the URL, matching what the user sees. */
  question: (sessionId: string, n: number) =>
    request<Question>(`/api/sessions/${sessionId}/questions/${n}`),

  answer: (sessionId: string, n: number, answer: string[]) =>
    request<AnswerResult>(`/api/sessions/${sessionId}/answers/${n}`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ answer }),
    }),

  result: (sessionId: string) =>
    request<Summary>(`/api/sessions/${sessionId}/result`),
};
