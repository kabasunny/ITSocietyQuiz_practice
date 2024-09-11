import { Dispatch, SetStateAction } from 'react';

const handleLogin = (loginOK: boolean, setIsLoggedIn: Dispatch<SetStateAction<boolean>>, setIsAdmin: Dispatch<SetStateAction<boolean>>, isAdmin: boolean) => {
  setIsLoggedIn(loginOK);
  setIsAdmin(isAdmin);
};

export default handleLogin;
