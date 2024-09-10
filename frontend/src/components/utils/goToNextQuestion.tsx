import { Dispatch, SetStateAction } from 'react';
import { QuizData, Answer } from '../../types';

const goToNextQuestion = (
  currentQuestion: number,
  setCurrentQuestion: Dispatch<SetStateAction<number>>,
  quizData: QuizData[],
  answers: Answer[],
  submitAnswers: (answers: Answer[], setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>>) => void,
  setShowScore: Dispatch<SetStateAction<boolean>>,
  setNext: Dispatch<SetStateAction<boolean>>,
  setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>>
) => {
  const nextQuestion = currentQuestion + 1;
  if (nextQuestion < quizData.length) { // API側で、1日のクイズ数は5問を想定
    setCurrentQuestion(nextQuestion);
  } else {
    submitAnswers(answers, setIsSubmitAnsewr); // APIに結果を送信
    setShowScore(true); // 結果のスコアを表示するためのフラグ
  }
  setNext(false);
};

export default goToNextQuestion;
