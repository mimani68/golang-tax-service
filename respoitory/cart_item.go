package repository

import (
	"context"
	"interview/domain"

	"gorm.io/gorm"
)

type cartItemRepository struct {
	database *gorm.DB
}

func NewCartItemRepository(db *gorm.DB) domain.CartItemRepository {
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

func (tr *cartItemRepository) Save(c context.Context, CartItemEntity string) error {
	return nil
}
