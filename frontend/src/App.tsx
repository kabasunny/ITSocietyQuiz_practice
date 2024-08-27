import { useState } from 'react';
import './App.css';
import Quiz from './components/Quiz';
import ScoreSection from './components/ScoreSection';
import Login from './components/Login';
import useQuizData from './hooks/useQuizData';

interface Answer {
  question: string;
  answer: string;
  correct: boolean;
}

function App() {
  const [currentQuestion, setCurrentQuestion] = useState<number>(0);
  const [next, setNext] = useState<boolean>(false);
  const [answers, setAnswers] = useState<Answer[]>([]);
  const [score, setScore] = useState<number>(0);
  const [feedback, setFeedback] = useState<string | null>(null);
  const [showScore, setShowScore] = useState<boolean>(false);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const quizData = useQuizData();

  // ログインが成功したかどうかを受け取り、その結果に基づいてログイン状態
  const onLogin = (loginOK: boolean) => {
    setIsLoggedIn(loginOK);
  };

  // ユーザーがクイズの回答を選択したときに呼び出され、回答の正誤を判定し、スコアやフィードバックを更新
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

  // 次のクイズに遷移する処理
  const goToNextQuestion = () => {
    const nextQuestion = currentQuestion + 1;
    if (nextQuestion < quizData.length) {
      setCurrentQuestion(nextQuestion);
    } else {
      handleResultsSubmit(); // 最終画面へ遷移する前に結果を送信
      setShowScore(true);
    }
    setNext(false);
  };

  // 結果送信の処理
  const handleResultsSubmit = () => {
    // 実装
    console.log('結果が送信されました');
  };

  // 終了ボタンがクリックされたときの処理
  const handleEnd = () => {
    // 実装
    console.log('クイズが終了しました');
  };

  return (
    <div className="quiz-container">
      {isLoggedIn ? ( // ログインが認証されている
        showScore ? ( // クイズ全問終了しているので、結果画面を表示し、結果を送信
          <ScoreSection score={score} answers={answers} onEnd={handleEnd} />
        ) : ( // 未回答クイズ問題有り
          quizData.length > 0 ? ( // クイズデータが取得されている
            <Quiz
              currentQuestion={currentQuestion}
              quizData={quizData}
              next={next}
              feedback={feedback}
              handleAnswer={handleAnswer}
              goToNextQuestion={goToNextQuestion}
            />
          ) : ( // クイズデータが取得できてないので、APIのレスポンス待ちのてい
            <div className="loading">
              <p>Loading...</p>
              <p>⁽⁽*( ᐖ )*⁾⁾ ₍₍*( ᐛ )*₎₎</p>
            </div>
          )
        )
      ) : ( // ログインが認証されていないので、ログイン画面へ遷移
        <Login onLogin={onLogin} />
      )}
    </div>
  );
}

export default App;
