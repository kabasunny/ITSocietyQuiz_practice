import { z } from 'zod';

export const validationSchema = z.object({
  empid: z
    .string()
    .min(1, "社員IDは必須です (ꐦ°᷄д°᷅)")
    .min(6, "社員IDは6文字以上で入力してください (ꐦ`•ω•´)"),
  password: z
    .string()
    .min(1, "パスワードは必須です (°ㅂ°ꐦ)")
    .min(6, "パスワードは6文字以上で入力してください (# `꒳´ )"),
});
