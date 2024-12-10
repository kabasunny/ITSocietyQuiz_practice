import axios from 'axios';
import { Question } from '../../types';

interface QuizResponse {
  quizdata: Question[] | null;
  todays_finish: boolean;
}

const fetchQuizData = async (
  isSubmitAnswer: boolean
): Promise<QuizResponse> => {
  try {
    // まずセッションストレージに保存されているデータを確認
    const storedQuizData = sessionStorage.getItem('quizdata');
    const storedTodaysFinish = sessionStorage.getItem('todays_finish');

    // 既にデータがある場合、そのまま返す
    if (storedQuizData && storedTodaysFinish && !isSubmitAnswer) {
      return {
        quizdata: JSON.parse(storedQuizData),
        todays_finish: JSON.parse(storedTodaysFinish) === 'true',
      };
    }

    // セッションストレージにデータがない場合、APIリクエストを行う
    const jwt = sessionStorage.getItem('token'); // ログイン時にAPIから取得したトークン
    const todaysCount = sessionStorage.getItem('todays_count'); // ローカルストレージからtodays_countを取得

    const response = await axios.get(
      `${process.env.REACT_APP_API_URL}/questions/oneday?todays_count=${todaysCount}`,
      {
        headers: {
          Authorization: `Bearer ${jwt}`,
        },
      }
    );

    // 取得したデータをコンソールに出力
    console.log('APIレスポンス:', response.data);

    // 取得したデータをセッションストレージに保存
    sessionStorage.setItem('quizdata', JSON.stringify(response.data.quizdata));
    sessionStorage.setItem(
      'todays_finish',
      JSON.stringify(response.data.todays_finish)
    );

    return {
      quizdata: response.data.quizdata,
      todays_finish: response.data.todays_finish,
    };
  } catch (error) {
    console.error('クイズデータの取得中にエラーが発生しました:', error);
    return {
      quizdata: null,
      todays_finish: false,
    };
  }
};

export default fetchQuizData;
