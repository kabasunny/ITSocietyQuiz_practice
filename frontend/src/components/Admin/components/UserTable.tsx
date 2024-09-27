import React from 'react';
import { formatDate } from '../utils/formatDate';
import { AdminsUser, UserTableProps } from '../../../types' 



const UserTable: React.FC<UserTableProps> = ({ users, onEditUser, onDeleteUser }) => {
  return (
    <div className="admin-table-container">
      <table className="admin-user-table">
        <thead>
          <tr>
            <th>社員ID</th>
            <th>社員氏名</th>
            <th>Eメールアドレス</th>
            <th>権限</th>
            <th>最終更新日</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          {users.map(user => (
            <tr key={user.empId}>
              <td>{user.empId}</td>
              <td>{user.name}</td>
              <td>{user.email}</td>
              <td>{user.roleName}</td>
              <td dangerouslySetInnerHTML={{ __html: formatDate(user.updatedAt > user.createdAt ? user.updatedAt : user.createdAt) }}></td>
              <td>
                <button className="edit-button" onClick={() => onEditUser(user)}>編集</button>
                <br/>
                <button className="delete-button" onClick={() => onDeleteUser(user.empId)}>削除</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default UserTable;