import axios from 'axios';
import shuffleArray from './shuffleArray';
import { QuizData, Option } from '../../types';

interface ApiQuizItem { // APIから受け取る型、QuizData型へ変換用
  Question: string;
  Options: string[];
  Supplement: string;
  Difficulty: number;
}

const fetchQuizData = async (setQuizData: (quizData: QuizData[]) => void, setDataFetched: (fetched: boolean) => void) => {
  try {
    const response = await axios.get('http://localhost:8082/questions/oneday');
    // ここでレスポンスは、Go言語のdtoで設計した (modelの構造になっていたため)
    // type QuizData struct {
    // 	Question   string   `json:"question"`
    // 	Options    []string `json:"options"`
    // 	Supplement string   `json:"supplement"`
    // 	Difficulty int      `json:"difficulty"` ( ← いらないかも)
    // } となるようにバック側のQuestionsServiceを修正した
    const mappedData: QuizData[] = response.data.quizData.map((item: ApiQuizItem) => {
      const correctAnswer: string = item.Options[0]; // レスポンスのOptionsキーの最初の配列が正解
      const optionsWithIndex: Option[]  = item.Options.map((option: string, index: number) => ({
        text: option,
        index: index
      })); // オブジェクト.array.map((element, index, array)
      const shuffledOptions = shuffleArray([...optionsWithIndex]);
      return {
        question: item.Question,
        options: shuffledOptions,
        correct: correctAnswer,
        supplement: item.Supplement,
      } as QuizData;
    });
    setQuizData(mappedData);
    setDataFetched(true);
  } catch (error) {
    console.error('クイズデータの取得中にエラーが発生しました:', error);
  }
};

export default fetchQuizData;
