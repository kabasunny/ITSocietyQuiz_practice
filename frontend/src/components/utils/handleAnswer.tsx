import { Answer, Option } from '../../types';
import { Dispatch, SetStateAction } from 'react';

const handleAnswer = (
  answer: Option,
  quizData: any[],
  currentQuestion: number,
  setScore: Dispatch<SetStateAction<number>>,
  setFeedback: Dispatch<SetStateAction<string | null>>,
  setAnswers: Dispatch<SetStateAction<Answer[]>>,
  setNext: Dispatch<SetStateAction<boolean>>,
  answers: Answer[]
) => {
  const newAnswer: Answer = {
    question: quizData[currentQuestion].question,
    answer: answer,
    correct: answer.text === quizData[currentQuestion].correct,
  };

  if (newAnswer.correct) {
    setScore((prevScore) => prevScore + 1);
    setFeedback("○");
  } else {
    setFeedback("×");
  }
  setAnswers([...answers, newAnswer]);
  setNext(true);
};

export default handleAnswer;
