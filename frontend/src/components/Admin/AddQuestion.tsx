import React, { useState, useRef } from 'react';
import axios from 'axios';
import './AdminScreen.css'; // CSSファイルをインポート

const UploadCSV: React.FC = () => {
  const [file, setFile] = useState<File | null>(null);
  const [message, setMessage] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files) {
      setFile(event.target.files[0]);
    }
  };

  const handleUpload = async () => {
    if (!file) return;

    const formData = new FormData();
    formData.append('file', file);

    try {
      const response = await axios.post('http://localhost:8082/admins/import_csv', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      if (response.status === 200) {
        setMessage('ファイルが正常にアップロードされました。');
      } else {
        setMessage('ファイルのアップロードに失敗しました。');
      }
      setFile(null); // ファイル選択をリセット
      if (fileInputRef.current) {
        fileInputRef.current.value = ''; // ファイル入力フィールドをリセット
      }
    } catch (error) {
      console.error('Error uploading file:', error);
      setMessage('ファイルのアップロード中にエラーが発生しました。');
    }
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">CSVファイルのアップロード</h2>
      <div className="button-group">
        <input type="file" accept=".csv" onChange={handleFileChange} ref={fileInputRef} className="file-input" id="file-upload" />
        <label htmlFor="file-upload" className="upload-csv-button">ファイルを選択</label>
        <button onClick={handleUpload} className="upload-button">アップロード</button>

      </div>
      {file && <p>選択されたファイル: {file.name}</p>}
      {message && <p>{message}</p>}
    </div>
  );
};

export default UploadCSV;
