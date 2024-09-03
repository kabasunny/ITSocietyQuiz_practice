import { Dispatch, SetStateAction } from 'react';
import { QuizData } from '../../types';

const goToNextQuestion = (
  currentQuestion: number,
  setCurrentQuestion: Dispatch<SetStateAction<number>>,
  quizData: QuizData[],
  handleResultsSubmit: () => void,
  setShowScore: Dispatch<SetStateAction<boolean>>,
  setNext: Dispatch<SetStateAction<boolean>>
) => {
  const nextQuestion = currentQuestion + 1;
  if (nextQuestion < quizData.length) { // 1日のクイズ数は5問を想定
    setCurrentQuestion(nextQuestion);
  } else {
    handleResultsSubmit(); // APIに結果を送信
    setShowScore(true); // 結果のスコアを表示するためのフラグ
  }
  setNext(false);
};

export default goToNextQuestion;
