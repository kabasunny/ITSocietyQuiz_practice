import { Answer } from '../../types';

const handleResultsSubmit = (answers: Answer[]) => {
  const jwt = localStorage.getItem('token'); // ログイン時にAPIから取得したトークン
  if (jwt) {
    const result = {
      answers: answers,
      jwt: jwt
    };
    alert(`結果が送信されました。\nトークン: ${jwt}\n結果: ${JSON.stringify(answers, null, 2)}`);
  } else {
    alert('トークンが見つかりません');
  }
};

export default handleResultsSubmit;
