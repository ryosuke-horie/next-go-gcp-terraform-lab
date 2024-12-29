"use client";

import { Checkbox, IconButton, ListItem, ListItemText } from "@mui/material";
import Button from "@mui/material/Button";
import type React from "react";
import type { TaskResponse } from "../types/Task";

interface TodoItemProps {
	task: TaskResponse;
	onToggle: (id: number) => void;
	onDelete: (id: number) => void;
}

const TodoItem: React.FC<TodoItemProps> = ({ task, onToggle, onDelete }) => {
	const handleToggle = () => {
		onToggle(task.id);
		alert(`Toggled TODO with id: ${task.id}`);
	};

	const handleDelete = () => {
		onDelete(task.id);
		alert(`Deleted TODO with id: ${task.id}`);
	};

	return (
		<ListItem
			secondaryAction={
				<Button variant="contained" onClick={handleDelete}>
					削除
				</Button>
			}
		>
			<Checkbox
				edge="start"
				checked={task.is_completed}
				tabIndex={-1}
				disableRipple
				onChange={handleToggle}
			/>
			<ListItemText
				primary={task.title}
				secondary={task.detail}
				style={{ textDecoration: task.is_completed ? "line-through" : "none" }}
			/>
		</ListItem>
	);
};

export default TodoItem;
