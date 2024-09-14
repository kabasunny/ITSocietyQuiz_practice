import { useState } from 'react';
import axios from 'axios';
import { LoginForm } from '../types';

export const useLogin = (onLogin: (loginOK: boolean, isAdmin: boolean) => void) => {
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState(''); // エラーメッセージの状態管理

  const onSubmit = async (loginForm: LoginForm) => {
    setLoading(true);
    try {
      if (process.env.REACT_APP_USE_API === 'true') {
        const response = await axios.post('http://localhost:8082/login', loginForm, {
          headers: {
            'Content-Type': 'application/json',
          },
        });
        if (response.status >= 200 && response.status < 300) {
          const result = response.data;
          localStorage.setItem('token', result.token); // トークンを保存
          localStorage.setItem('todays_count', result.todays_count); // 本日の実施済みクイズ数を保存
          onLogin(true, result.admin); // adminフラグを渡す
        } else {
          setErrorMessage('社員IDまたはパスワードが間違っています'); // エラーメッセージを設定
          onLogin(false, false);
        }
      } else { // ローカルデータでの検証はとりあえず消しますか
        
      }
    } catch (error) {
      setErrorMessage('ログインに失敗しました'); // エラーメッセージを設定
      onLogin(false, false);
    } finally {
      setLoading(false);
    }
  };

  return { onSubmit, loading, errorMessage }; // エラーメッセージを返す
};
