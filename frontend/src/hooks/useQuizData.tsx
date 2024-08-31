import { useState, useEffect } from 'react';
import { questions as testQuizData } from '../data/Questions';
import { QuizData } from '../types';
import shuffleArray from './utils/shuffleArray';
import fetchQuizData from './utils/fetchQuizData';

const useQuizData = () => {
  const [quizData, setQuizData] = useState<QuizData[]>([]);
  const [dataFetched, setDataFetched] = useState<boolean>(false);

  useEffect(() => {
    console.log('REACT_APP_USE_API:', process.env.REACT_APP_USE_API);
    if (process.env.REACT_APP_USE_API === 'true') {
      fetchQuizData(setQuizData, setDataFetched);

      const interval = setInterval(() => {
        if (!dataFetched) {
          fetchQuizData(setQuizData, setDataFetched);
        } else {
          clearInterval(interval);
        }
      }, 10000);

      return () => clearInterval(interval);
    } else {  //  process.env.REACT_APP_USE_API === 'false' ローカルデータでの検証
      const mappedTestQuizData = testQuizData.map((item) => {
        const correctAnswer = item.options[0];
        const optionsWithIndex = item.options.map((option: string, index: number) => ({
          text: option,
          index: index
        }));
        const shuffledOptions = shuffleArray([...optionsWithIndex]);
        return {
          ...item,
          options: shuffledOptions,
          correct: correctAnswer,
        };
      });
      setQuizData(mappedTestQuizData);
    }
  }, [dataFetched]);

  return quizData;
};

export default useQuizData;
