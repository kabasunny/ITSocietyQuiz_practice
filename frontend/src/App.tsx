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
              handleAnswer={(answer: Option) => handleAnswer(
                answer, quizData, currentQuestion, setScore, setFeedback, setAnswers, setNext, answers
              )}
              goToNextQuestion={() => goToNextQuestion(
                currentQuestion, setCurrentQuestion, quizData, () => handleResultsSubmit(answers), setShowScore, setNext
              )}
            />
          ) : (
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
