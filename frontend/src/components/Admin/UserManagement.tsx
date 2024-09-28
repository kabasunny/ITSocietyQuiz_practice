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
  const [newUser, setNewUser] = useState<AdminsUser>({
    dbId: 0,
    empId: '',
    name: '',
    email: '',
    password_1: '',
    password_2: '',
    roleId: 2, // 一般ユーザーのroleIdをデフォルトに設定
    roleName: '',
    updatedAt: '',
    createdAt: ''
  });
  const [showAddForm, setShowAddForm] = useState(false); // 新しいユーザー追加フォームの表示状態

  const jwt = sessionStorage.getItem('token');

  useEffect(() => {
    // ユーザー情報を取得するAPI呼び出し
    axios.get('http://localhost:8082/admins/userslist', {
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    })
      .then(response => {
        setUsers(response.data.userlist);
        console.log(users);
      })
      .catch(error => {
        console.error('Error fetching users:', error);
      });
  }, [jwt, editingUser]);

  // const handleInputChange = (e: ChangeEvent<HTMLInputElement>, field: keyof AdminsUser) => {
  //   if (editingUser) {
  //     setEditingUser({ ...editingUser, [field]: e.target.value });
  //   } else {
  //     setNewUser({ ...newUser, [field]: e.target.value });
  //   }
  // };

  const handleAddUser = (data: AdminsUser) => {
    // ユーザー情報を新規追加するAPI呼び出し
    axios.post('http://localhost:8082/admins/addusers', data, {
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    })
      .then(response => {
        setUsers([...users, response.data.userlist]);
        setShowAddForm(false); // フォームを閉じる
      })
      .catch(error => {
        console.error('Error adding user:', error);
      });
  };

  const handleEditUser = (user: AdminsUser) => {
    setEditingUser(user);
  };

  const handleUpdateUser = (data: AdminsUser) => {
    if (editingUser) {
      data.dbId = editingUser.dbId; // 初期に受け取った dbId を設定
    }
    axios.put(`http://localhost:8082/admins/updateusers/${data.dbId}`, data, {
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    })
      .then(response => {
        if (response.status === 200) {
          setUsers(users.map(user => user.dbId === data.dbId ? response.data : user));
          alert('ユーザー情報が正常に更新されました。');
        } else {
          console.error('Unexpected response status:', response.status);
        }
      })
      .catch(error => {
        console.error('Error updating user:', error);
      });
    setEditingUser(null);
  };
  

  const handleDeleteUser = (empId: string) => {
    const confirmDelete = window.confirm('本当にこのユーザーを削除しますか？'); // 確認ダイアログを表示
    if (confirmDelete) {
      // ユーザー情報を削除するAPI呼び出し
      axios.delete(`http://localhost:8082/users/${empId}`, {
        headers: {
          Authorization: `Bearer ${jwt}`,
        },
      })
        .then(() => {
          setUsers(users.filter(user => user.empId !== empId));
        })
        .catch(error => {
          console.error('Error deleting user:', error);
        });
    }
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">ユーザー管理（社員情報を更新、削除、追加）</h2>
      {!showAddForm && !editingUser && (
        <UserTable users={users} onEditUser={handleEditUser} onDeleteUser={handleDeleteUser} />
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
        <button className="add-button" onClick={() => setShowAddForm(true)}>新規ユーザーの登録</button>
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
