import { useState, useEffect } from 'react';
import axios from 'axios';
import { AdmQuestion } from '../../../types';

const useQuestions = (jwt: string | null) => {
  const [questions, setQuestions] = useState<AdmQuestion[]>([]);
  const [editingQuestion, setEditingQuestion] = useState<AdmQuestion | null>(null);

  useEffect(() => {
    if (!jwt) {
      console.error('Token not found');
      return;
    }

    axios.get('http://localhost:8082/admins', {
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    })
      .then((response) => {
        setQuestions(response.data.adm_data);
      })
      .catch((error) => {
        console.error('Error fetching data:', error);
      });
  }, [jwt, editingQuestion]); // editingQuestionに依存させることで、編集後に更新させ、データ抜けを解消

  return { questions, setQuestions, editingQuestion, setEditingQuestion };
};

export default useQuestions;
