// ログイン関連のインターフェース
export interface LoginProps {
  onLogin: (data: boolean) => void;
}

export interface LoginForm {
  empid: string;
  password: string;
}

// クイズ関連のインターフェース
export interface QuizData {
  id: number; // GORM.Modelから
  question: string;
  options: Option[]; // 選択肢シャッフル後も、インデックスを保持する目的
  correct: string;
  supplement: string;
}

export interface QuizProps {
  currentQuestion: number;
  quizData: QuizData[];
  next: boolean;
  feedback: string | null;
  handleAnswer: (selectedAnswer: Option) => void;
  goToNextQuestion: () => void;
}

export interface Question {
  id: number; // GORM.Modelから
  question: string;
  options: string[];
  supplement: string;
  difficulty: number;
}

export interface Option {
  text: string;
  index: number;
}

export interface Answer {
  question_id: number;
  question: string;
  answer_id: number;
  answer_text: string;
  correct: boolean;
}

export interface ResAnswer {
  question_id: number;
  answer_id: number;
}

// スコア関連のインターフェース
export interface ScoreSectionProps {
  score: number;
  answers: Answer[];
}

// ユーザー関連のインターフェース
export interface Role {
  id: number;
  role_name: string;
}

export interface User {
  empid: string;
  username: string;
  email: string;
  password: string;
  total_questions: number;
  correct_answers: number;
}
