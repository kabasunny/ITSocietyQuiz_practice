import React, { useState, useEffect, ChangeEvent } from 'react';
import axios from 'axios';
import './css/AdminScreen.css'; // CSSファイルをインポート
import './css/UserManagement.css'; // 新しいCSSファイルをインポート

interface User {
  empid: string;
  name: string;
  email: string;
  password: string;
  role: string;
  updatedAt: string;
}

const UserManagement: React.FC = () => {
  const [users, setUsers] = useState<User[]>([]);
  const [editingUser, setEditingUser] = useState<User | null>(null);
  const [newUser, setNewUser] = useState<User>({
    empid: '',
    name: '',
    email: '',
    password: '',
    role: '',
    updatedAt: ''
  });

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
      })
      .catch(error => {
        console.error('Error fetching users:', error);
      });
  }, []);

  const handleInputChange = (e: ChangeEvent<HTMLInputElement>, field: keyof User) => {
    setNewUser({ ...newUser, [field]: e.target.value });
  };

  const handleAddUser = () => {
    axios.post('http://localhost:8082/users', newUser)
      .then(response => {
        setUsers([...users, response.data.userlist]);
        setNewUser({
          empid: '',
          name: '',
          email: '',
          password: '',
          role: '',
          updatedAt: ''
        });
      })
      .catch(error => {
        console.error('Error adding user:', error);
      });
  };

  const handleEditUser = (user: User) => {
    setEditingUser(user);
  };

  const handleUpdateUser = () => {
    if (editingUser) {
      axios.put(`http://localhost:8082/users/${editingUser.empid}`, editingUser)
        .then(response => {
          setUsers(users.map(user => user.empid === editingUser.empid ? response.data : user));
          setEditingUser(null);
        })
        .catch(error => {
          console.error('Error updating user:', error);
        });
    }
  };

  const handleDeleteUser = (empid: string) => {
    axios.delete(`http://localhost:8082/users/${empid}`)
      .then(() => {
        setUsers(users.filter(user => user.empid !== empid));
      })
      .catch(error => {
        console.error('Error deleting user:', error);
      });
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">ユーザー管理（社員情報を更新、削除、追加）</h2>
      <div className="admin-table-container">
      <div className="table-container">
        <table className="admin-table">
          <thead>
            <tr>
              <th>社員ID</th>
              <th>社員氏名</th>
              <th>Eメールアドレス</th>
              <th>パスワード</th>
              <th>権限</th>
              <th>最終更新日</th>
              <th>操作</th>
            </tr>
          </thead>
          <tbody>
            {users.map(user => (
              <tr key={user.empid}>
                <td>{user.empid}</td>
                <td>{user.name}</td>
                <td>{user.email}</td>
                <td>{user.password}</td>
                <td>{user.role}</td>
                <td>{user.updatedAt}</td>
                <td>
                  <button className="edit-button" onClick={() => handleEditUser(user)}>編集</button>
                  <button className="delete-button" onClick={() => handleDeleteUser(user.empid)}>削除</button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div></div>
      {editingUser && (
        <div className="edit-form">
          <h3>ユーザー情報を編集</h3>
          <label>
            社員ID:
            <input type="text" value={editingUser.empid} onChange={(e) => setEditingUser({ ...editingUser, empid: e.target.value })} />
          </label>
          <label>
            社員氏名:
            <input type="text" value={editingUser.name} onChange={(e) => setEditingUser({ ...editingUser, name: e.target.value })} />
          </label>
          <label>
            Eメールアドレス:
            <input type="email" value={editingUser.email} onChange={(e) => setEditingUser({ ...editingUser, email: e.target.value })} />
          </label>
          <label>
            パスワード:
            <input type="password" value={editingUser.password} onChange={(e) => setEditingUser({ ...editingUser, password: e.target.value })} />
          </label>
          <label>
            権限:
            <input type="text" value={editingUser.role} onChange={(e) => setEditingUser({ ...editingUser, role: e.target.value })} />
          </label>
          <button onClick={handleUpdateUser}>更新</button>
          <button onClick={() => setEditingUser(null)}>キャンセル</button>
        </div>
      )}
      <div className="user-form">
        <h3>新しいユーザーを追加</h3>
        <label>
          社員ID:
          <input type="text" value={newUser.empid} onChange={(e) => handleInputChange(e, 'empid')} />
        </label>
        <label>
          社員氏名:
          <input type="text" value={newUser.name} onChange={(e) => handleInputChange(e, 'name')} />
        </label>
        <label>
          Eメールアドレス:
          <input type="email" value={newUser.email} onChange={(e) => handleInputChange(e, 'email')} />
        </label>
        <label>
          パスワード:
          <input type="password" value={newUser.password} onChange={(e) => handleInputChange(e, 'password')} />
        </label>
        <label>
          権限:
          <input type="text" value={newUser.role} onChange={(e) => handleInputChange(e, 'role')} />
        </label>
        <button onClick={handleAddUser}>追加</button>
      </div>
    </div>
  );
  
};

export default UserManagement;
