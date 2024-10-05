import React, { useState, useEffect } from 'react';
import './css/AdminScreen.css'; // CSSファイルをインポート

const Statistics: React.FC = () => {
  const [image, setImage] = useState<string | null>(null);  // 画像データを保存するための状態を定義

  useEffect(() => {
    // FlaskバックエンドのエンドポイントにPOSTリクエストを送信
    // とりあえずテストで、GoのAPIを介さず、pythonのFlaskサーバに直接
    // 将来的には、Goのルーティングに従う
    fetch('http://127.0.0.1:5001/api/visualize', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        employee_id: ['EMP1', 'EMP2', 'EMP3', 'EMP4', 'EMP5', 'EMP6', 'EMP7'],  // 送信するデータ。実際はGoでDBからデータを取得する
        score_indicator: [15, 20, 25, , 10, 15, 20],
      }),
    })
      .then((response) => response.json())  // レスポンスをJSON形式に変換
      .then((data) => {
        if (data.image) {
          setImage(data.image);  // 取得した画像データを状態に保存
        }
      })
      .catch((error) => {
        console.error('Error fetching the image:', error);  // エラーメッセージをコンソールに表示
      });
  }, []);  // コンポーネントがマウントされたときに実行

  return (
    <div>
      <h2 className="admin-h2">統計と分析（社員の成績を確認）</h2>
      {image && <img src={`data:image/png;base64,${image}`} alt="Employee Scores" />}
      {/*画像データが存在する場合、画像を表示 */}
    </div>
  );
};

export default Statistics;
