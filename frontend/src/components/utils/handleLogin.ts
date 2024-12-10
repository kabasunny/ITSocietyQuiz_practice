import { Dispatch, SetStateAction } from 'react';

const handleLogin = (
  loginOK: boolean,
  setIsLoggedIn: Dispatch<SetStateAction<boolean>>,
  setIsAdmin: Dispatch<SetStateAction<boolean>>,
  isAdmin: boolean
) => {
  if (loginOK) {
    setIsLoggedIn(true);
    setIsAdmin(isAdmin);
    window.location.reload(); // ブラウザをリロード
  } else {
    setIsLoggedIn(false);
  }
};

export default handleLogin;
