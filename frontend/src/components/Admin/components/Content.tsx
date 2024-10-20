import React, { useState, useEffect } from 'react';
import '../css/AdminScreen.css'; // AdminScreen.cssをインポート
import '../css/Statistics.css'; // Statistics.cssをインポート

interface ContentProps {
  graphImage: string | null;
  title: string;
}

const Content: React.FC<ContentProps> = ({ graphImage, title }) => {
  const [image, setImage] = useState<string | null>(graphImage);

  useEffect(() => {
    setImage(graphImage);
  }, [graphImage]);

  return (
    <div className="content">
      <h2 className="admin-h2">{title}</h2>
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
