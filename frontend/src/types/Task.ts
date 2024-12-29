export interface TaskResponse {
    id: number;
    title: string;
    detail: string;
    is_completed: boolean;
    created_at: string; // ISO形式
}