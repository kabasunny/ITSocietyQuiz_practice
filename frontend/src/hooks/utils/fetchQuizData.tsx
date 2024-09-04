import axios from 'axios';
import { Question } from '../../types';

const fetchQuizData = async (): Promise<Question[]> => {
  try {
    const response = await axios.get('http://localhost:8082/questions/oneday');
    return response.data.data;
  } catch (error) {
    console.error('クイズデータの取得中にエラーが発生しました:', error);
    return [];
  }
};

export default fetchQuizData;
