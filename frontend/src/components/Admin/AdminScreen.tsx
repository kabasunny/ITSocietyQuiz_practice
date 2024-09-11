import React from 'react';
import { Link } from 'react-router-dom';

interface AdminScreenProps {
  isAdmin: boolean;
}

const AdminScreen: React.FC<AdminScreenProps> = ({ isAdmin }) => {
  return (
    <div className="admin-container">
      <h1>管理者ダッシュボード</h1>
      <nav>
        <ul>
          {isAdmin && <li><Link to="/add-question">問題の追加</Link></li>}
          {isAdmin && <li><Link to="/update-delete-question">問題の更新、削除</Link></li>}
          {isAdmin && <li><Link to="/user-management">ユーザー管理</Link></li>}
          {isAdmin && <li><Link to="/statistics">統計と分析</Link></li>}
        </ul>
      </nav>
    </div>
  );
};

export default AdminScreen;
