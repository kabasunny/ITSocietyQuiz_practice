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
  question: string;
  answer: Option;
  correct: boolean;
}

export interface QuizData {
  question: string;
  options: Option[];
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


  

