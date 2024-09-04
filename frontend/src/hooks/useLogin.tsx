import { useState } from 'react';
import axios from 'axios';
import { LoginForm } from '../types';
import { users } from '../data/Users'; // ローカルデータをインポート

export const useLogin = (onLogin: (login: boolean) => void) => {
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
        // console.log('API Response:', result); // レスポンスデータをログに出力
        if (response.status >= 200 && response.status < 300) {
          const result = response.data;
          localStorage.setItem('token', result.token); // トークンを保存
          onLogin(true);
        } else {
          setErrorMessage('社員IDまたはパスワードが間違っています'); // エラーメッセージを設定
        }
      } else { // process.env.REACT_APP_USE_API === 'false' ローカルデータでの検証
        const loginOK = users.find(user => user.empid === loginForm.empid && user.password === loginForm.password);
        if (loginOK) {
          onLogin(true);
        } else {
          setErrorMessage('社員IDまたはパスワードが間違っています'); // エラーメッセージを設定
        }
      }
    } catch (error) {
      setErrorMessage('ログインに失敗しました'); // エラーメッセージを設定
    } finally {
      setLoading(false);
    }
  };

  return { onSubmit, loading, errorMessage }; // エラーメッセージを返す
};
