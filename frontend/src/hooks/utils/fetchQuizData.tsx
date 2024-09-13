import axios from 'axios';
import { Question } from '../../types';

const fetchQuizData = async (): Promise<Question[]> => {
  try {
    const jwt = localStorage.getItem('token'); // ログイン時にAPIから取得したトークン
    const todaysCount = localStorage.getItem('todays_count'); // ローカルストレージからtodays_countを取得

    const response = await axios.get(`http://localhost:8082/questions/oneday?todays_count=${todaysCount}`, { // とりあえず本日の回答済みクイズ数をURLパラメータに入れる
      headers: {
        'Authorization': `Bearer ${jwt}`
      }
    });

    return response.data.quizdata;
  } catch (error) {
    console.error('クイズデータの取得中にエラーが発生しました:', error);
    return [];
  }
};

export default fetchQuizData;
