import React from 'react';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { AdminsUser, UserFormProps } from '../../../types';
import { addUserSchema, editUserSchema } from '../../utils/ValidationSchema';

const UserForm: React.FC<UserFormProps> = ({ user, onSave, onCancel, isEditing }) => {
  // 編集モードか新規追加モードかに応じてスキーマを選択
  const schema = isEditing ? editUserSchema : addUserSchema;

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<AdminsUser>({
    resolver: zodResolver(schema),
    defaultValues: isEditing ? user : { ...user, roleId: 2 }, // 新規追加の場合はroleIdを2に設定
  });

  const onSubmit = (data: AdminsUser) => {
    console.log(data); // 送信データをコンソールに出力
    onSave(data);
  };


  return (
    <form className="user-form" onSubmit={handleSubmit(onSubmit)} autoComplete="off">{/* 自動補完をオフ */}
      <h3>ユーザー情報を{isEditing ? '編集は、変更箇所のみ編集ください。' : '追加は、全項目の入力が必須です。'}</h3>
      <div className="form-group">
        <label>社員ID : </label>
        <input type="text" {...register('empId')} autoComplete="off" />
        {errors.empId && <p className="error-message">{errors.empId?.message}</p>}
      </div>
      <div className="form-group">
        <label>社員氏名 : </label>
        <input type="text" {...register('name')} autoComplete="off" />
        {errors.name && <p className="error-message">{errors.name?.message}</p>}
      </div>
      <div className="form-group">
        <label>Eメールアドレス : </label>
        <input type="email" {...register('email')} autoComplete="off" />
        {errors.email && <p className="error-message">{errors.email?.message}</p>}
      </div>
      <div className="form-group">
        <label>{isEditing ? '現在のパスワード' : 'パスワード'} : </label>
        <input type="password" {...register('password_1')} autoComplete="new-password" />
        {errors.password_1 && <p className="error-message">{errors.password_1?.message}</p>}
      </div>
      <div className="form-group">
        <label>{isEditing ? '新規パスワード' : 'パスワードの確認'} : </label>
        <input type="password" {...register('password_2')} autoComplete="new-password" />
        {errors.password_2 && <p className="error-message">{errors.password_2?.message}</p>}
      </div>
      <div className="form-group">
        <label>権限 : </label>
        <select {...register('roleId', { valueAsNumber: true })} defaultValue={user.roleId || 2} onChange={(e) => console.log('Selected roleId:', e.target.value)}>
          <option value={2}>一般</option>
          <option value={1}>クイズ管理者</option>
        </select>

        {errors.roleId && <p className="error-message">{errors.roleId?.message}</p>}
      </div>
      <div className="button-container">
        <button type="submit" className="add-button">
          {isEditing ? '更新' : 'ユーザーの追加'}
        </button>
        <button className="cancel-button" type="button" onClick={onCancel}>
          キャンセル
        </button>
      </div>
    </form>
  );
};

export default UserForm;
