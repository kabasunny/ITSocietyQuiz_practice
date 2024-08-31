import { Dispatch, SetStateAction } from 'react';

const handleLogin = (loginOK: boolean, setIsLoggedIn: Dispatch<SetStateAction<boolean>>) => {
  setIsLoggedIn(loginOK);
};

export default handleLogin;
