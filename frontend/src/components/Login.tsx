import React from 'react';
import { useForm } from 'react-hook-form'; //カスタムフック
import { zodResolver } from '@hookform/resolvers/zod';
import { validationSchema } from './utils/ValidationSchema';
import './Login.css';

interface LoginProps {
  onLogin: (name: string) => void;
}

interface LoginForm {
    empid: string;
    name: string;
    email: string;
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
    onLogin(data.name);
  };

  return (
    <div className="form-container">
      <h1>ITSocietyQuiz</h1>
      <form onSubmit={handleSubmit(onSubmit)}>
      <label htmlFor="empid">社員IDを入力してね \( ᐕ )/</label>
        <input
          type="text"
          id="empid"
          // validationSchemaを使用しない場合
          // {...register("empid", {
          //   required: "社員IDは必須です",
          //   minLength: { value: 6, message: "6文字以上で入力してください" },
          // })}
          {...register("empid")}
        />
        {errors.empid && <p>{errors.empid.message as React.ReactNode}</p>}
        {/* <p>{errors.empid?.message as React.ReactNode}</p> */}
        {/* as React.ReactNode は型アサーション */}

        {/* <label htmlFor="name">社員IDを入力してね </label>
        <input
          type="text"
          id="name"
          // validationSchemaを使用しない場合
          // {...register("name", {
          //   required: "名前は必須です",
          //   minLength: { value: 4, message: "4文字以上で入力してください" },
          // })}
          {...register("name")}
        />
        {errors.name && <p>{errors.name.message as React.ReactNode}</p>} */}

        {/* <label htmlFor="email">メールアドレス</label>
        <input
          type="email"
          id="email"
          // validationSchemaを使用しない場合
          // {...register("email", {
          //   required: "メールアドレスは必須です",
          //   pattern: {
          //     value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
          //     message: "正しいメールアドレスを入力してください",
          //   },
          // })}
          {...register("email")}
        />
        {errors.email && <p>{errors.email.message as React.ReactNode}</p>} */}

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
