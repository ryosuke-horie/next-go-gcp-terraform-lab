package repository

import (
	"context"

	"github.com/ryosuke-horie/next-go-gcp-terraform-k8s-lab/models"
)

// TaskRepository はタスクに対するDB操作のインターフェース
type TaskRepository interface {
	CreateTask(ctx context.Context, task *models.Task) error
	ListTasks(ctx context.Context) ([]models.Task, error)
	DeleteTask(ctx context.Context, id int) error
}
