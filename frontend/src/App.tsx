import { useState } from 'react';
import './App.css';
import Quiz from './components/Quiz';
import ScoreSection from './components/ScoreSection';
import Login from './components/Login';
import useQuizData from './hooks/useQuizData';
import { Answer, Option } from './types';
import handleLogin from './components/utils/handleLogin';
import handleAnswer from './components/utils/handleAnswer';
import goToNextQuestion from './components/utils/goToNextQuestion';
import handleResultsSubmit from './components/utils/handleResultsSubmit';

function App() {
  const [currentQuestion, setCurrentQuestion] = useState<number>(0);
  const [next, setNext] = useState<boolean>(false);
  const [answers, setAnswers] = useState<Answer[]>([]);
  const [score, setScore] = useState<number>(0);
  const [feedback, setFeedback] = useState<string | null>(null);
  const [showScore, setShowScore] = useState<boolean>(false);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const quizData = useQuizData();

  return (
    <div className="quiz-container">
      {isLoggedIn ? ( // ログイン済みのとき
        showScore ? ( // クイズを解き終えている
          <ScoreSection score={score} answers={answers} /> // スコアを表示
        ) : ( // クイズ解き終えていない
          quizData.length > 0 ? ( // クイズデータ取得済み
            <Quiz // クイズ出題コンポーネント
              currentQuestion={currentQuestion}
              quizData={quizData}
              next={next}
              feedback={feedback}
              handleAnswer={(answer: Option) => handleAnswer(
                answer, quizData, currentQuestion, setScore, setFeedback, setAnswers, setNext, answers
              )}
              goToNextQuestion={() => goToNextQuestion(
                currentQuestion, setCurrentQuestion, quizData, () => handleResultsSubmit(answers), setShowScore, setNext
              )}
            />
          ) : ( // クイズデータがないとき、待ち画面
            <div className="loading"> 
              <p>Loading...</p>
              <p>⁽⁽*( ᐖ )*⁾⁾ ₍₍*( ᐛ )*₎₎</p>
            </div>
          )
        )
      ) : (
        <Login onLogin={(loginOK: boolean) => handleLogin(loginOK, setIsLoggedIn)} />
      )}
    </div>
  );
}

export default App;
