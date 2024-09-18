import { useState, useEffect } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import './App.css';
import Quiz from './components/Quiz';
import ScoreSection from './components/ScoreSection';
import Login from './components/Login';
import useQuizData from './hooks/useQuizData';
import { Answer, Option } from './types';
import handleLogin from './components/utils/handleLogin';
import handleAnswer from './components/utils/handleAnswer';
import goToNextQuestion from './components/utils/goToNextQuestion';
import submitAnswers from './components/utils/submitAnswers';
import AdminScreen from './components/admin/AdminScreen'; // 管理者用画面をインポート
import AddQuestion from './components/admin/AddQuestion';
import UpdateDeleteQuestion from './components/admin/EditQuestion';
import UserManagement from './components/admin/UserManagement';
import Statistics from './components/admin/Statistics';
import handleAdminLogout from './hooks/utils/handleLogout'; // handleAdminLogoutをインポート

function App() {
  const [currentQuestion, setCurrentQuestion] = useState<number>(0);
  const [next, setNext] = useState<boolean>(false);
  const [answers, setAnswers] = useState<Answer[]>([]);
  const [score, setScore] = useState<number>(0);
  const [feedback, setFeedback] = useState<string | null>(null);
  const [showScore, setShowScore] = useState<boolean>(false);
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false); // ログイン状態の管理
  const [isAdmin, setIsAdmin] = useState<boolean>(false); // 管理者フラグの状態管理

  const [isSubmitAnsewr, setIsSubmitAnsewr] = useState<boolean>(false);
  const [todaysFinish, setTodaysFinish] = useState<boolean>(false); // 日のノルマフラグの状態管理

  const navigate = useNavigate();
  const token = sessionStorage.getItem('token'); // トークンを取得

  // トークンに基づいてログイン状態を設定
  useEffect(() => {
    if (token) {
      setIsLoggedIn(true);
      const admin = sessionStorage.getItem('admin') === 'true';
      const todaysCount = sessionStorage.getItem('todays_count') === 'true';
      setIsAdmin(admin);
      setTodaysFinish(todaysCount);
    } else {
      setIsLoggedIn(false);
    }
  }, [token]); // `token` が変更されたときにのみ実行

  // 常に useQuizData を呼び出すが、結果を条件付きで使用する
  const quizData = useQuizData(isLoggedIn, todaysFinish, isAdmin, setTodaysFinish);

  useEffect(() => {
    if (!isLoggedIn) {
      navigate('/'); // トークンがない場合はログインページにリダイレクト
    }
  }, [isLoggedIn, navigate]);

  return (
    <div className="quiz-container">
      {isLoggedIn ? (
        isAdmin ? ( // 管理者の場合
          <AdminScreen isAdmin={isAdmin} onAdminLogout={() => handleAdminLogout(setIsLoggedIn, setIsAdmin, navigate)} /> // 管理者用画面を表示
        ) : (
          todaysFinish ? ( // 日のノルマが達成された場合
            <div className="quota-met">
              <h1>本日の受験は終了しました</h1><h1>┌┤´д`├┐ﾀﾞﾙ～</h1>
              <button className="logout-button" onClick={() => handleAdminLogout(setIsLoggedIn, setIsAdmin, navigate)}>ログアウト</button>
            </div>
          ) : (
            showScore ? (
              <ScoreSection score={score} answers={answers} isSubmitAnsewr={isSubmitAnsewr} />
            ) : (
              quizData.length > 0 ? (
                <Quiz
                  currentQuestion={currentQuestion}
                  quizData={quizData}
                  next={next}
                  feedback={feedback}
                  handleAnswer={(selectedAnswer: Option) => handleAnswer(
                    selectedAnswer, quizData, currentQuestion, setScore, setFeedback, setAnswers, setNext, answers
                  )}
                  goToNextQuestion={() => goToNextQuestion(
                    currentQuestion, setCurrentQuestion, quizData, answers, submitAnswers, setShowScore, setNext, setIsSubmitAnsewr
                  )}
                />
              ) : (
                <div className="loading">
                  <p>Loading...</p>
                  <p>⁽⁽*( ᐖ )*⁾⁾ ₍₍*( ᐛ )*₎₎</p>
                </div>
              )
            )
          )
        )
      ) : (
        <Login onLogin={(loginOK: boolean, isAdmin: boolean) => handleLogin(loginOK, setIsLoggedIn, setIsAdmin, isAdmin)} />
      )}
      <Routes>
        <Route path="/admin" element={<AdminScreen isAdmin={isAdmin} onAdminLogout={() => handleAdminLogout(setIsLoggedIn, setIsAdmin, navigate)} />} />
        <Route path="/add-question" element={<AddQuestion />} />
        <Route path="/edit-question" element={<UpdateDeleteQuestion />} />
        <Route path="/user-management" element={<UserManagement />} />
        <Route path="/statistics" element={<Statistics />} />
      </Routes>
    </div>
  );
}

export default App;
