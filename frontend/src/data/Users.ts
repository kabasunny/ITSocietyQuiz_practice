// Users.ts
export interface User {
    id: number;
    username: string;
    email: string;
    total_questions: number;
    correct_answers: number;
  }
  
  export const users: User[] = [
    {
      id: 1,
      username: "ユーザー1",
      email: "user1@example.com",
      total_questions: 10,
      correct_answers: 8,
    },
    {
      id: 2,
      username: "ユーザー2",
      email: "user2@example.com",
      total_questions: 15,
      correct_answers: 12,
    },
  ];
  