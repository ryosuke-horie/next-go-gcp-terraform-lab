import TaskList from "@/components/TaskList";
import { z } from "zod";
import { type TaskResponse, TaskResponseSchema } from "../types/Task";

const Home = async () => {
	// サーバーサイドでデータをフェッチ
	const response = await fetch("http://localhost:3333/task", {
		// 再生成を防ぐためのオプション
		cache: "no-store",
	});

	if (!response.ok) {
		throw new Error("タスクの取得に失敗しました。");
	}

	const data = await response.json();

	// Zodによるデータ検証
	const parsed = z.array(TaskResponseSchema).safeParse(data);

	if (!parsed.success) {
		console.error("取得したデータが無効です:", parsed.error.format());
		throw new Error("取得したデータの形式が無効です。");
	}

	const tasks: TaskResponse[] = parsed.data;

	return (
		<main>
			<TaskList initialTasks={tasks} />
		</main>
	);
};

export default Home;
