import React, { useState, useEffect, ChangeEvent } from 'react';
import axios from 'axios';
import './css/AdminScreen.css'; // CSSファイルをインポート
import './css/UserManagement.css'; // 新しいCSSファイルをインポート
import UserTable from './components/UserTable';
import UserForm from './components/UserForm';
import { AdminsUser } from '../../types';

const UserManagement: React.FC = () => {
  const [users, setUsers] = useState<AdminsUser[]>([]);
  const [editingUser, setEditingUser] = useState<AdminsUser | null>(null);
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
  const [showAddForm, setShowAddForm] = useState(false); // 新しいユーザー追加フォームの表示状態

  const jwt = sessionStorage.getItem('token');

  useEffect(() => {
    // ユーザー情報を取得するAPI呼び出し
    axios
      .get('http://localhost:8082/admins/userdata/userslist', {
        headers: {
          Authorization: `Bearer ${jwt}`,
        },
      })
      .then((response) => {
        setUsers(response.data.userlist);
        console.log(users);
      })
      .catch((error) => {
        console.error('Error fetching users:', error);
      });
  }, [jwt, editingUser]);

  const handleAddUser = (data: AdminsUser) => {
    console.log(data);
    // ユーザー情報を新規追加するAPI呼び出し
    axios
      .post('http://localhost:8082/admins/userdata/addusers', data, {
        headers: {
          Authorization: `Bearer ${jwt}`,
          'Content-Type': 'application/json',
        },
      })
      .then((response) => {
        setUsers([...users, response.data]);
        alert('ユーザー情報が正常に更新されました。');
        setShowAddForm(false); // フォームを閉じる
      })
      .catch((error) => {
        console.error('Error adding user:', error);
        console.error(
          'Error details:',
          error.response ? error.response.data : error.message
        );
      });
  };

  const handleEditUser = (user: AdminsUser) => {
    setEditingUser(user);
  };

  const handleUpdateUser = (data: AdminsUser) => {
    if (editingUser) {
      data.dbId = editingUser.dbId; // 初期に受け取った dbId を設定
    }
    axios
      .put(
        `http://localhost:8082/admins/userdata/updateusers/${data.dbId}`,
        data,
        {
          headers: {
            Authorization: `Bearer ${jwt}`,
          },
        }
      )
      .then((response) => {
        if (response.status === 200) {
          const updatedUser = response.data;
          setUsers(
            users.map((user) =>
              user.dbId === data.dbId ? { ...user, ...updatedUser } : user
            )
          );
          console.log(updatedUser);
          alert('ユーザー情報が正常に更新されました。');
        } else {
          console.error('Unexpected response status:', response.status);
        }
      })

      .catch((error) => {
        console.error('Error updating user:', error);
      });
    setEditingUser(null);
  };

  useEffect(() => {
    console.log('Users updated:', users);
  }, [users]);

  const handleDeleteUser = (dbId: number) => {
    const confirmDelete = window.confirm('本当にこのユーザーを削除しますか？'); // 確認ダイアログを表示
    if (confirmDelete) {
      // ユーザー情報を削除するAPI呼び出し
      axios
        .delete(`http://localhost:8082/admins/userdata/deleteuser/${dbId}`, {
          headers: {
            Authorization: `Bearer ${jwt}`,
          },
        })
        .then(() => {
          setUsers(users.filter((user) => user.dbId !== dbId));
        })
        .catch((error) => {
          console.error('Error deleting user:', error);
        });
    }
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">ユーザー管理（社員情報を更新、削除、追加）</h2>
      {!showAddForm && !editingUser && (
        <UserTable
          users={users}
          onEditUser={handleEditUser}
          onDeleteUser={handleDeleteUser}
        />
      )}
      {editingUser && (
        <UserForm
          user={editingUser}
          onSave={handleUpdateUser}
          onCancel={() => setEditingUser(null)}
          isEditing={true}
        />
      )}
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
