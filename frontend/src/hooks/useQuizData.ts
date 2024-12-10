import { useState, useEffect } from 'react';
import { QuizData } from '../types';
import fetchQuizData from './utils/fetchQuizData';
import mapQuizData from './utils/mapQuizData';

const useQuizData = (isLoggedIn: boolean, todaysFinish: boolean, isAdmin: boolean, isSubmitAnsewr: boolean, setTodaysFinish: (finish: boolean) => void) => {
  const [quizData, setQuizData] = useState<QuizData[]>([]);
  const [dataFetched, setDataFetched] = useState<boolean>(false);

  useEffect(() => {
    if (isLoggedIn && !isAdmin) {
      const fetchData = async () => {
        const response = await fetchQuizData(isSubmitAnsewr);
        if (response.todays_finish) {
          setTodaysFinish(true);
          return; // ノルマが達成されている場合は早期にリターン
        }
        if (response.quizdata) {
          const mappedData = mapQuizData(response.quizdata);
          setQuizData(mappedData);
        }
        setDataFetched(true);
      };

      fetchData(); // 初回レンダリング時にデータをフェッチ

      const interval = setInterval(() => {
        if (!dataFetched && !todaysFinish && !isAdmin) {
          fetchData();
        }
      }, 10000);

      return () => clearInterval(interval); // useEffect フックのクリーンアップ関数
    }
  }, [isLoggedIn, dataFetched, setTodaysFinish]);

  return quizData;
};

export default useQuizData;
