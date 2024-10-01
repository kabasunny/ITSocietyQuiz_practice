import React, { useState, useRef, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import './css/AdminScreen.css';
import './css/AddQuestion.css';

const UploadCSV: React.FC = () => {
  const [file, setFile] = useState<File | null>(null);
  const [message, setMessage] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null); // コンポーネントが再レンダリングされても参照が保持
  // const navigate = useNavigate();
  // const isAdmin = sessionStorage.getItem('admin') === 'true';

  // useEffect(() => {
  //   if (!isAdmin) {
  //     navigate('/'); // 管理者でない場合はホームページにリダイレクト
  //   }
  // }, [isAdmin, navigate]);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    // ファイル入力フィールドで発生する変更イベント
    if (event.target.files) {
      // ユーザーが選択したファイルのリストを含む FileList オブジェクト
      setFile(event.target.files[0]); // ユーザーが選択した最初のファイルをfileに設定
    }
  };

  const handleUpload = async () => {
    if (!file) {
      setMessage('ファイルが選択されていません。');
      return;
    }

    const formData = new FormData(); // キーバリューのペアを簡単に構築するためのオブジェクトで、特にファイルやデータをサーバーに送信する際に使用
    formData.append('file', file);
    const jwt = sessionStorage.getItem('token');

    try {
      const response = await axios.post(
        'http://localhost:8082/admins/questionsdata/import_csv',
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data',
            Authorization: `Bearer ${jwt}`,
          },
        }
      );
      if (response.status === 200) {
        setMessage('ファイルが正常にアップロードされました。');
      } else {
        setMessage('ファイルのアップロードに失敗しました。');
      }
      handleClear();
      // setFile(null);
      // if (fileInputRef.current) {
      //   fileInputRef.current.value = '';
      // }
    } catch (error) {
      console.error('Error uploading file:', error);
      setMessage('ファイルのアップロード中にエラーが発生しました。');
    }
  };

  const handleClear = () => {
    setFile(null);
    if (fileInputRef.current) {
      fileInputRef.current.value = '';
    }
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">CSVファイルのアップロード</h2>
      <div className="button-group">
        <input
          type="file"
          accept=".csv"
          onChange={handleFileChange}
          ref={fileInputRef}
          className="file-input"
          id="file-upload"
        />
        <label htmlFor="file-upload" className="upload-csv-button">
          ファイルを選択
        </label>
        <button onClick={handleClear} className="clear-button">
          選択を解除
        </button>
        <button onClick={handleUpload} className="upload-button">
          アップロード
        </button>
      </div>
      {file && <p className="selected-file">選択されたファイル: {file.name}</p>}
      {message && <p className="message">{message}</p>}
    </div>
  );
};

export default UploadCSV;
