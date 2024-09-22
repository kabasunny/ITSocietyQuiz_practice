import React from 'react';
import { useNavigate } from 'react-router-dom';
import './AdminScreen.css'; // インポート文の追加

interface AdminScreenProps {
  isAdmin: boolean;
  onLogout: () => void;
}

const AdminScreen: React.FC<AdminScreenProps> = ({ onLogout }) => {
  const navigate = useNavigate();

  const handleNavigation = (path: string) => {
    navigate(path, { state: { fromLink: true } });
  };

  const handleLogout = () => {
    const confirmed = window.confirm('本当にログアウトしますか？');
    if (confirmed) {
      onLogout();
    }
  };

  return (
    <div className="admin-container">
      <h1>管理者ダッシュボード(´◉◞౪◟◉)</h1>
      <div className="button-group">
        <button onClick={() => handleNavigation('/add-question')}>問題の追加</button>
        <button onClick={() => handleNavigation('/edit-question')}>問題の編集</button>
        <button onClick={() => handleNavigation('/user-management')}>ユーザー管理</button>
        <button onClick={() => handleNavigation('/statistics')}>統計と分析</button>
        <button className="logout-button" onClick={handleLogout}>ログアウト</button>
      </div>
    </div>
  );
};

export default AdminScreen;
