import React, { useState } from 'react';
import './ScoreSection.css';
import { Answer } from '../types';
import { NavigateFunction } from 'react-router-dom';

interface ScoreSectionProps {
  score: number;
  answers: Answer[];
  isSubmitAnsewr: boolean;
  handleLogout: (setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>, setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>, setIsSubmitAnsewr: React.Dispatch<React.SetStateAction<boolean>>, setShowScore: React.Dispatch<React.SetStateAction<boolean>>, navigate: NavigateFunction) => void;
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>;
  setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>;
  setIsSubmitAnsewr: React.Dispatch<React.SetStateAction<boolean>>;
  setShowScore: React.Dispatch<React.SetStateAction<boolean>>;
  navigate: NavigateFunction;
}

const ScoreSection: React.FC<ScoreSectionProps> = ({ score, answers, isSubmitAnsewr, handleLogout, setIsLoggedIn, setIsAdmin, setIsSubmitAnsewr, setShowScore, navigate }) => {
  const [isEnded, setIsEnded] = useState<boolean>(false); // 終了状態を管理

  const handleEndClick = () => {
    if (isSubmitAnsewr) {
      setIsEnded(true);
      setTimeout(() => {
        handleLogout(setIsLoggedIn, setIsAdmin, setIsSubmitAnsewr, setShowScore, navigate); // 必要な引数を渡す
      }, 2000); // 2秒待つ
    } else {
      alert('データ送信に失敗しました');
    }
    sessionStorage.setItem('currentQuestion', JSON.stringify(0)); // 終了ボタンを押すとリセット
    sessionStorage.setItem('answers', JSON.stringify(null)); 
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
          {answers.map((answer, index) => (
            <tr key={index} className={answer.correct ? "correct" : "wrong"}>
              <td>{answer.question}</td>
              <td>{answer.answer_text}</td>
              <td className="judgement">{answer.correct ? "○" : "×"}</td>
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
          終了<br/>
          & ログアウト
        </button>
      )}
    </div>
  );
}

export default ScoreSection;
