import { NavigateFunction } from 'react-router-dom';

const handleLogout = (
  setIsLoggedIn: React.Dispatch<React.SetStateAction<boolean>>,
  setIsAdmin: React.Dispatch<React.SetStateAction<boolean>>,
  setIsSubmitAnsewr: React.Dispatch<React.SetStateAction<boolean>>,
  setShowScore: React.Dispatch<React.SetStateAction<boolean>>,
  navigate: NavigateFunction
) => {
  // sessionStorage.removeItem('quizdata');
  // sessionStorage.removeItem('todays_finish');
  // sessionStorage.removeItem('currentQuestion');
  // sessionStorage.removeItem('answers');

  sessionStorage.clear(); //scoreがセッションに残っていて、他のユーザーに加算されていた

  setIsLoggedIn(false);
  setIsAdmin(false);
  setIsSubmitAnsewr(false);
  setShowScore(false);
  navigate('/'); // ホームページに遷移
};

export default handleLogout;
