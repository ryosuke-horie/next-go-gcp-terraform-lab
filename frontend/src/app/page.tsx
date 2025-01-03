import TaskList from "@/components/TaskList";
import { getCloudflareContext } from "@opennextjs/cloudflare";
import { z } from "zod";
import { type TaskResponse, TaskResponseSchema } from "../types/Task";

const Home = async () => {
	try {
		// Cloudflare Workersのコンテキストから環境変数を取得
		const { env } = await getCloudflareContext();
		const apiBaseUrl = env.NEXT_PUBLIC_API_BASE_URL;
		console.log(apiBaseUrl);

		if (!apiBaseUrl) {
			throw new Error("NEXT_PUBLIC_API_BASE_URLが設定されていません。");
		}

		// サーバーサイドでデータをフェッチ
		const response = await fetch(`${apiBaseUrl}/task`, {
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
	} catch (error) {
		console.error("サーバーサイドエラー:", error);
		throw error;
	}
};

export default Home;
