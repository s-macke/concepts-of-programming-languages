// Mirrors the JSON produced by the Go server. Note what is absent: a question
// never carries its correct answer, that only arrives inside an AnswerResult
// after the answer has been submitted.

export type QuestionType = "single" | "multiple" | "truefalse" | "text";

export interface QuizImage {
  data: string;
  mime?: string;
  alt?: string;
}

export interface Code {
  lang?: string;
  source: string;
}

export interface Option {
  id: string;
  text: string;
  image?: QuizImage;
}

export interface Question {
  type: QuestionType;
  text: string;
  code?: Code;
  image?: QuizImage;
  options?: Option[];
  points: number;
}

export interface QuizListEntry {
  id: string;
  title: string;
  description?: string;
  numQuestions: number;
}

export interface SessionInfo {
  sessionId: string;
  title: string;
  description?: string;
  numQuestions: number;
  totalPoints: number;
}

export interface AnswerResult {
  correct: boolean;
  points: number;
  correctAnswer: string[];
  explanation?: string;
}

export interface Summary {
  answered: number;
  numQuestions: number;
  points: number;
  totalPoints: number;
  correct: number;
}
