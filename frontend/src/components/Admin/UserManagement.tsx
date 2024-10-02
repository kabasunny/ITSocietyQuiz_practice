import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './css/AdminScreen.css';
import './css/UserManagement.css';
import UserTable from './components/UserTable';
import UserForm from './components/UserForm';
import { AdminsUser } from '../../types';
import { fetchUsers, addUser, updateUser, deleteUser } from './utils/userApi';

const UserManagement: React.FC = () => {
  const [users, setUsers] = useState<AdminsUser[]>([]); // ユーザーリストの状態管理
  const [editingUser, setEditingUser] = useState<AdminsUser | null>(null); // 編集中のユーザーの状態管理
  const [showAddForm, setShowAddForm] = useState(false); // 新規ユーザー追加フォームの表示状態管理

  const jwt = sessionStorage.getItem('token'); // セッションストレージからJWTトークンを取得

  // 新規ユーザーのデフォルト値
  const newUser: AdminsUser = {
    dbId: 0,
    empId: '',
    name: '',
    email: '',
    password_1: '',
    password_2: '',
    roleId: 2, // 一般ユーザーのroleIdをデフォルトに設定
    roleName: '',
    updatedAt: '',
    createdAt: '',
  };

  // コンポーネントのマウント時および編集ユーザーの変更時にユーザーリストを取得
  useEffect(() => {
    fetchUsers(jwt, setUsers);
  }, [jwt, editingUser]);

  // 新規ユーザーを追加する関数
  const handleAddUser = (data: AdminsUser) => {
    addUser(data, jwt, users, setUsers, setShowAddForm);
  };

  // ユーザー編集を開始する関数
  const handleEditUser = (user: AdminsUser) => {
    setEditingUser(user);
  };

  // ユーザー情報を更新する関数
  const handleUpdateUser = (data: AdminsUser) => {
    updateUser(data, editingUser, jwt, users, setUsers, setEditingUser);
  };

  // ユーザーを削除する関数
  const handleDeleteUser = (dbId: number) => {
    deleteUser(dbId, jwt, users, setUsers);
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">ユーザー管理（社員情報を更新、削除、追加）</h2>
      {/* 新規ユーザー追加フォームと編集フォームが表示されていない場合にユーザーテーブルを表示 */}
      {!showAddForm && !editingUser && (
        <UserTable
          users={users}
          onEditUser={handleEditUser}
          onDeleteUser={handleDeleteUser}
        />
      )}
      {/* 編集フォームを表示 */}
      {editingUser && (
        <UserForm
          user={editingUser}
          onSave={handleUpdateUser}
          onCancel={() => setEditingUser(null)}
          isEditing={true}
        />
      )}
      {/* 新規ユーザー追加ボタンを表示 */}
      {!showAddForm && !editingUser && (
        <div className="add-user-button-container">
          <button
            className="add-user-button"
            onClick={() => setShowAddForm(true)}
          >
            新規ユーザーの登録
          </button>
        </div>
      )}
      {/* 新規ユーザー追加フォームを表示 */}
      {showAddForm && (
        <UserForm
          user={newUser}
          onSave={handleAddUser}
          onCancel={() => setShowAddForm(false)}
          isEditing={false}
        />
      )}
    </div>
  );
};

export default UserManagement;
