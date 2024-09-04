import { useState, useEffect } from 'react';
import { questions as testQuizData } from '../data/Questions';
import { QuizData } from '../types';
import fetchQuizData from './utils/fetchQuizData';
import mapQuizData from './utils/mapQuizData';

const useQuizData = () => {
  const [quizData, setQuizData] = useState<QuizData[]>([]);
  const [dataFetched, setDataFetched] = useState<boolean>(false);

  useEffect(() => {
    console.log('REACT_APP_USE_API:', process.env.REACT_APP_USE_API);
    const fetchData = async () => {
      let data;
      if (process.env.REACT_APP_USE_API === 'true') {
        data = await fetchQuizData();
      } else {
        data = testQuizData;
      }
      const mappedData = mapQuizData(data);
      setQuizData(mappedData);
      setDataFetched(true);
    };

    fetchData(); // 初回レンダリング時にデータをフェッチ

    const interval = setInterval(() => {
      if (!dataFetched) {
        fetchData();
      }
    }, 10000);

    return () => clearInterval(interval); // useEffect フックのクリーンアップ関数
  }, [dataFetched]);

  return quizData;
};

export default useQuizData;
