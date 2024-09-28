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
    onSave(data);
  };

  return (
    <form className="user-form" onSubmit={handleSubmit(onSubmit)}>
      <h3>ユーザー情報を{isEditing ? '編集' : '追加'}</h3>
      <div className="form-group">
        <label>社員ID : </label>
        <input type="text" {...register('empId')} />
        {errors.empId && <p className="error-message">{errors.empId?.message}</p>}
      </div>
      <div className="form-group">
        <label>社員氏名 : </label>
        <input type="text" {...register('name')} />
        {errors.name && <p className="error-message">{errors.name?.message}</p>}
      </div>
      <div className="form-group">
        <label>Eメールアドレス : </label>
        <input type="email" {...register('email')} />
        {errors.email && <p className="error-message">{errors.email?.message}</p>}
      </div>
      <div className="form-group">
        <label>{isEditing ? '現在のパスワード' : 'パスワード'} : </label>
        <input type="password" {...register('password_1')} />
        {errors.password_1 && <p className="error-message">{errors.password_1?.message}</p>}
      </div>
      <div className="form-group">
        <label>{isEditing ? '新規パスワード' : 'パスワードの確認'} : </label>
        <input type="password" {...register('password_2')} />
        {errors.password_2 && <p className="error-message">{errors.password_2?.message}</p>}
      </div>
      <div className="form-group">
        <label>権限 : </label>
        <select {...register('roleId', { valueAsNumber: true })} defaultValue={user.roleId || 2}>
          <option value={2}>一般</option>
          <option value={1}>クイズ管理者</option>
        </select>
        {errors.roleId && <p className="error-message">{errors.roleId?.message}</p>}
      </div>
      <button type="submit" className="add-button">
        {isEditing ? '更新' : 'ユーザーの追加'}
      </button>
      <button type="button" onClick={onCancel}>
        キャンセル
      </button>
    </form>
  );
};

export default UserForm;
