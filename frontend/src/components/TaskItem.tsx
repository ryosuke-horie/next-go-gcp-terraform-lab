"use client";

import { Checkbox, ListItem, ListItemText } from "@mui/material";
import Button from "@mui/material/Button";
import type React from "react";
import type { TaskResponse } from "../types/Task";

interface TaskItemProps {
	task: TaskResponse;
	onToggle: (id: number) => void;
	onDelete: (id: number) => void;
}

const TaskItem: React.FC<TaskItemProps> = ({ task, onToggle, onDelete }) => {
	const handleToggle = () => {
		onToggle(task.id);
	};

	const handleDelete = () => {
		onDelete(task.id);
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

export default TaskItem;
