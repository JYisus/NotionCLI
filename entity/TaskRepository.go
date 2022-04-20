package entity

import "context"

type TaskRepository interface {
	AddTask(ctx context.Context, database Database, task string) error
	DeleteTask(ctx context.Context, task string) error
	ListTasks(ctx context.Context, database Database) ([]Task, error)
}
