import React, { useState } from 'react';
import './ScoreSection.css';
import { Answer } from '../types';

interface ScoreSectionProps {
  score: number;
  answers: Answer[];
}

const ScoreSection: React.FC<ScoreSectionProps> = ({ score, answers }) => {
  const [isEnded, setIsEnded] = useState<boolean>(false); // 終了状態を管理

  const handleEndClick = () => {
    setIsEnded(true);
  };

  return (
    <div className='score-section'>
      {/* 環境変数からタイトルを読み込む */}
      <h1>{process.env.REACT_APP_TITLE}</h1>
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
              <td>{item.answer.text}</td>
              <td className="judgement">{item.correct ? "○" : "×"}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <h2 className="final-score">
        スコア : {score}問正解 ( 全{answers.length}問中 )
      </h2>
      {isEnded ? (
        <h1 className="end-message">本日の学習は終了です٩( 'ω' )و</h1>
      ) : (
        <button onClick={handleEndClick} className="end-button">
          終了
        </button>
      )}
    </div>
  );
}

export default ScoreSection;
