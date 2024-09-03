export interface LoginProps {
  onLogin: (data: boolean) => void;
}

export interface LoginForm {
  empid: string;
  password: string;
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


export interface QuizData {
  id: number; // GORM.Modelから
  question: string;
  options: Option[]; // 選択肢シャッフル後も、インデックスを保持する目的
  correct: string;
  supplement: string;
}

export interface QuizProps {
  currentQuestion: number;
  quizData: {
    question: string;
    correct: string;
    supplement: string;
    options: Option[];
  }[];
  next: boolean;
  feedback: string | null;
  handleAnswer: (answer: Option) => void;
  goToNextQuestion: () => void;
}


  

