import React from 'react';
import '../css/AdminScreen.css'; // AdminScreen.cssをインポート
import '../css/Statistics.css'; // Statistics.cssをインポート

interface RankingProps {
  ranking: any[];
}

const Ranking: React.FC<RankingProps> = ({ ranking }) => {
  return (
    <div className="ranking">
      <h2 className="admin-h2">ランキング</h2>
      <div className="ranking-table-container">
        <table className="ranking-table">
          <thead>
            <tr>
              <th>順位</th>
              <th>社員ID</th>
              <th>回答数</th>
              <th>正答率</th>
              <th>進捗数</th>
              <th>指標値</th>
            </tr>
          </thead>
          <tbody>
            {ranking.map((user) => (
              <tr key={user.empId}>
                <td>{user.rank}</td>
                <td>{user.empId}</td>
                <td>{user.totalQuestions}</td>
                <td>{user.correctAnswerRate.toFixed(1)}%</td>
                <td>{user.currentQID}</td>
                <td>{user.performanceIndicator.toFixed(1)}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Ranking;
