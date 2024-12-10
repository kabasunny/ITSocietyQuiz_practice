import { Dispatch, SetStateAction } from 'react';
import { Answer, Option, QuizData } from '../../types';

const goToNextQuestion = (
  setNext: Dispatch<SetStateAction<boolean>>,
  setCurrentQuestion: Dispatch<SetStateAction<number>>,
  setShowScore: Dispatch<SetStateAction<boolean>>,
  answers: Answer[],
  quizData: QuizData[]
) => {
  const answerLength = answers.length; // 定数にすることで繰り返しを防ぐ

  if (answerLength < quizData.length) {
    setCurrentQuestion(answerLength);
  } else if (answerLength >= quizData.length) { // 最後の問題の回答後
    setShowScore(true); // 結果のスコアを表示するためのフラグ
  }

  setNext(false);
};

export default goToNextQuestion;
