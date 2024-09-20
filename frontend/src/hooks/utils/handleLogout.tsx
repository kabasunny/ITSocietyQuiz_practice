import { NavigateFunction } from 'react-router-dom';

const handleLogout = (
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>,
  setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>,
  setIsSubmitAnsewr: React.Dispatch<React.SetStateAction<boolean>>,
  setShowScore: React.Dispatch<React.SetStateAction<boolean>>,
  navigate: NavigateFunction
) => {
  sessionStorage.removeItem('quizdata');
  sessionStorage.removeItem('todays_finish');
  sessionStorage.removeItem('currentQuestion');
  sessionStorage.removeItem('answers');

  // コンソールログで確認
  console.log('quizdata:', sessionStorage.getItem('quizdata'));
  console.log('todays_finish:', sessionStorage.getItem('todays_finish'));
  console.log('currentQuestion:', sessionStorage.getItem('currentQuestion'));
  console.log('answers:', sessionStorage.getItem('answers'));

  setIsLoggedIn(false);
  setIsAdmin(false);
  setIsSubmitAnsewr(false);
  setShowScore(false);
  navigate('/'); // ホームページに遷移
  // 他のログアウト処理をここに追加
};

export default handleLogout;
