package repository

import (
	"context"

	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/models"
)

type MockTaskRepository struct {
	// メソッドの呼び出しを追跡するためのフラグや結果を設定可能
	CreateTaskFunc func(ctx context.Context, task *models.Task) error
	ListTasksFunc  func(ctx context.Context) ([]models.Task, error)
	DeleteTaskFunc func(ctx context.Context, id int) error
}

func (m *MockTaskRepository) CreateTask(ctx context.Context, task *models.Task) error {
	if m.CreateTaskFunc != nil {
		return m.CreateTaskFunc(ctx, task)
	}
	return nil
}

func (m *MockTaskRepository) ListTasks(ctx context.Context) ([]models.Task, error) {
	if m.ListTasksFunc != nil {
		return m.ListTasksFunc(ctx)
	}
	return []models.Task{}, nil
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, id int) error {
	if m.DeleteTaskFunc != nil {
		return m.DeleteTaskFunc(ctx, id)
	}
	return nil
}
