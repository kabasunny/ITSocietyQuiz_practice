import { useState, useEffect } from 'react';
import { Route, Routes } from 'react-router-dom';
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
import AdminScreen from './components/Admin/AdminScreen'; // 管理者用画面をインポート
import AddQuestion from './components/Admin/AddQuestion';
import UpdateDeleteQuestion from './components/Admin/UpdateDeleteQuestion';
import UserManagement from './components/Admin/UserManagement';
import Statistics from './components/Admin/Statistics';

function App() {
  const [currentQuestion, setCurrentQuestion] = useState<number>(0);
  const [next, setNext] = useState<boolean>(false);
  const [answers, setAnswers] = useState<Answer[]>([]);
  const [score, setScore] = useState<number>(0);
  const [feedback, setFeedback] = useState<string | null>(null);
  const [showScore, setShowScore] = useState<boolean>(false);
  
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false);
  const [isAdmin, setIsAdmin] = useState<boolean>(false); // 管理者フラグの状態管理
    
  const [isSubmitAnsewr, setIsSubmitAnsewr] = useState<boolean>(false);
  const quizData = useQuizData(isLoggedIn); // バグ修正、ログイン済みでないとリクエストできない

  return (
    <div className="quiz-container">
      {isLoggedIn ? (
        isAdmin ? ( // 管理者の場合
          <AdminScreen isAdmin={isAdmin} /> // 管理者用画面を表示
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
      ) : (
        <Login onLogin={(loginOK: boolean, isAdmin: boolean) => handleLogin(loginOK, setIsLoggedIn, setIsAdmin, isAdmin)} />
      )}
      <Routes>
        <Route path="/admin" element={<AdminScreen isAdmin={isAdmin} />} />
        <Route path="/add-question" element={<AddQuestion />} />
        <Route path="/update-delete-question" element={<UpdateDeleteQuestion />} />
        <Route path="/user-management" element={<UserManagement />} />
        <Route path="/statistics" element={<Statistics />} />
      </Routes>
    </div>
  );
}

export default App;
