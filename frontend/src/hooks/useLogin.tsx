import { useState } from 'react';
import axios from 'axios';
import { LoginForm } from '../components/Login';
import { users } from '../data/Users'; // ローカルデータをインポート

export const useLogin = (onLogin: (login: boolean) => void) => {
  const [loading, setLoading] = useState(false);

  const onSubmit = async (data: LoginForm) => {
    setLoading(true);
    try {
      if (process.env.REACT_APP_USE_API === 'true') {
        const response = await axios.post('http://localhost:8082/questions/login', data, {
          headers: {
            'Content-Type': 'application/json',
          },
        });
        const result = response.data;
        if (result.success) {
          onLogin(true);
        } else {
          console.error(result.message);
        }
      } else { // process.env.REACT_APP_USE_API === 'false' ローカルデータでの検証
        const loginOK = users.find(u => u.empid === data.empid && u.password === data.password);
        if (loginOK) {
          onLogin(true);
        } else {
          console.error('社員IDまたはパスワードが間違っています');
        }
      }
    } catch (error) {
      console.error('サーバーとの通信に失敗しました', error);
    } finally {
      setLoading(false);
    }
  };

  return { onSubmit, loading };
};
