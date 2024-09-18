import axios from 'axios';
import { Question } from '../../types';

interface QuizResponse {
  quizdata: Question[] | null;
  todays_finish: boolean;
}

const fetchQuizData = async (): Promise<QuizResponse> => {
  try {
    const jwt = sessionStorage.getItem('token'); // ログイン時にAPIから取得したトークン
    const todaysCount = sessionStorage.getItem('todays_count'); // ローカルストレージからtodays_countを取得
    

    const response = await axios.get(`http://localhost:8082/questions/oneday?todays_count=${todaysCount}`, { // とりあえず本日の回答済みクイズ数をURLパラメータに入れる
      headers: {
        'Authorization': `Bearer ${jwt}`
      }
    });

    return {
      quizdata: response.data.quizdata,
      todays_finish: response.data.todays_finish
    };
  } catch (error) {
    console.error('クイズデータの取得中にエラーが発生しました:', error);
    return {
      quizdata: null,
      todays_finish: false
    };
  }
};

export default fetchQuizData;
