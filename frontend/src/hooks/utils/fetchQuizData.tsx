import axios from 'axios';
import shuffleArray from './shuffleArray';
import { QuizData, Option } from '../../types';

interface ApiQuizItem { // APIから受け取る型、QuizData型へ変換用
  id: number;
  question: string;
  options: string[];
  supplement: string;
  difficulty: number;
}

const fetchQuizData = async (setQuizData: (quizData: QuizData[]) => void, setDataFetched: (fetched: boolean) => void) => {
  try {
    const response = await axios.get('http://localhost:8082/questions/oneday');
    console.log('APIレスポンス:', response.data); // レスポンスデータをログに出力
    // ここでレスポンスは、Go言語のdtoで設計した (modelの構造になっていたため)
    // type QuizData struct {
    // 	Question   string   `json:"question"`
    // 	Options    []string `json:"options"`
    // 	Supplement string   `json:"supplement"`
    // 	Difficulty int      `json:"difficulty"` ( ← いらないかも)
    // } となるようにバック側のQuestionsServiceを修正した
    const mappedData: QuizData[] = response.data.data.map((item: ApiQuizItem) => {
      console.log('マッピング前のアイテム:', item); // マッピング前のデータをログに出力
      const correctAnswer: string = item.options[0]; // レスポンスのOptionsキーの最初の配列が正解
      const optionsWithIndex: Option[] = item.options.map((option: string, index: number) => ({
          text: option,
          index: index
      })); // オブジェクト.array.map((element, index, array)
      const shuffledOptions: Option[] = shuffleArray([...optionsWithIndex]);
      const quizItem: QuizData = {
        id: item.id,
        question: item.question,
        options: shuffledOptions,
        correct: correctAnswer,
        supplement: item.supplement,
      };
      console.log('マッピング後のアイテム:', quizItem); // マッピング後のデータをログに出力
      return quizItem;
    });

  
    setQuizData(mappedData);
    setDataFetched(true);
  } catch (error) {
    console.error('クイズデータの取得中にエラーが発生しました:', error);
  }
};

export default fetchQuizData;
