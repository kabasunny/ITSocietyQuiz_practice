import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom'; // 追加
import './index.css';
import App from './App';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <BrowserRouter> {/* 追加 */}
      <App />
    </BrowserRouter> {/* 追加 */}
  </React.StrictMode>
);


// 変更の経緯
// 管理者ダッシュボード内で画面遷移に Link コンポーネントを使用したところ、エラーが発生。
// これは、Link コンポーネントが Router コンテキスト内でのみ動作するように設計されているため。
// Router コンテキストがないと、Link コンポーネントは正しく動作しない。
// そこで、App コンポーネント全体を BrowserRouter でラップすることで、Router コンテキストが提供されるように。
// これにより、Link コンポーネントが正しく動作し、エラーが解消された。


// import React from 'react';
// import ReactDOM from 'react-dom/client';
// import './index.css';
// import App from './App';

// const root = ReactDOM.createRoot(
//   document.getElementById('root') as HTMLElement
// );
// root.render(
//   <React.StrictMode>
//     <App />
//   </React.StrictMode>
// );

