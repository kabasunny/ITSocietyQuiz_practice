import { NavigateFunction } from 'react-router-dom';

const handleAdminLogout = (
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>,
  setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>,
  navigate: NavigateFunction
) => {
  setIsLoggedIn(false);
  setIsAdmin(false);
  navigate('/'); // ホームページに遷移
  // 他のログアウト処理をここに追加
};

export default handleAdminLogout;
