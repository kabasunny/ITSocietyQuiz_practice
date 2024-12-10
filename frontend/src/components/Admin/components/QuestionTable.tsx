import React from 'react';
import { AdmQuestion } from '../../../types';
import { formatDate } from '../utils/formatDate';

interface QuestionTableProps {
  questions: AdmQuestion[];
  handleEdit: (question: AdmQuestion) => void;
  handleDelete: (id: number) => void; // handleDelete プロップを追加
}

const QuestionTable: React.FC<QuestionTableProps> = ({ questions, handleEdit, handleDelete }) => {
  // questions配列をコピーしてid順にソート...なぜか編集後にソートがうまく行かない⁉
  const sortedQuestions = [...questions].sort((a, b) => {
    return Number(a.id) - Number(b.id);
  });
  console.log(sortedQuestions)

  return (
    <table className="admin-table">
      <thead>
        <tr>
          <th>DB ID</th>
          <th>Qiz ID</th>
          <th>問題文</th>
          <th>正解</th>
          <th>捕捉情報</th>
          <th>不正解1</th>
          <th>不正解2</th>
          <th>不正解3</th>
          <th>難度</th>
          <th>最終更新</th>
          <th>Act</th>
        </tr>
      </thead>
      <tbody>
        {sortedQuestions.map((question) => (
          <tr key={question.id}>
            <td>{question.id}</td>
            <td>{question.userQuestionID || ''}</td>
            <td>{question.question || ''}</td>
            <td>{question.options?.[0] || ''}</td>
            <td>{question.supplement || ''}</td>
            <td>{question.options?.[1] || ''}</td>
            <td>{question.options?.[2] || ''}</td>
            <td>{question.options?.[3] || ''}</td>
            <td>{question.difficulty !== undefined ? question.difficulty.toString() : ''}</td>
            <td dangerouslySetInnerHTML={{ __html: formatDate(question.updatedAt > question.createdAt ? question.updatedAt : question.createdAt) }}></td>
            <td>
              <div className="button-container">
                <button className="edit-button" onClick={() => handleEdit(question)}>編集</button><br/>
                <button className="delete-button" onClick={() => handleDelete(question.id)}>削除</button>
              </div>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default QuestionTable;
