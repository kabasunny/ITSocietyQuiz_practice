import React, { useState, useEffect } from 'react';
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
  const [currentPage, setCurrentPage] = useState(1); // 現在のページ番号の状態管理
  const itemsPerPage = 20; // 1ページあたりのアイテム数

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

  // ページを変更する関数
  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber); // 現在のページ番号を設定
  };

  const indexOfLastItem = currentPage * itemsPerPage; // 現在のページの最後のアイテムのインデックスを計算
  const indexOfFirstItem = indexOfLastItem - itemsPerPage; // 現在のページの最初のアイテムのインデックスを計算
  const currentItems = users.slice(indexOfFirstItem, indexOfLastItem); // 現在のページに表示するアイテムを取得

  return (
    <div className="admin-container">
      <h2 className="admin-h2">ユーザー管理（社員情報を更新、削除、追加）</h2>
      {/* 新規ユーザー追加フォームと編集フォームが表示されていない場合にユーザーテーブルを表示 */}
      {!showAddForm && !editingUser && (
        <UserTable
          users={currentItems}
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
      {/* ページネーションボタンを表示 */}
      {!editingUser && (
        <div className="pagination">
          {Array.from(
            { length: Math.ceil(users.length / itemsPerPage) },
            (_, index) => (
              <button
                key={index + 1}
                onClick={() => handlePageChange(index + 1)}
                className={currentPage === index + 1 ? 'active' : ''}
              >
                {index + 1}
              </button>
            )
          )}
        </div>
      )}
    </div>
  );
};

export default UserManagement;
