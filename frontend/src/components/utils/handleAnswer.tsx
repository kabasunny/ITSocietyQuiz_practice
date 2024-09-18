import { Answer, Option, QuizData } from '../../types';
import { Dispatch, SetStateAction } from 'react';

const handleAnswer = (
  selectedAnswer: Option,
  quizData: QuizData[],
  currentQuestion: number,
  setScore: Dispatch<SetStateAction<number>>,
  setFeedback: Dispatch<SetStateAction<string | null>>,
  setAnswers: Dispatch<SetStateAction<Answer[]>>,
  setNext: Dispatch<SetStateAction<boolean>>,
  answers: Answer[]
) => {
  const newAnswer: Answer = {
    question_id: quizData[currentQuestion].id,
    question: quizData[currentQuestion].question,
    answer_id: selectedAnswer.index,
    answer_text: selectedAnswer.text,
    correct: selectedAnswer.index === 0,
  };

  if (newAnswer.correct) {
    setScore((prevScore) => prevScore + 1);
    setFeedback("○");
  } else {
    setFeedback("×");
  }
  const updatedAnswers = [...answers, newAnswer];
  setAnswers(updatedAnswers);
  sessionStorage.setItem('answers', JSON.stringify(updatedAnswers)); // セッションストレージに保存
  setNext(true);
};

export default handleAnswer;
