package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/models"
)

// TaskRepositoryインターフェースの実装として定義
type TaskRepositoryImpl struct {
	DB *sql.DB
}

// NewTaskRepository は TaskRepositoryImpl のコンストラクタ
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{DB: db}
}

// DBにタスクを挿入する
func (r *TaskRepositoryImpl) CreateTask(ctx context.Context, task *models.Task) error {
	return task.Insert(ctx, r.DB)
}

// タスクの全件取得
func (r *TaskRepositoryImpl) ListTasks(ctx context.Context) ([]models.Task, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT * FROM tasks;`)
	if err != nil {
		log.Printf("ListTasks: DBクエリ失敗: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Detail, &task.IsCompleted, &task.CreatedAt); err != nil {
			log.Printf("ListTasks: 行のスキャン失敗: %v", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		log.Printf("ListTasks: rowsのイテレーション中にエラー: %v", err)
		return nil, err
	}

	return tasks, nil
}

// 指定したIDのタスクをDBから削除
func (r *TaskRepositoryImpl) DeleteTask(ctx context.Context, id int) error {
	// タスクを取得して存在確認
	task, err := models.TaskByID(ctx, r.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("タスクが見つかりません")
		}
	}

	return task.Delete(ctx, r.DB)
}
