"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Button, List, Paper, Stack, TextField } from "@mui/material";
import type React from "react";
import { useForm } from "react-hook-form";
import useSWR, { mutate } from "swr";
import { z } from "zod";
import {
	type NewTask,
	NewTaskSchema,
	type TaskResponse,
	TaskResponseSchema,
} from "../types/Task";
import TaskItem from "./TaskItem"; // 修正後のコンポーネント名

interface TaskListProps {
	initialTasks: TaskResponse[];
}

// swrのフェッチ関数
const fetcher = async (url: string): Promise<TaskResponse[]> => {
	const response = await fetch(url);
	if (!response.ok) {
		throw new Error("タスクの取得に失敗しました。");
	}

	const data = await response.json();
	const parsed = z.array(TaskResponseSchema).safeParse(data);
	if (!parsed.success) {
		throw new Error("データ形式が無効です");
	}
	return parsed.data;
};

const TaskList: React.FC<TaskListProps> = ({ initialTasks }) => {
	// 環境変数からAPIのベースURLを取得
	const apiBaseUrl = process.env.NEXT_PUBLIC_API_BASE_URL;

	if (!apiBaseUrl) {
		throw new Error("NEXT_PUBLIC_API_BASE_URLが設定されていません。");
	}

	const { data: tasks, error } = useSWR<TaskResponse[]>(
		`${apiBaseUrl}/task`,
		fetcher,
		{
			fallbackData: initialTasks,
		},
	);

	// RHFの初期化
	const {
		register,
		handleSubmit,
		reset,
		formState: { errors, isSubmitting },
	} = useForm<NewTask>({
		resolver: zodResolver(NewTaskSchema),
	});

	// ローディングおよびエラーステートの処理
	if (error) return <div>タスクの読み込みに失敗しました。</div>;
	if (!tasks) return <div>読み込み中...</div>;

	// フォーム送信時の処理
	const onSubmit = async (data: NewTask) => {
		const newTaskItem: TaskResponse = {
			id: Date.now(), // 一意のIDを生成
			title: data.title.trim(),
			detail: data.detail?.trim() || "",
			is_completed: false,
			created_at: new Date().toISOString(),
		};

		try {
			const response = await fetch(`${apiBaseUrl}/task`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(newTaskItem),
			});
			if (!response.ok) {
				const errorData: { message?: string } = await response.json();
				throw new Error(errorData.message || "タスク追加に失敗しました");
			}
		} catch (error: any) {
			// エラーメッセージを取得するため any に
			alert(error.message || "タスク作成に失敗しました");
		}

		// swrキャッシュの再検証（最新を取得）
		mutate(`${apiBaseUrl}/task`);
		reset(); // フォームをリセット
	};

	// 更新処理
	const handleToggle = async (id: number) => {
		try {
			const task = tasks.find((task) => task.id === id);
			if (!task) {
				throw new Error("タスクが見つかりません");
			}

			const updatedTask = {
				...task,
				is_completed: !task.is_completed,
			};

			console.log("Updating task:", updatedTask);

			const response = await fetch(`${apiBaseUrl}/task`, {
				method: "PUT",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(updatedTask),
			});

			console.log("Response status:", response.status);

			if (!response.ok) {
				const errorData: { message?: string } = await response.json();
				throw new Error(errorData.message || "タスクの更新に失敗しました");
			}

			// SWRのキャッシュを再検証
			mutate(`${apiBaseUrl}/task`);
		} catch (error: any) {
			// エラーメッセージを取得するため any に
			console.error("Error updating task:", error);
			alert(error.message || "タスクの更新に失敗しました");
		}
	};

	const handleDelete = async (id: number) => {
		try {
			const response = await fetch(`${apiBaseUrl}/task`, {
				method: "DELETE",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify({ id: id }),
			});

			if (!response.ok) {
				const errorData: { message?: string } = await response.json();
				throw new Error(errorData.message || "タスク削除に失敗しました");
			}

			// SWRのキャッシュを再検証
			mutate(`${apiBaseUrl}/task`);
		} catch (error: any) {
			// エラーメッセージを取得するため any に
			console.error(error);
			alert(error.message || "失敗しました。");
		}
	};

	return (
		<Paper style={{ padding: 16, maxWidth: 600, margin: "auto" }}>
			<h1>TODOリスト</h1>
			<form onSubmit={handleSubmit(onSubmit)}>
				<Stack spacing={2} marginBottom={2}>
					{/* タイトル入力フィールド */}
					<TextField
						label="新しいTODO"
						variant="outlined"
						fullWidth
						{...register("title")}
						error={!!errors.title}
						helperText={errors.title ? errors.title.message : ""}
					/>
					{/* 詳細入力フィールド（テキストエリア） */}
					<TextField
						label="詳細"
						variant="outlined"
						fullWidth
						multiline
						rows={4}
						{...register("detail")}
						error={!!errors.detail}
						helperText={errors.detail ? errors.detail.message : ""}
					/>
					{/* 追加ボタン */}
					<Button
						type="submit"
						variant="contained"
						color="primary"
						disabled={isSubmitting}
					>
						追加
					</Button>
				</Stack>
			</form>
			<List>
				{tasks.map((task) => (
					<TaskItem
						key={task.id}
						task={task}
						onToggle={handleToggle}
						onDelete={handleDelete}
					/>
				))}
			</List>
		</Paper>
	);
};

export default TaskList;
