import { QuizData, Option, Question } from '../../types';
import shuffleArray from './shuffleArray';

const mapQuizData = (data: Question[]): QuizData[] => {
  return data.map((item) => {
    const correctAnswer: string = item.options[0];
    const optionsWithIndex: Option[] = item.options.map((option: string, index: number) => ({
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
};

export default mapQuizData;
