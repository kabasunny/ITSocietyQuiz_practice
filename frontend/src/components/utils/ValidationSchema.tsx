import { z } from 'zod';

// ログイン用バリデーションスキーマ
export const validationSchema = z.object({
  empid: z
    .string()
    .min(1, '社員IDは必須です (ꐦ°᷄д°᷅)')
    .min(6, '社員IDは6文字以上で入力してください (ꐦ`•ω•´)'),
  password: z
    .string()
    .min(1, 'パスワードは必須です (°ㅂ°ꐦ)')
    .min(6, 'パスワードは6文字以上で入力してください (# `꒳´ )'),
});


// ユーザー管理用バリデーションスキーマを定義
export const addUserSchema = z.object({
  empId: z
  .string()
  .min(1, '社員IDは必須です (ꐦ°᷄д°᷅)')
  .min(6, '社員IDは6文字以上で入力してください (ꐦ`•ω•´)'),
  name: z
  .string()
  .min(1, '社員氏名は必須です(ꐦ`•ω•´)'),
  email: z
  .string()
  .email('無効なメールアドレスです'),
  password_1: z
  .string()
  .min(6, 'パスワードは6文字以上必要です(ꐦ°᷄д°᷅)'),
  password_2: z
  .string()
  .min(6, 'パスワードは6文字以上必要です(ꐦ`•ω•´)'),
  roleId: z
  .number()
  .min(1, '権限を選択してください'),
}).refine(data => {
  // パスワードの一致を確認
  if (data.password_1 && data.password_2 && data.password_1 !== data.password_2) {
    return false;
  }
  return true;
}, {
  message: 'パスワードが一致しません (ꐦ°᷄д°᷅)',
  path: ['password_2'], // エラーメッセージを表示するフィールド;
});

export const editUserSchema = z.object({
  empId: z
  .string()
  .min(1, '社員IDは必須です (ꐦ°᷄д°᷅)')
  .min(6, '社員IDは6文字以上で入力してください (ꐦ`•ω•´)'),
  name: z
  .string()
  .min(1, '社員氏名は必須です(ꐦ`•ω•´)'),
  email: z
  .string()
  .email('無効なメールアドレスです'),
  password_1: z
    .string()
    .refine(value => value === '' || value.length >= 6, {
      message: 'パスワードは6文字以上必要です (ꐦ°᷄д°᷅)',
    })
    .optional(),
  password_2: z
    .string()
    .refine(value => value === '' || value.length >= 6, {
      message: 'パスワードは6文字以上必要です (ꐦ`•ω•´)',
    })
    .optional(),
  

});
