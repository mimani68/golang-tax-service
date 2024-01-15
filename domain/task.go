package domain

import (
	"context"
)

const (
	CollectionTask = "tasks"
)

type Task struct {
	ID     string `bson:"_id" json:"-"`
	Title  string `bson:"title" form:"title" binding:"required" json:"title"`
	UserID string `bson:"userID" json:"-"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID string) ([]Task, error)
}
