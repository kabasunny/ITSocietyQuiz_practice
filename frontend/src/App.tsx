import { useState, useEffect } from 'react';
// import axios from 'axios';
import './App.css';
import Quiz from './components/Quiz';
import ScoreSection from './components/ScoreSection';
import { quizData } from './data/quizData';
import Login from './components/Login';

interface Answer {
  question: string;
  answer: string;
  correct: boolean;
}

interface QuizData {
  question: string;
  options: string[];
  correct: string;
  supplement: string;
}

function App() {
  const [currentQuestion, setCurrentQuestion] = useState<number>(0);
  const [next, setNext] = useState<boolean>(false);
  const [answers, setAnswers] = useState<Answer[]>([]);
  const [score, setScore] = useState<number>(0);
  const [feedback, setFeedback] = useState<string | null>(null);
  const [showScore, setShowScore] = useState<boolean>(false);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [username, setUsername] = useState<string>('');
  // const [quizData, setQuizData] = useState<QuizData[]>([]); // WebAPIから取得
  // const [dataFetched, setDataFetched] = useState<boolean>(false); // データ取得フラグ

  const handleLogin = (username: string) => {
    setIsLoggedIn(true);
    setUsername(username);
  };

  // WebAPIを使用する場合
  // useEffect(() => {
  //   const fetchQuizData = async () => {
  //     try {
  //       const response = await axios.get('http://localhost:8082/quiz_data');
  //       const mappedData = response.data.data.map((item: any) => ({
  //         question: item.Question,
  //         options: item.Options,
  //         correct: item.Correct,
  //         supplement: item.Supplement,
  //       }));
  //       setQuizData(mappedData);
  //       setDataFetched(true); // データ取得フラグを更新
  //     } catch (error) {
  //       console.error('クイズデータの取得中にエラーが発生しました:', error);
  //     }
  //   };

  //   fetchQuizData();

  //   const interval = setInterval(() => {
  //     if (!dataFetched) {
  //       fetchQuizData();
  //     } else {
  //       clearInterval(interval); // データが取得されたらインターバルをクリア
  //     }
  //   }, 10000); // 10秒ごとにデータを取得

  //   return () => clearInterval(interval); // コンポーネントがアンマウントされたときにインターバルをクリア
  // }, [dataFetched]);

  const handleAnswer = (answer: string) => {
    const newAnswer: Answer = {
      question: quizData[currentQuestion].question,
      answer: answer,
      correct: answer === quizData[currentQuestion].correct,
    };

    if (newAnswer.correct) {
      setScore((prevScore) => prevScore + 1);
      setFeedback("○");
    } else {
      setFeedback("×");
    }
    setAnswers([...answers, newAnswer]);
    setNext(true);
  };

  const goToNextQuestion = () => {
    const nextQuestion = currentQuestion + 1;
    if (nextQuestion < quizData.length) {
      setCurrentQuestion(nextQuestion);
    } else {
      setShowScore(true);
    }
    setNext(false);
  };

  return (
    <div className="quiz-container">
      {isLoggedIn ? (
        showScore ? (
          <ScoreSection score={score} answers={answers} />
        ) : (
          quizData.length > 0 ? (
            <Quiz
              currentQuestion={currentQuestion}
              quizData={quizData}
              next={next}
              feedback={feedback}
              handleAnswer={handleAnswer}
              goToNextQuestion={goToNextQuestion}
            />
          ) : (
            <div className="loading">
              <p>Loading...</p>
              <p>⁽⁽*( ᐖ )*⁾⁾ ₍₍*( ᐛ )*₎₎</p>
            </div>
          )
        )
      ) : (
        <Login onLogin={handleLogin} />
      )}
    </div>
  );
}

export default App;
