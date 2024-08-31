import { Dispatch, SetStateAction } from 'react';

const goToNextQuestion = (
  currentQuestion: number,
  setCurrentQuestion: Dispatch<SetStateAction<number>>,
  quizData: any[],
  handleResultsSubmit: () => void,
  setShowScore: Dispatch<SetStateAction<boolean>>,
  setNext: Dispatch<SetStateAction<boolean>>
) => {
  const nextQuestion = currentQuestion + 1;
  if (nextQuestion < quizData.length) {
    setCurrentQuestion(nextQuestion);
  } else {
    handleResultsSubmit();
    setShowScore(true);
  }
  setNext(false);
};

export default goToNextQuestion;
