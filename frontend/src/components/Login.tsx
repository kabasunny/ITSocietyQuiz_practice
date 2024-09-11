import React, { useState } from 'react';
import { useForm } from 'react-hook-form'; //カスタムフックでフォームのバリデーションと送信を管理
import { zodResolver } from '@hookform/resolvers/zod';
import { validationSchema } from './utils/ValidationSchema';
import './Login.css';
import { useLogin } from '../hooks/useLogin';
import { LoginForm, LoginProps } from '../types';

const Login: React.FC<LoginProps> = ({ onLogin }) => {
  const { onSubmit, loading, errorMessage } = useLogin(onLogin);
  const {
    register, //React Hook Formに登録しフィールドとして管理するための関数、inputタグに含める
    handleSubmit, // フォームが送信されたときに呼び出される関数
    formState: { errors },
  } = useForm<LoginForm>({
    mode: "onChange", //バリデーション発火のタイミング
    resolver: zodResolver(validationSchema),
  });

  return (
    <div className="form-container">
      {/* 環境変数からタイトルを読み込む */}
      <h1>{process.env.REACT_APP_TITLE}</h1>
      <form onSubmit={handleSubmit(onSubmit)} autoComplete="off">
        {/* 社員IDの入力フィールド */}
        <label htmlFor="empid">社員IDを入力してね \( ᐕ )/</label>
        <input
          type="text"
          id="empid"
          {...register("empid")}
          autoComplete="off"
        />
        {errors.empid && <p>{errors.empid.message as React.ReactNode}</p>}
        
        {/* パスワードの入力フィールド */}
        <label htmlFor="password">パスワードを入力してね \( ᐛ )/</label>
        <input
          id="password"
          type="password"
          {...register("password")}
          autoComplete="off"
        />
        {errors.password && <p>{errors.password.message as React.ReactNode}</p>}

         {/* ログインボタン、onSubmit取得後にloadingはtrue*/}
         <button type="submit" disabled={loading}>
          {loading ? '送信中...' : (
            <>
              ログイン<br />& スタート！
            </>
          )}
        </button>
      </form>

      {/* エラーメッセージの表示 */}
      {errorMessage && <p style={{ color: 'red' }}>{errorMessage}</p>}

      <p><br />60分以内に全問解答できない場合は、<br/>全問不正解になります。</p>
    </div>
  );
}

export default Login;
