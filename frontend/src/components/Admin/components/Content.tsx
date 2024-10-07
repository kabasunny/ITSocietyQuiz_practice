import React, { useState, useEffect } from 'react';
import '../css/AdminScreen.css'; // AdminScreen.cssをインポート
import '../css/Statistics.css'; // Statistics.cssをインポート

const Content: React.FC = () => {
  const [image, setImage] = useState<string | null>(null);

  useEffect(() => {
    fetch('http://127.0.0.1:5001/api/visualize', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        employee_id: ['EMP1', 'EMP2', 'EMP3', 'EMP4', 'EMP5', 'EMP6', 'EMP7'],
        score_indicator: [15, 20, 25, 5, 10, 15, 20],
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.image) {
          setImage(data.image);
        }
      })
      .catch((error) => {
        console.error('Error fetching the image:', error);
      });
  }, []);

  return (
    <div className="content">
      <h2 className="admin-h2">傾向グラフ</h2>
      {image && (
        <img
          src={`data:image/png;base64,${image}`}
          alt="Employee Scores"
          className="graph-image"
        />
      )}
    </div>
  );
};

export default Content;
