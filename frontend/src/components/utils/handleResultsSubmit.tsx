import axios from 'axios';
import { Answer, ResAnswer } from '../../types';

const handleResultsSubmit = async (answers: Answer[]) => {
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
      // alert(`結果が送信されました。\nレスポンス: ${JSON.stringify(response.data, null, 2)}`);
    } catch (error) {
      
      alert('結果の送信中にエラーが発生しました。');
    }
  } else {
    alert('トークンが見つかりません');
  }
};

export default handleResultsSubmit;
