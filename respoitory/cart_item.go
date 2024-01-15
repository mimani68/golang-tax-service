package repository

import (
	"context"
	"interview/db"
	"interview/domain"
)

type cartItemRepository struct {
	database db.Database
}

func NewCartItemRepository(db db.Database) domain.CartItemRepository {
	return &cartItemRepository{
		database: db,
	}
}

func (tr *cartItemRepository) Create(c context.Context, cartItem *domain.CartItemEntity) error {
	return nil
}

func (tr *cartItemRepository) Where(c context.Context, criteria string) []domain.CartItemEntity {
	return []domain.CartItemEntity{}
}

func (tr *cartItemRepository) Delete(c context.Context, cartItem *domain.CartItemEntity) error {
	return nil
}
