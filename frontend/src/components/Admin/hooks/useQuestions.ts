import { useState, useEffect, useCallback } from 'react';
import axios from 'axios';
import { AdmQuestion } from '../../../types';

const useQuestions = (jwt: string | null) => {
  const [questions, setQuestions] = useState<AdmQuestion[]>([]);
  const [editingQuestion, setEditingQuestion] = useState<AdmQuestion | null>(null);

  const refreshQuestions = useCallback(() => {
    if (!jwt) {
      console.error('Token not found');
      return;
    }

    axios
      .get('http://localhost:8082/admins/questionsdata/all', {
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
  }, [jwt]);

  useEffect(() => {
    refreshQuestions();
  }, [refreshQuestions]);

  return { questions, setQuestions, editingQuestion, setEditingQuestion, refreshQuestions };
};

export default useQuestions;
