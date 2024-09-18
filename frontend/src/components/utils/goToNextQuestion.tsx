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
  
  // 現在の問題番号をセッションに保存する関数
  const saveCurrentQuestionToSession = (question: number) => {
    sessionStorage.setItem('currentQuestion', JSON.stringify(question));
  };

  if (nextQuestion < quizData.length) { // 次の問題がまだある場合
    setCurrentQuestion((prevQuestion) => {
      const updatedQuestion = prevQuestion + 1;
      saveCurrentQuestionToSession(updatedQuestion); // セッションに保存
      return updatedQuestion;
    });
  } else {
    submitAnswers(answers, setIsSubmitAnsewr); // APIに結果を送信
    setShowScore(true); // 結果のスコアを表示するためのフラグ
  }

  setNext(false);
};

export default goToNextQuestion;
