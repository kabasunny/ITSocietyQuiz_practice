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

  const shuffleArray = (array: any[]) => {
    for (let i = array.length - 1; i > 0; i--) {
      const j = Math.floor(Math.random() * (i + 1));
      [array[i], array[j]] = [array[j], array[i]];
    }
    return array;
  };

  useEffect(() => {
    console.log('REACT_APP_USE_API:', process.env.REACT_APP_USE_API);
    if (process.env.REACT_APP_USE_API === 'true') {
      const fetchQuizData = async () => {
        try {
          // const response = await axios.get('http://localhost:8082/quiz_data');
          const response = await axios.get('http://localhost:8082/questions/oneday');
          const mappedData = response.data.data.map((item: any) => {
            const correctAnswer = item.Options[0];
            const shuffledOptions = shuffleArray([...item.Options]);
            return {
              question: item.Question,
              options: shuffledOptions,
              correct: correctAnswer,
              supplement: item.Supplement,
            };
          });
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
      const mappedTestQuizData = testQuizData.map((item) => {
        const correctAnswer = item.options[0];
        const shuffledOptions = shuffleArray([...item.options]);
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
