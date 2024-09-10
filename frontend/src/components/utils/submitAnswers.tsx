import axios from 'axios';
import { Answer, ResAnswer } from '../../types';
import { Dispatch, SetStateAction } from 'react';

const submitAnswers = async (
  answers: Answer[],
  setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>>
) => {
  if (process.env.REACT_APP_USE_API === 'true') {
    const jwt = localStorage.getItem('token'); // ログイン時にAPIから取得したトークン
    if (jwt) {
      // Answer を ResAnswer に変換
      const resAnswers: ResAnswer[] = answers.map(answer => ({
        question_id: answer.question_id,
        answer_id: answer.answer_id
      }));

      try {
        const response = await axios.post('http://localhost:8082/answers', resAnswers, {
          headers: {
            'Authorization': `Bearer ${jwt}`,
            'Content-Type': 'application/json'
          }
        });
        setIsSubmitAnsewr(true); // 成功時にsetIsSubmitAnsewr(true)を呼び出す
      } catch (error) {
        alert('結果の送信中にエラーが発生しました。');
      }
    } else {
      alert('トークンが見つかりません');
    }
  } else {
    setIsSubmitAnsewr(true);
    console.log('API送信は無効化されています');
  }
};

export default submitAnswers;
