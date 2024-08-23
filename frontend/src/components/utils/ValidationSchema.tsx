import { z } from 'zod';

export const validationSchema = z.object({
  name: z
    .string()
    .min(1, "名前は必須です")
    .min(4, "名前は4文字以上で入力してください"),
  email: z
    .string()
    .min(1, "メールアドレスは必須です")
    .email("正しいメールアドレスを入力してください"),
  password: z
    .string()
    .min(1, "パスワードは必須です")
    .min(6, "パスワードは6文字以上で入力してください"),
});
