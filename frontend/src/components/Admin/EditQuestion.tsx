import React, { useState } from 'react';
import axios from 'axios';
import useQuestions from './hooks/useQuestions';
import QuestionTable from './components/QuestionTable';
import EditForm from './components/EditForm';
import { AdmQuestion } from '../../types';
import './css/AdminScreen.css';
import './css/EditQuestion.css';

const EditQuestion: React.FC = () => {
  const jwt = sessionStorage.getItem('token');
  const { questions, setQuestions, editingQuestion, setEditingQuestion } = useQuestions(jwt);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 20;

  // 入力フィールドの変更を処理する関数
  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement>, // イベントオブジェクトの型を指定
    field: keyof AdmQuestion, // 変更するフィールドのキーを指定
    index?: number // オプションのインデックス（オプションフィールドの場合に使用）
  ) => {
    if (editingQuestion) { // 編集中の質問が存在する場合に処理を行う
      if (field === 'options' && index !== undefined) { // フィールドがオプションで、インデックスが指定されている場合
        const updatedOptions = [...editingQuestion.options]; // 現在のオプションをコピー
        updatedOptions[index] = e.target.value; // 指定されたインデックスのオプションを更新
        setEditingQuestion({ ...editingQuestion, options: updatedOptions }); // 更新されたオプションを含む新しい質問オブジェクトを設定
      } else if (field === 'difficulty') { // 難度フィールドの場合
        setEditingQuestion({ ...editingQuestion, [field]: parseInt(e.target.value, 10) }); // 数値に変換して設定
      } else { // その他のフィールドの場合
        setEditingQuestion({ ...editingQuestion, [field]: e.target.value }); // 指定されたフィールドを更新
      }
    }
  };
  
  // 編集中の質問を保存する関数
  const handleSave = (id: number) => {
    if (editingQuestion) {
      axios.put(`http://localhost:8082/admins/${id}`, editingQuestion, { // 編集中の質問を指定されたIDで更新するためのPUTリクエストを送信
        headers: {
          Authorization: `Bearer ${jwt}`, // 認証トークンをヘッダーに含める
        },
      })
        .then((response) => {
          setQuestions(
            questions.map((q) => (q.id === id ? response.data : q)) // 更新された質問データで質問リストを更新
          );
          setEditingQuestion(null); // 編集中の質問をクリア
        })
        .catch((error) => {
          console.error('Error updating data:', error); // エラーメッセージをコンソールに表示
        });
    }
  };

  // 質問を削除する関数
const handleDelete = (id: number) => {
  const confirmDelete = window.confirm('本当にこの質問を削除しますか？'); // 確認ダイアログを表示
  if (confirmDelete) {
    axios.delete(`http://localhost:8082/admins/${id}`, { // 指定されたIDの質問を削除するためのDELETEリクエストを送信
      headers: {
        Authorization: `Bearer ${jwt}`, // 認証トークンをヘッダーに含める
      },
    })
      .then(() => {
        setQuestions(questions.filter((q) => q.id !== id)); // 削除された質問を質問リストから除外
      })
      .catch((error) => {
        console.error('Error deleting data:', error); // エラーメッセージをコンソールに表示
      });
  }
};


  // 質問の編集を開始する関数
  const handleEdit = (question: AdmQuestion) => {
    setEditingQuestion(question); // 編集中の質問を設定
  };

  // 編集をキャンセルする関数
  const handleCancelEdit = () => {
    setEditingQuestion(null); // 編集中の質問をクリア
  };

  // ページを変更する関数
  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber); // 現在のページ番号を設定
  };

  const indexOfLastItem = currentPage * itemsPerPage; // 現在のページの最後のアイテムのインデックスを計算
  const indexOfFirstItem = indexOfLastItem - itemsPerPage; // 現在のページの最初のアイテムのインデックスを計算
  const currentItems = questions.slice(indexOfFirstItem, indexOfLastItem); // 現在のページに表示するアイテムを取得

return (
  <div className="admin-container">
    <h2 className="admin-h2">問題の編集（クイズ問題の更新、削除）</h2>
    <div className="admin-table-container">
      <QuestionTable questions={currentItems} handleEdit={handleEdit} handleDelete={handleDelete} />
    </div>
    {editingQuestion ? (
      <EditForm
        editingQuestion={editingQuestion}
        handleInputChange={handleInputChange}
        handleSave={handleSave}
        handleCancelEdit={handleCancelEdit}
      />
    ) : (
      <div className="pagination">
        {Array.from({ length: Math.ceil(questions.length / itemsPerPage) }, (_, index) => (
          // Array.from メソッドを使用して、ページ番号のボタンを動的に生成
          // length: Math.ceil(questions.length / itemsPerPage) は、質問の総数を1ページあたりのアイテム数で割った値を切り上げて、必要なページ数を計算
          <button
            key={index + 1 // Reactでリストをレンダリングする際に必要な一意のキーを設定
            }
            onClick={() => handlePageChange(index + 1)}
            className={currentPage === index + 1 ? 'active' : '' // 現在のページのボタンの色が変わる
            } 
          >
            {index + 1}
          </button>
        ))}
      </div>
    )}
  </div>
);
};

export default EditQuestion;
