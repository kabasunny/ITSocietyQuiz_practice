import React from 'react';
import { EditFormProps } from '../../../types';

const EditForm: React.FC<EditFormProps> = ({
  editingQuestion,
  handleInputChange,
  handleSave,
  handleCancelEdit,
}) => {
  return (
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
          value={editingQuestion.difficulty !== undefined ? editingQuestion.difficulty.toString() : '' //表する際はstring型でデータとして扱うときは数値にする
          }
          onChange={(e) => handleInputChange(e, 'difficulty')}
        />
      </label>
      <button onClick={() => handleSave(editingQuestion.id)}>保存 & 更新</button>
      <button onClick={handleCancelEdit}>編集取止め</button>
    </div>
  );
};

export default EditForm;
