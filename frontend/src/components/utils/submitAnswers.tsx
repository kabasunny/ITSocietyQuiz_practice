import axios from 'axios';
import { Answer, ResAnswer } from '../../types';
import { Dispatch, SetStateAction } from 'react';

const submitAnswers = async (
  answers: Answer[],
  setIsSubmitAnsewr: Dispatch<SetStateAction<boolean>>
) => {
    const jwt = sessionStorage.getItem('token'); // ログイン時にAPIから取得したトークン
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
        console.log('APIレスポンス:', response.data); // レスポンスデータをログに出力
        setIsSubmitAnsewr(true); // 成功時にsetIsSubmitAnsewr(true)を呼び出す
      } catch (error) {
        alert('結果の送信中にエラーが発生しました。');
      }
    } else {
      alert('トークンが見つかりません');
    }
 
};

export default submitAnswers;
