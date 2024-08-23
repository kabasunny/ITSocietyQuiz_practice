import React from 'react';
import { useForm } from 'react-hook-form'; //カスタムフック
import { zodResolver } from '@hookform/resolvers/zod';
import { validationSchema } from './utils/ValidationSchema';
import './Login.css';
import { users } from '../data/Users';

interface LoginProps {
  onLogin: (data: boolean) => void;
}

export interface LoginForm {
  empid: string;
  password: string;
}

const Login: React.FC<LoginProps> = ({ onLogin }) => {
  const {
    register, //inputタグに含める
    handleSubmit,
    formState: { errors },
  } = useForm<LoginForm>({
    mode: "onChange", //バリデーション発火のタイミング
    resolver: zodResolver(validationSchema),
  });
  
  const onSubmit = (data: LoginForm) => {
    const loginOK = users.find(u => u.empid === data.empid && u.password === data.password);
    if (loginOK) {
      onLogin(true);
    } else {
      console.error('社員IDまたはパスワードが間違っています');
    }
  };

  return (
    <div className="form-container">
      <h1>ITSocietyQuiz</h1>
      <form onSubmit={handleSubmit(onSubmit)}>
      <label htmlFor="empid">社員IDを入力してね \( ᐕ )/</label>
        <input
          type="text"
          id="empid"
          {...register("empid")}
        />
        {errors.empid && <p>{errors.empid.message as React.ReactNode}</p>}
        <label htmlFor="password">パスワードを入力してね \( ᐛ )/</label>
        <input
          id="password"
          type="password"
          {...register("password")}
        />
        {errors.password && <p>{errors.password.message as React.ReactNode}</p>}

        <button type="submit">ログイン</button>
      </form>
    </div>
  );
}

export default Login;
