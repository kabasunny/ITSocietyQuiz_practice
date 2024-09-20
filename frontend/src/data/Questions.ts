// Questions.ts

import { Question } from '../types';
  
  export const questions: Question[] = [
    {
      id: 1, // GORM.Modelから
      question: "サンプル質問1",
      options: ["選択肢1", "選択肢2", "選択肢3", "選択肢4"],
      supplement: "補足情報1",
      difficulty: 1,
    },
    {
      id: 2, // 
      question: "サンプル質問2",
      options: ["選択肢A", "選択肢B", "選択肢C", "選択肢D"],
      supplement: "補足情報2",
      difficulty: 2,
    },
  ];
  