import React from 'react';
import './ScoreSection.css';

interface Answer {
  question: string;
  answer: string;
  correct: boolean;
}

interface ScoreSectionProps {
  score: number;
  answers: Answer[];
}

const ScoreSection: React.FC<ScoreSectionProps> = ({ score, answers }) => {
  return (
    <div className='score-section'>
      <h1>スコア</h1>
      <h2 className="final-score">
        {score}問正解 ( 全{answers.length}問中 )
      </h2>
      <table className="answer-table">
        <thead>
          <tr>
            <td>出題内容</td>
            <td>あなたの解答</td>
            <td>判定</td>
          </tr>
        </thead>
        <tbody>
          {answers.map((item, index) => (
            <tr key={index} className={item.correct ? "correct" : "wrong"}>
              <td>{item.question}</td>
              <td>{item.answer}</td>
              <td>{item.correct ? "○" : "×"}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default ScoreSection;
