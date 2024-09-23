import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './css/AdminScreen.css'; // CSSファイルをインポート
import './css/EditQuestion.css'; // CSSファイルをインポート
import { AdmQuestion } from '../../types';

const EditQuestion: React.FC = () => {
  const [questions, setQuestions] = useState<AdmQuestion[]>([]);
  const [editingQuestion, setEditingQuestion] = useState<AdmQuestion | null>(null);
  const [currentPage, setCurrentPage] = useState(1);
  const itemsPerPage = 50;

  const jwt = sessionStorage.getItem('token'); // ログイン時にAPIから取得したトークン

  useEffect(() => {
    // トークンが存在するか確認
    if (!jwt) {
      console.error('Token not found');
      return;
    }

    // APIからクイズデータを取得
    axios.get('http://localhost:8082/admins', {
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    })
      .then((response) => {
        setQuestions(response.data.adm_data); // 修正箇所
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, [jwt, editingQuestion]);

  // ページネーションのロジック
  const indexOfLastItem = currentPage * itemsPerPage;
  const indexOfFirstItem = indexOfLastItem - itemsPerPage;
  const currentItems = questions.slice(indexOfFirstItem, indexOfLastItem);

  const handleInputChange = (
    e: React.ChangeEvent<HTMLInputElement>,
    field: keyof AdmQuestion,
    index?: number // オプションの index パラメータを追加
  ) => {
    if (editingQuestion) {
      if (field === 'options' && index !== undefined) {
        const updatedOptions = [...editingQuestion.options];
        updatedOptions[index] = e.target.value;
        setEditingQuestion({ ...editingQuestion, options: updatedOptions });
      } else {
        setEditingQuestion({ ...editingQuestion, [field]: e.target.value });
      }
    }
  };

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

  const handleEdit = (question: AdmQuestion) => {
    setEditingQuestion(question);
  };

  const handleCancelEdit = () => {
    setEditingQuestion(null);
  };

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月は0から始まるため+1
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    return `${year}年<br/>${month}月${day}日<br/>${hours}時${minutes}分`;
  };

  const handlePageChange = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  return (
    <div className="admin-container">
      <h2 className="admin-h2">問題の編集（クイズ問題の更新、削除）</h2>
      <div className="admin-table-container">
        <table className="admin-table">
          <thead>
            <tr>
              <th>DB ID</th>
              <th>Qiz ID</th>
              <th>問題文</th>
              <th>正解</th>
              <th>捕捉情報</th> {/*順番変えた*/}
              <th>不正解1</th>
              <th>不正解2</th>
              <th>不正解3</th>
              <th>難度</th>
              <th>作成日</th>
              <th>Act</th>
            </tr>
          </thead>
          <tbody>
            {Array.isArray(currentItems) && currentItems.map((question) => (
              <tr key={question.id}>
                <td>{question.id}</td>
                <td>{question.userQuestionID || ''}</td>
                <td>{question.question || ''}</td>
                <td>{question.options?.[0] || ''}</td>
                <td>{question.supplement || ''}</td>  {/*順番変えた*/}
                <td>{question.options?.[1] || ''}</td>
                <td>{question.options?.[2] || ''}</td>
                <td>{question.options?.[3] || ''}</td>
                <td>{question.difficulty !== undefined ? question.difficulty.toString() : ''}</td>
                <td dangerouslySetInnerHTML={{ __html: formatDate(question.createdAt) }}></td>
                <td>
                  <div className="button-container">
                    <button className="edit-button" onClick={() => handleEdit(question)}>編集</button>
                    <button className="delete-button">削除</button>
                  </div>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
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
        <div className="edit-form">
          <h2>問題の編集 ( DB_ID = {editingQuestion.id} )</h2>
          <label>
            Qiz ID:
            <input
              type="text"
              value={editingQuestion.userQuestionID || ''}
              onChange={(e) => handleInputChange(e, 'userQuestionID')}
            />
          </label>
          <label>
            問題文:
            <input
              type="text"
              value={editingQuestion.question}
              onChange={(e) => handleInputChange(e, 'question')}
            />
          </label>
          <label>
            正解:
            <input
              type="text"
              value={editingQuestion.options?.[0] || ''}
              onChange={(e) => handleInputChange(e, 'options', 0)}
            />
          </label>
          <label>
            捕捉情報:
            <input
              type="text"
              value={editingQuestion.supplement || ''}
              onChange={(e) => handleInputChange(e, 'supplement')}
            />
          </label>
          <label>
            不正解1:
            <input
              type="text"
              value={editingQuestion.options?.[1] || ''}
              onChange={(e) => handleInputChange(e, 'options', 1)}
            />
          </label>
          <label>
            不正解2:
            <input
              type="text"
              value={editingQuestion.options?.[2] || ''}
              onChange={(e) => handleInputChange(e, 'options', 2)}
            />
          </label>
          <label>
            不正解3:
            <input
              type="text"
              value={editingQuestion.options?.[3] || ''}
              onChange={(e) => handleInputChange(e, 'options', 3)}
            />
          </label>
          <label>
            難度:
            <input
              type="number"
              value={editingQuestion.difficulty !== undefined ? editingQuestion.difficulty.toString() : ''}
              onChange={(e) => handleInputChange(e, 'difficulty')}
            />
          </label>
          <button onClick={() => handleSave(editingQuestion.id)}>保存 & 更新</button>
          <button onClick={handleCancelEdit}>編集取止め</button>
        </div>
      )}
    </div>
  );
};

export default EditQuestion;
