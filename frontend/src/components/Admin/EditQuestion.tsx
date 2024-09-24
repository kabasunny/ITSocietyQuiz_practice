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
  const itemsPerPage = 50;

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
    } else { // その他のフィールドの場合
      setEditingQuestion({ ...editingQuestion, [field]: e.target.value }); // 指定されたフィールドを更新
    }
  }
};


  // 編集中の質問を保存する関数
  const handleSave = (id: number) => {
    if (editingQuestion) {
      axios.put(`http://localhost:8082/admins/${id}`, editingQuestion, {
        headers: {
          Authorization: `Bearer ${jwt}`,
        },
      })
        .then((response) => {
          setQuestions(
            questions.map((q) => (q.id === id ? response.data : q))
          );
          setEditingQuestion(null);
        })
        .catch((error) => {
          console.error('Error updating data:', error);
        });
    }
  };

  // 質問を削除する関数
  const handleDelete = (id: number) => {
    axios.delete(`http://localhost:8082/admins/${id}`, {
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    })
      .then(() => {
        setQuestions(questions.filter((q) => q.id !== id));
      })
      .catch((error) => {
        console.error('Error deleting data:', error);
      });
  };

  // 質問の編集を開始する関数
  const handleEdit = (question: AdmQuestion) => {
    setEditingQuestion(question);
  };

  // 編集をキャンセルする関数
  const handleCancelEdit = () => {
    setEditingQuestion(null);
  };

  // ページを変更する関数
  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = questions.slice(indexOfFirstItem, indexOfLastItem);

  return (
    <div className="admin-container">
      <h2 className="admin-h2">問題の編集（クイズ問題の更新、削除）</h2>
      <div className="admin-table-container">
        <QuestionTable questions={currentItems} handleEdit={handleEdit} handleDelete={handleDelete} />
      </div>
      <div className="pagination">
        {Array.from({ length: Math.ceil(questions.length / itemsPerPage) }, (_, index) => (
          <button
            key={index + 1}
            onClick={() => handlePageChange(index + 1)}
            className={currentPage === index + 1 ? 'active' : ''}
          >
            {index + 1}
          </button>
        ))}
      </div>
      {editingQuestion && (
        <EditForm
          editingQuestion={editingQuestion}
          handleInputChange={handleInputChange}
          handleSave={handleSave}
          handleCancelEdit={handleCancelEdit}
        />
      )}
    </div>
  );
};

export default EditQuestion;
