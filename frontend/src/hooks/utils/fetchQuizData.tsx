import axios from 'axios';
import shuffleArray from './shuffleArray';

const fetchQuizData = async (setQuizData: (data: any[]) => void, setDataFetched: (fetched: boolean) => void) => {
  try {
    const response = await axios.get('http://localhost:8082/questions/oneday');
    const mappedData = response.data.data.map((item: any) => {
      const correctAnswer = item.Options[0];
      const optionsWithIndex = item.Options.map((option: string, index: number) => ({
        text: option,
        index: index
      }));
      const shuffledOptions = shuffleArray([...optionsWithIndex]);
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

export default fetchQuizData;
