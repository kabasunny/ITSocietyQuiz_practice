
import { NavigateFunction } from 'react-router-dom';

export interface LoginProps {
  onLogin: (loginOK: boolean, isAdmin: boolean) => void;
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
  isSubmitAnsewr: boolean;
  handleLogout: (setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>, setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>, setIsSubmitAnsewr: React.Dispatch<React.SetStateAction<boolean>>, navigate: NavigateFunction) => void;
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>;
  setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>;
  setIsSubmitAnsewr: React.Dispatch<React.SetStateAction<boolean>>;
  setShowScore: React.Dispatch<React.SetStateAction<boolean>>;
  navigate: NavigateFunction;
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


export interface AdmQuestion {
  id: number;
  userQuestionID: string | null;
  question: string;
  options: string[];
  supplement: string;
  difficulty: number;
  createdAt: string; // 作成日を追加
}