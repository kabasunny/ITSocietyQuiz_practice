import React from 'react';
import './Quiz.css';
import { Option, QuizProps } from '../types';

const Quiz: React.FC<QuizProps> = ({ currentQuestion, quizData, next, feedback, handleAnswer, goToNextQuestion }) => {
  return (
    <div className='question-section'>
      {/* 環境変数からタイトルを読み込む */}
      <h1>{process.env.REACT_APP_TITLE}</h1>
      <h1>{currentQuestion + 1} 問目 (  全{quizData.length}問中 )</h1>
      <h2>{quizData[currentQuestion].question}</h2>

      {next ? ( // クイズを解き終えている
        <div className='feedback-section'>
          <h2 className='large-feedback'>{feedback}</h2>
          <h3>解答: {quizData[currentQuestion].correct}</h3>
          <p>補足：{quizData[currentQuestion].supplement}</p>
          <button onClick={goToNextQuestion}>次の問題へ</button>
          <p><br />60分以内に全問解答できない場合は、全問不正解になります。</p>
        </div>
      ) : ( // これからクイズを解く
        <div className='answer-section'>
          <>
          {quizData[currentQuestion].options.map((item, index) => (
            
            <button 
              key={index}
              onClick={() => handleAnswer(item)}
              className={`quiz-option-button option-${item.index}`} // cssでindexに対応した装飾
            >
              {item.text}
            </button>
          ))}
            <p><br />60分以内に全問解答できない場合は、全問不正解になります。</p>
          </>
        </div>
      )}
    </div>
  );
}

export default Quiz;
