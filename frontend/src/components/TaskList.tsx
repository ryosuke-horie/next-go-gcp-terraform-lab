"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { Button, List, Paper, TextField } from "@mui/material";
import type React from "react";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import {
	type NewTask,
	NewTaskSchema,
	type TaskResponse,
	TaskResponseSchema,
} from "../types/Task";
import TaskItem from "./TaskItem";

const initialTasks = [
	{
		id: 1,
		title: "サンプルTODO 1",
		detail: "詳細1",
		is_completed: false,
		created_at: new Date().toISOString(),
	},
	{
		id: 2,
		title: "サンプルTODO 2",
		detail: "詳細2",
		is_completed: true,
		created_at: new Date().toISOString(),
	},
];

// 検証
const parsedTasks = z.array(TaskResponseSchema).safeParse(initialTasks);

// zodによるバリデーション
if (!parsedTasks.success) {
	console.error("初期データが無効:", parsedTasks.error.format());
	throw new Error("無効な初期データ");
}

const TaskList: React.FC = () => {
	const [tasks, setTasks] = useState<TaskResponse[]>(parsedTasks.data);

	// RHFの初期化
	const {
		register,
		handleSubmit,
		reset,
		formState: { errors, isSubmitting },
	} = useForm<NewTask>({
		resolver: zodResolver(NewTaskSchema),
	});

	// フォーム送信時の処理
	const onSubmit = (data: NewTask) => {
		const newTaskItem: TaskResponse = {
			id: 3, // 本来は＋1する
			title: data.title,
			detail: "", // 未実装のため空欄
			is_completed: false,
			created_at: new Date().toISOString(),
		};

		setTasks((prevTasks) => [...prevTasks, newTaskItem]);
		reset(); // フォームをリセット
		alert(`追加したTodo: "${newTaskItem.title}"`);
	};

	const handleToggle = (id: number) => {
		setTasks(
			tasks.map((task) =>
				task.id === id ? { ...task, is_completed: !task.is_completed } : task,
			),
		);
	};

	const handleDelete = (id: number) => {
		setTasks(tasks.filter((task) => task.id !== id));
	};

	return (
		<Paper style={{ padding: 16, maxWidth: 600, margin: "auto" }}>
			<h1>TODOリスト</h1>
			<form
				onSubmit={handleSubmit(onSubmit)}
				style={{ display: "flex", marginBottom: 16 }}
			>
				<TextField
					label="新しいTODO"
					variant="outlined"
					fullWidth
					{...register("title")}
					error={!!errors.title}
					helperText={errors.title ? errors.title.message : ""}
				/>
				<Button
					type="submit"
					variant="contained"
					color="primary"
					style={{ marginLeft: 8 }}
					disabled={isSubmitting}
				>
					追加
				</Button>
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
