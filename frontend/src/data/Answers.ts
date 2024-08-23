// Answers.ts
export interface Answer {
    id: number;
    user_id: number;
    question_id: number;
    answer: number;
    is_correct: boolean;
    timestamp: string;
  }
  
  export const answers: Answer[] = [
    {
      id: 1,
      user_id: 1,
      question_id: 1,
      answer: 2,
      is_correct: true,
      timestamp: "2024-08-22T12:00:00Z",
    },
    {
      id: 2,
      user_id: 2,
      question_id: 2,
      answer: 3,
      is_correct: false,
      timestamp: "2024-08-22T12:05:00Z",
    },
  ];
  