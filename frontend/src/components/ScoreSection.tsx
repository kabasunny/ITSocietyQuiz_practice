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
  onEnd: () => void; // 終了ボタンのためのプロップ
}

const ScoreSection: React.FC<ScoreSectionProps> = ({ score, answers, onEnd }) => {
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
              <td>{item.answer}</td>
              <td className="judgement">{item.correct ? "○" : "×"}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <h2 className="final-score">
        スコア : {score}問正解 ( 全{answers.length}問中 )
      </h2>
      {/* 終了ボタン */}
      <button onClick={onEnd} className="end-button">
        終了
      </button>
    </div>
  );
}

export default ScoreSection;
