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

  const onLogin = (loginOK: boolean) => {
    setIsLoggedIn(loginOK);
  };

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
        <Login onLogin={onLogin} />
      )}
    </div>
  );
}

export default App;
