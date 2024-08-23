// Users.ts
export interface User {
  empid: string; // 型を string に変更
  username: string;
  email: string;
  password: string;
  total_questions: number;
  correct_answers: number;
}

export const users: User[] = [
  {
    empid: "EMP1234",
    username: "ITSocietyQuiz",
    email: "quize@example.com",
    password: "password",
    total_questions: 10,
    correct_answers: 8,
  },
  {
    empid: "EMP2345",
    username: "ITSocietyQuiz_2",
    email: "quize_2@example.com",
    password: "password_2",
    total_questions: 15,
    correct_answers: 12,
  },
];
