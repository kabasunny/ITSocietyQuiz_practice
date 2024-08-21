import { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';
import Quiz from './components/Quiz';
import ScoreSection from './components/ScoreSection';

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
  const [quizData, setQuizData] = useState<QuizData[]>([]);

  useEffect(() => {
    const fetchQuizData = async () => {
      try {
        const response = await axios.get('http://localhost:8080/quiz_data');
        const mappedData = response.data.data.map((item: any) => ({
          question: item.Question,
          options: item.Options,
          correct: item.Correct,
          supplement: item.Supplement,
        }));
        setQuizData(mappedData);
      } catch (error) {
        console.error('クイズデータの取得中にエラーが発生しました:', error);
      }
    };

    fetchQuizData();
  }, []);

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
      {showScore ? (
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
          <div>Loading...</div>
        )
      )}
    </div>
  );
}

export default App;


// import { useState } from 'react';
// import './App.css';
// import Quiz from './components/Quiz';
// import ScoreSection from './components/ScoreSection';
// import { quizData } from './data/quizData';

// interface Answer {
//   question: string;
//   answer: string;
//   correct: boolean;
// }

// function App() {
//   const [currentQuestion, setCurrentQuestion] = useState<number>(0);
//   const [next, setNext] = useState<boolean>(false);
//   const [answers, setAnswers] = useState<Answer[]>([]);
//   const [score, setScore] = useState<number>(0);
//   const [feedback, setFeedback] = useState<string | null>(null);
//   const [showScore, setShowScore] = useState<boolean>(false);

//   const handleAnswer = (answer: string) => {
//     const newAnswer: Answer = {
//       question: quizData[currentQuestion].question,
//       answer: answer,
//       correct: answer === quizData[currentQuestion].correct,
//     };

//     if (newAnswer.correct) {
//       setScore((prevScore) => prevScore + 1);
//       setFeedback("○");
//     } else {
//       setFeedback("×");
//     }
//     setAnswers([...answers, newAnswer]);
//     setNext(true);
//   };

//   const goToNextQuestion = () => {
//     const nextQuestion = currentQuestion + 1;
//     if (nextQuestion < quizData.length) {
//       setCurrentQuestion(nextQuestion);
//     } else {
//       setShowScore(true);
//     }
//     setNext(false);
//   };

//   return (
//     <div className="quiz-container">
//       {showScore ? (
//         <ScoreSection score={score} answers={answers} />
//       ) : (
//         <Quiz
//           currentQuestion={currentQuestion}
//           quizData={quizData}
//           next={next}
//           feedback={feedback}
//           handleAnswer={handleAnswer}
//           goToNextQuestion={goToNextQuestion}
//         />
//       )}
//     </div>
//   );
// }

// export default App;
