import { z } from 'zod';

export const validationSchema = z.object({
  empid: z
    .string()
    .min(1, "社員IDは必須です (ꐦ°᷄д°᷅)")
    .min(6, "社員IDは6文字以上で入力してください (ꐦ`•ω•´)"),
  name: z
    .string()
    .min(1, "名前は必須です (╬＾ω＾)")
    .min(4, "名前は4文字以上で入力してください ( ꐦ◜ω◝ )"),
  email: z
    .string()
    .min(1, "メールアドレスは必須です (੭ꐦ •̀Д•́ )੭*")
    .email("正しいメールアドレスを入力してください (#°Д°)"),
  password: z
    .string()
    .min(1, "パスワードは必須です (°ㅂ°ꐦ)")
    .min(6, "パスワードは6文字以上で入力してください (# `꒳´ )"),
});
