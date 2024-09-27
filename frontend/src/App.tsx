import { useState, useEffect, useCallback } from 'react';
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
import handleLogout from './hooks/utils/handleLogout'; // handleLogoutをインポート

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
  
  const quizData = useQuizData(isLoggedIn, todaysFinish, isAdmin, isSubmitAnsewr, setTodaysFinish);
  console.log('quizData:', quizData);

  const navigate = useNavigate();
  const token = sessionStorage.getItem('token'); // トークンを取得

  // トークンに基づいてログイン状態を設定
  useEffect(() => {
    if (token) {
      setIsLoggedIn(true);
      const admin = sessionStorage.getItem('admin') === 'true';
      const todaysFinish = sessionStorage.getItem('todaysFinish') === 'true';
      setIsAdmin(admin);
      setTodaysFinish(todaysFinish);
    } else {
      setIsLoggedIn(false);
    }
  }, [token]); // `token` が変更されたときにのみ実行


  useEffect(() => {
    if (!isLoggedIn) {
      navigate('/'); // トークンがない場合はログインページにリダイレクト
    }
  }, [isLoggedIn, navigate]);

  // useEffect(() => {
  //   if (!isAdmin) {
  //     navigate('/'); // トークンがない場合はログインページにリダイレクト
  //   }
  // }, [isAdmin, navigate]);


   // コンポーネントがマウントされたときにセッションから `currentQuestion` を復元する
   useEffect(() => {
    const todaysFinish = sessionStorage.getItem('todays_finish') === 'true';
    setTodaysFinish(todaysFinish);
    console.log('todaysFinish:', todaysFinish);

    const savedCurrentQuestion = sessionStorage.getItem('currentQuestion');
    console.log('savedCurrentQuestion:', savedCurrentQuestion);
    if (savedCurrentQuestion) {
      setCurrentQuestion(JSON.parse(savedCurrentQuestion));
    }
    const savedAnswers = sessionStorage.getItem('answers');
    console.log('savedAnswers:', savedAnswers);
    if (savedAnswers) {
      setAnswers(JSON.parse(savedAnswers));
    }
    const savedScore = sessionStorage.getItem('score');
  console.log('savedScore:', savedScore);
  if (savedScore) {
    setScore(JSON.parse(savedScore));
  }
    setNext(false);
  }, []);



  const handleBeforeUnload = useCallback((event: BeforeUnloadEvent) => {
    // submitAnswers関数を呼び出す条件をチェック
    if (isLoggedIn && !isAdmin && !todaysFinish && !showScore) {
        submitAnswers(answers, setIsSubmitAnsewr);
      }
  }, [answers]);

  useEffect(() => {
    window.addEventListener('beforeunload', handleBeforeUnload);

    return () => {
      window.removeEventListener('beforeunload', handleBeforeUnload);
    };
  }, [handleBeforeUnload]);
  
  

  return (
    <div className="quiz-container">
      {isLoggedIn ? (
        isAdmin ? ( // 管理者の場合
          <AdminScreen isAdmin={isAdmin} onLogout={() => handleLogout(setIsLoggedIn, setIsAdmin, setIsSubmitAnsewr, setShowScore, navigate)} /> // 管理者用画面を表示
        ) : (
          todaysFinish ? ( // 日のノルマが達成された場合
            <div className="quota-met">
              <h1>本日の受験は終了しました</h1><h1>┌┤´д`├┐ﾀﾞﾙ～</h1>
              <button className="logout-button" onClick={() => handleLogout(setIsLoggedIn, setIsAdmin, setIsSubmitAnsewr, setShowScore, navigate)}>ログアウト</button>
            </div>
          ) : (
            showScore ? (
              <ScoreSection 
                score={score} 
                answers={answers} 
                isSubmitAnsewr={isSubmitAnsewr} 
                handleLogout={handleLogout} 
                setIsLoggedIn={setIsLoggedIn}
                setIsAdmin={setIsAdmin}
                setIsSubmitAnsewr={setIsSubmitAnsewr}
                setShowScore={setShowScore}
                navigate={navigate}
              />
             ) : (
              quizData.length > 0 ? (
                <Quiz
                  currentQuestion={currentQuestion}
                  quizData={quizData}
                  next={next}
                  feedback={feedback}
                  handleAnswer={(selectedAnswer: Option) => handleAnswer(
                    selectedAnswer, quizData, currentQuestion, setCurrentQuestion, setScore, setFeedback, setAnswers, setNext, submitAnswers, setShowScore, setIsSubmitAnsewr, answers
                  )}
                  goToNextQuestion={() => goToNextQuestion(
                    setNext, 
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
        <Route path="/admin" element={<AdminScreen isAdmin={isAdmin} onLogout={() => handleLogout(setIsLoggedIn, setIsAdmin, setIsSubmitAnsewr, setShowScore, navigate)} />} />
        <Route path="/add-question" element={<AddQuestion />} />
        <Route path="/edit-question" element={<UpdateDeleteQuestion />} />
        <Route path="/user-management" element={<UserManagement />} />
        <Route path="/statistics" element={<Statistics />} />
      </Routes>
    </div>
  );
}

export default App;
