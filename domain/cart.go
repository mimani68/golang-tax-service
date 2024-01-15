package domain

import (
	"context"

	"gorm.io/gorm"
)

const (
	CartOpen   = "open"
	CartClosed = "closed"
)

type CartEntity struct {
	gorm.Model
	Total     float64
	SessionID string
	Status    string
}

type CartUsecase interface {
	// Create(c context.Context, task *Task) error
	// FetchByUserID(c context.Context, userID string) ([]Task, error)
}

type CartRepository interface {
	Create(c context.Context, cart *CartEntity) error
	Where(c context.Context, criteria string) []Task
	Delete(c context.Context, cart *CartEntity) error
}
