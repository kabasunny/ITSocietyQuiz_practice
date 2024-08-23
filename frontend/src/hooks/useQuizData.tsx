import { useState, useEffect } from 'react';
import axios from 'axios';
import { quizData as testQuizData } from '../data/quizData';

interface QuizData {
  question: string;
  options: string[];
  correct: string;
  supplement: string;
}

const useQuizData = () => {
  const [quizData, setQuizData] = useState<QuizData[]>([]);
  const [dataFetched, setDataFetched] = useState<boolean>(false);

  useEffect(() => {
    console.log('REACT_APP_USE_API:', process.env.REACT_APP_USE_API);
    if (process.env.REACT_APP_USE_API === 'true') {
      const fetchQuizData = async () => {
        try {
          const response = await axios.get('http://localhost:8082/quiz_data');
          const mappedData = response.data.data.map((item: any) => ({
            question: item.Question,
            options: item.Options,
            correct: item.Correct,
            supplement: item.Supplement,
          }));
          setQuizData(mappedData);
          setDataFetched(true);
        } catch (error) {
          console.error('クイズデータの取得中にエラーが発生しました:', error);
        }
      };

      fetchQuizData();

      const interval = setInterval(() => {
        if (!dataFetched) {
          fetchQuizData();
        } else {
          clearInterval(interval);
        }
      }, 10000);

      return () => clearInterval(interval);
    } else {
      setQuizData(testQuizData);
    }
  }, [dataFetched]);

  return quizData;
};

export default useQuizData;
