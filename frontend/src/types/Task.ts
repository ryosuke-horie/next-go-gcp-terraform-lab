import { z } from "zod";

// zodスキーマ
export const TaskResponseSchema = z.object({
	id: z.number().positive(), // 正の数
	title: z.string().min(1, "タイトルは必須です"), // 必須で1文字以上
	detail: z.string().optional(), // 任意の文字列
	is_completed: z.boolean(), // 真偽値
	created_at: z.string().datetime(), // ISO形式の日時文字列
});

// 型定義をスキーマから生成
export type TaskResponse = z.infer<typeof TaskResponseSchema>;
