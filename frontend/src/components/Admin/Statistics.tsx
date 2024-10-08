import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './css/AdminScreen.css'; // AdminScreen.cssをインポート
import './css/Statistics.css'; // Statistics.cssをインポート
import Ranking from './components/Ranking';
import Content from './components/Content';

const getHeaders = (jwt: string | null) => ({
  Authorization: `Bearer ${jwt}`,
  'Content-Type': 'application/json',
});

const Statistics: React.FC = () => {
  const [ranking, setRanking] = useState<any[]>([]);
  const jwt = sessionStorage.getItem('token'); // セッションストレージからJWTトークンを取得

  useEffect(() => {
    axios
      .get(`${process.env.REACT_APP_API_URL}/admins/analizedata/ranking`, {
        headers: getHeaders(jwt),
      })
      .then((response) => {
        setRanking(response.data.ranking);
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, [jwt]);

  return (
    <div className="container">
      <div className="header">
        <h2 className="admin-h2">統計と分析（社員の成績を確認）</h2>
      </div>
      <div className="content-container">
        <Ranking ranking={ranking} />
        <Content />
      </div>
    </div>
  );
};

export default Statistics;
