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
  const handleLogin = (loginOK: boolean) => {
    setIsLoggedIn(loginOK);
  };

  // ユーザーがクイズの回答を選択したときに呼び出され、回答の正誤を判定し、スコアやフィードバックを更新
  const handleAnswer = (answer: string) => {
    const newAnswer: Answer = {
      question: quizData[currentQuestion].question,
      answer: answer,
      correct: answer === quizData[currentQuestion].correct,
    };

    if (newAnswer.correct) { // 正解なら、スコアに1を足し、フィードバックに○
      setScore((prevScore) => prevScore + 1);
      setFeedback("○");
    } else { // 不正解なら、フィードバックに×
      setFeedback("×");
    }
    setAnswers([...answers, newAnswer]);
    setNext(true); // 次の質問に遷移するフラグ(結果を表示して、次のクイズへ)
  };

  // 次のクイズに遷移する処理
  const goToNextQuestion = () => {
    const nextQuestion = currentQuestion + 1;
    if (nextQuestion < quizData.length) { //次のクイズ追番が、クイズの数より小さい場合
      setCurrentQuestion(nextQuestion);
    } else { //次のクイズ追番が、クイズの数以上の場合
      handleResultsSubmit(); // 最終画面へ遷移する前に結果を送信
      setShowScore(true); // 最終画面へ遷移する許可
    }
    setNext(false); // 次の質問に遷移するフラグ(今のクイズが表示される)
  };

  // 結果送信の処理
  const handleResultsSubmit = () => {
    alert('結果が送信されました');
  };


  return (
    <div className="quiz-container">
      {isLoggedIn ? ( // ログイン認証されている
        showScore ? ( // クイズ全問終了しているので、結果画面を表示し、結果を送信
          <ScoreSection score={score} answers={answers} />
        ) : ( // 未回答クイズ問題有り
          quizData.length > 0 ? ( // クイズデータが取得されている状態
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
      ) : ( // ログイン認証されていないので、ログイン画面へ遷移
        <Login onLogin={handleLogin} />
      )}
    </div>
  );
}

export default App;
