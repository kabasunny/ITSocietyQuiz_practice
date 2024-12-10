import { useState } from 'react';
import axios from 'axios';
import { LoginForm } from '../types';

export const useLogin = (
  onLogin: (loginOK: boolean, isAdmin: boolean) => void
) => {
  const [loading, setLoading] = useState(false);
  const [errorMessage, setErrorMessage] = useState(''); // エラーメッセージの状態管理
  // console.log('API URL:', process.env.REACT_APP_API_URL);

  const onSubmit = async (loginForm: LoginForm) => {
    setLoading(true);
    try {
      const response = await axios.post(
        `${process.env.REACT_APP_API_URL}/login`, // 環境変数を使用してベースURLを設定
        loginForm,
        {
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );
      if (response.status >= 200 && response.status < 300) {
        const result = response.data;
        sessionStorage.setItem('token', result.token); // トークンを保存
        sessionStorage.setItem('admin', result.admin); // 管理者かどうか
        sessionStorage.setItem('todays_count', result.todays_count); // 本日の実施済みクイズ数を保存
        onLogin(true, result.admin); // adminフラグを渡す
      } else {
        setErrorMessage('社員IDまたはパスワードが間違っています'); // エラーメッセージを設定
        onLogin(false, false);
      }
    } catch (error) {
      console.error('Login error:', error); // エラーの詳細をログに記録
      setErrorMessage('ログインに失敗しました'); // エラーメッセージを設定
      onLogin(false, false);
    } finally {
      setLoading(false);
    }
  };

  return { onSubmit, loading, errorMessage }; // エラーメッセージを返す
};
