import React from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { z } from 'zod';
import { validationSchema } from './utils/ValidationSchema';
import './Login.css';

interface LoginProps {
  onLogin: (name: string) => void;
}

interface LoginForm {
    name: string;
    email: string;
    password: string;
  }

const Login: React.FC<LoginProps> = ({ onLogin }) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginForm>({
    mode: "onChange",
    resolver: zodResolver(validationSchema),
  });

  const onSubmit = (data: LoginForm) => {
    onLogin(data.name);
  };

  return (
    <div className="form-container">
      <h1>ITSocietyQuiz</h1>
      <form onSubmit={handleSubmit(onSubmit)}>
        <label htmlFor="">名前</label>
        <input
          type="text"
          id="name"
          // {...register("name", {
          //   required: "名前は必須です",
          //   minLength: { value: 4, message: "4文字以上で入力してください" },
          // })}
          {...register("name")}
        />
        {errors.name && <p>{errors.name.message as React.ReactNode}</p>}

        <label htmlFor="email">メールアドレス</label>
        <input
          type="email"
          id="email"
          // {...register("email", {
          //   required: "メールアドレスは必須です",
          //   pattern: {
          //     value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
          //     message: "正しいメールアドレスを入力してください",
          //   },
          // })}
          {...register("email")}
        />
        {errors.email && <p>{errors.email.message as React.ReactNode}</p>}

        <label htmlFor="password">パスワード</label>
        <input
          id="password"
          type="password"
          // {...register("password", {
          //   required: "パスワードは必須です",
          //   minLength: { value: 8, message: "8文字以上で入力してください" },
          // })}
          {...register("password")}
        />
        {errors.password && <p>{errors.password.message as React.ReactNode}</p>}

        <button type="submit">送信</button>
      </form>
    </div>
  );
}

export default Login;
