'use client'
import TaskItem from "../components/TaskItem";
import { TaskResponse } from "@/types/Task";

const mockTaskItem: TaskResponse = {
    id: 1,
    title: "test task",
    detail: "task detail",
    is_completed: false,
    created_at: "2024_01_01"
} 

const ToggleTask = () => {
    alert("Toggle")
}

const DeleteTask = () => {
    alert("Delete")
}

export default function Home() {
	return (
		<main>
            <TaskItem 
            task={mockTaskItem}
            onToggle={ToggleTask}
            onDelete={DeleteTask}
            />
		</main>
	);
}
