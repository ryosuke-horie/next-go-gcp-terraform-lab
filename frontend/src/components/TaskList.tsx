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
import TaskItem from "./TaskItem";

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
	const { data: tasks, error } = useSWR<TaskResponse[]>(
		"http://localhost:3333/task",
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
	const onSubmit = (data: NewTask) => {
		const newTaskItem: TaskResponse = {
			id: Date.now(), // 一意のIDを生成
			title: data.title.trim(),
			detail: data.detail?.trim() || "",
			is_completed: false,
			created_at: new Date().toISOString(),
		};

		// swrキャッシュの再検証（最新を取得）
		mutate("http://localhost:3333/task");

		reset(); // フォームをリセット
		alert(`追加したTodo: "${newTaskItem.title}"`);
	};

	const handleToggle = (id: number) => {
		alert(id);
	};

	const handleDelete = (id: number) => {
		alert(id);
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
