import { Dispatch, SetStateAction } from 'react';
import { Answer, Option, QuizData } from '../../types';

const goToNextQuestion = (
  setNext: Dispatch<SetStateAction<boolean>>,
  setCurrentQuestion: Dispatch<SetStateAction<number>>,
  answers: Answer[],
  quizData: QuizData[]
) => {
  if (answers.length < quizData.length) {
    setCurrentQuestion(answers.length);
  }
  setNext(false);
};
export default goToNextQuestion;
