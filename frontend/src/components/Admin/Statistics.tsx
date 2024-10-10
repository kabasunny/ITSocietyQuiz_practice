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
  const [graphData, setGraphData] = useState<string | null>(null); // 平均のグラフデータの状態を追加
  const [empGraphData, setEmpGraphData] = useState<string | null>(null); // 個人のグラフデータの状態を追加
  const [empID, setEmpID] = useState<string | null>(null); // empIDの状態を追加
  const jwt = sessionStorage.getItem('token'); // セッションストレージからJWTトークンを取得

  useEffect(() => {
    axios
      .get(
        `${process.env.REACT_APP_API_URL}/admins/analizedata/initialdata/AVG100`,
        {
          headers: getHeaders(jwt),
        }
      )
      .then((response) => {
        setRanking(response.data.ranking);
        setGraphData(response.data.graphData.image); // グラフデータを設定
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, [jwt]);

  useEffect(() => {
    axios
      .get(
        `${process.env.REACT_APP_API_URL}/admins/analizedata/initialdata/${empID}`,
        {
          headers: getHeaders(jwt),
        }
      )
      .then((response) => {
        setRanking(response.data.ranking);
        setEmpGraphData(response.data.graphData.image); // グラフデータを設定
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, [empID]);

  return (
    <div className="container">
      <div className="header">
        <h2 className="admin-h2">統計と分析（社員の成績を確認）</h2>
      </div>
      <div className="content-container">
        <div className="ranking">
          <Ranking ranking={ranking} setEmpID={setEmpID} />
        </div>
        <div className="content">
          {empID ? (
            <Content
              graphImage={empGraphData}
              title={`${empID}(個人)の成績グラフ`}
            />
          ) : null}
          <Content graphImage={graphData} title="全体平均の傾向グラフ" />
        </div>
      </div>
    </div>
  );
};

export default Statistics;
