"use client";

import { Button, List, Paper, TextField } from "@mui/material";
import type React from "react";
import { useState } from "react";
import { z } from "zod";
import { type TaskResponse, TaskResponseSchema } from "../types/Task";
import TaskItem from "./TaskItem";

const initialTodos = [
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
const parsedTodos = z.array(TaskResponseSchema).safeParse(initialTodos);

// zodによるバリデーション
if (!parsedTodos.success) {
	console.error("初期データが無効:", parsedTodos.error.format());
	throw new Error("無効な初期データ");
}

const TaskList: React.FC = () => {
	const [todos, setTodos] = useState<TaskResponse[]>(parsedTodos.data);
	const [newTodo, setNewTodo] = useState<string>("");

	const handleAddTodo = () => {
		if (newTodo.trim() === "") {
			alert("TODOアイテムを入力してください。");
			return;
		}

		const newTaskItem: TaskResponse = {
			id: Date.now(),
			title: newTodo.trim(),
			detail: "",
			is_completed: false,
			created_at: new Date().toISOString(),
		};

		setTodos([...todos, newTaskItem]);
		setNewTodo("");
		alert(`追加したTODO: "${newTaskItem.title}"`);
	};

	const handleToggle = (id: number) => {
		setTodos(
			todos.map((todo) =>
				todo.id === id ? { ...todo, is_completed: !todo.is_completed } : todo,
			),
		);
	};

	const handleDelete = (id: number) => {
		setTodos(todos.filter((todo) => todo.id !== id));
	};

	return (
		<Paper style={{ padding: 16, maxWidth: 600, margin: "auto" }}>
			<h1>TODOリスト</h1>
			<div style={{ display: "flex", marginBottom: 16 }}>
				<TextField
					label="新しいTODO"
					variant="outlined"
					fullWidth
					value={newTodo}
					onChange={(e) => setNewTodo(e.target.value)}
				/>
				<Button
					variant="contained"
					color="primary"
					style={{ marginLeft: 8 }}
					onClick={handleAddTodo}
				>
					追加
				</Button>
			</div>
			<List>
				{todos.map((todo) => (
					<TaskItem
						key={todo.id}
						task={todo}
						onToggle={handleToggle}
						onDelete={handleDelete}
					/>
				))}
			</List>
		</Paper>
	);
};

export default TaskList;
