import React, { ChangeEvent } from 'react';
import { AdminsUser, UserFormProps } from '../../../types';

const UserForm: React.FC<UserFormProps> = ({ user, onChange, onSave, onCancel, isEditing }) => {
  const handleSelectChange = (e: ChangeEvent<HTMLSelectElement>) => {
    const value = parseInt(e.target.value, 10);
    onChange(e as unknown as ChangeEvent<HTMLInputElement>, 'roleId');
  };

  return (
    <div className="user-form">
      <h3>ユーザー情報を{isEditing ? '編集' : '追加'}</h3>
      <div className="form-group">
        <label>社員ID : </label>
        <input type="text" value={user.empId} onChange={(e) => onChange(e, 'empId')} />
      </div>
      <div className="form-group">
        <label>社員氏名 : </label>
        <input type="text" value={user.name} onChange={(e) => onChange(e, 'name')} />
      </div>
      <div className="form-group">
        <label>Eメールアドレス : </label>
        <input type="email" value={user.email} onChange={(e) => onChange(e, 'email')} />
      </div>
      <div className="form-group">
        <label>{isEditing ? '現在のパスワード' : 'パスワード'} : </label>
        <input type="password" value={user.password_1} onChange={(e) => onChange(e, 'password_1')} />
      </div>
      <div className="form-group">
        <label>{isEditing ? '新規パスワード' : 'パスワードの確認'} : </label>
        <input type="password" value={user.password_2} onChange={(e) => onChange(e, 'password_2')} />
      </div>
      <div className="form-group">
        <label>権限 : </label>
        <select value={user.roleId} onChange={handleSelectChange}>
          <option value={2}>一般</option>
          <option value={1}>クイズ管理者</option>
        </select>
      </div>
      <button className="add-button" onClick={onSave}>{isEditing ? '更新' : 'ユーザーの追加'}</button>
      <button onClick={onCancel}>キャンセル</button>
    </div>
  );
};

export default UserForm;
