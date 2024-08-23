// Users.ts
export interface User {
  employeeId: string; // 型を string に変更
  username: string;
  email: string;
  password: string;
  total_questions: number;
  correct_answers: number;
}

export const users: User[] = [
  {
    employeeId: "EMP1234", // 例: アルファベットを含む社員ID
    username: "ITSocietyQuiz",
    email: "quize@example.com",
    password: "password",
    total_questions: 10,
    correct_answers: 8,
  },
  {
    employeeId: "EMP2345", // 例: アルファベットを含む社員ID
    username: "ITSocietyQuiz_2",
    email: "quize_2@example.com",
    password: "password_2",
    total_questions: 15,
    correct_answers: 12,
  },
];
