import { Answer, Option, QuizData } from '../../types';
import { Dispatch, SetStateAction } from 'react';

const handleAnswer = (
  selectedAnswer: Option,
  quizData: QuizData[],
  currentQuestion: number,
  // setCurrentQuestion: Dispatch<SetStateAction<number>>,
  setScore: Dispatch<SetStateAction<number>>,
  setFeedback: Dispatch<SetStateAction<string | null>>,
  setAnswers: Dispatch<SetStateAction<Answer[]>>,
  setNext: Dispatch<SetStateAction<boolean>>,
  submitAnswers: (answers: Answer[], setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>> ) => void,
  setShowScore: Dispatch<SetStateAction<boolean>>,
  setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>>,
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
    setScore((prevScore) => {
      const newScore = prevScore + 1;
      console.log('newScore:', newScore);
      sessionStorage.setItem('score', JSON.stringify(newScore)); // セッションストレージに保存
      return newScore;
    });

    setFeedback("○");
  } else {
    setFeedback("×");
  }
  const updatedAnswers = [...answers, newAnswer];
  setAnswers(updatedAnswers);
  sessionStorage.setItem('answers', JSON.stringify(updatedAnswers)); // セッションストレージに保存

  console.log('answersの長さ:', updatedAnswers.length); // answersの長さを表示
  

  // 現在の問題番号をセッションに保存する関数
  const saveCurrentQuestionToSession = (question: number) => {
    sessionStorage.setItem('currentQuestion', JSON.stringify(question));
  };

  // if (updatedAnswers.length < quizData.length) { // 次の問題がまだある場合
  //   setCurrentQuestion(() => {
  //     const updatedQuestion = updatedAnswers.length;
  //     saveCurrentQuestionToSession(updatedQuestion); // セッションに保存
  //     return updatedQuestion;
  //   });
  // } else {
  //   submitAnswers(updatedAnswers, setIsSubmitAnsewr); // APIに結果を送信
  //   setShowScore(true); // 結果のスコアを表示するためのフラグ
  // }

  const updatedQuestion = updatedAnswers.length;
  saveCurrentQuestionToSession(updatedQuestion); // セッションに保存

  if (updatedAnswers.length >= quizData.length) { // 最後の問題の回答後
    submitAnswers(updatedAnswers, setIsSubmitAnsewr); // APIに結果を送信
    setShowScore(true); // 結果のスコアを表示するためのフラグ
  }
  
  setNext(true);
};

export default handleAnswer;
