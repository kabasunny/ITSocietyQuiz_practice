import axios from 'axios';
import { Answer, ResAnswer } from '../../types';
import { Dispatch, SetStateAction } from 'react';

const submitAnswers = async (
  answers: Answer[],
  setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>>
) => {
  const jwt = sessionStorage.getItem('token'); // ログイン時にAPIから取得したトークン
  const storedQuizData = sessionStorage.getItem('quizdata');
  var todaysFinish = false;

  if (jwt) {
    // Answer を ResAnswer に変換
    const resAnswers: ResAnswer[] = answers.map((answer) => ({
      question_id: answer.question_id,
      answer_id: answer.answer_id,
    }));

    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API_URL}/answers`,
        resAnswers,
        {
          headers: {
            Authorization: `Bearer ${jwt}`,
            'Content-Type': 'application/json',
          },
        }
      );
      // response.data.todays_finishを変数に取り出す
      if (answers.length === storedQuizData?.length) {
        todaysFinish = true;
      }

      // 変数をコンソールに表示
      console.log('todays_finish:', todaysFinish);

      // セッションに保存
      sessionStorage.setItem('todays_finish', JSON.stringify(todaysFinish));

      setIsSubmitAnsewr(true); // 成功時にsetIsSubmitAnsewr(true)を呼び出す
      sessionStorage.removeItem('quizdata');
    } catch (error) {
      alert('結果の送信中にエラーが発生しました。');
    }
  } else {
    alert('トークンが見つかりません');
  }
};

export default submitAnswers;
