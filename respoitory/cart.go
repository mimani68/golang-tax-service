package repository

import (
	"context"
	"interview/domain"

	"gorm.io/gorm"
)

type cartRepository struct {
	database *gorm.DB
}

func NewCartRepository(db *gorm.DB) domain.CartRepository {
	return &cartRepository{
		database: db,
	}
}

func (tr *cartRepository) Create(c context.Context, cart *domain.CartEntity) error {
	return nil
}

func (tr *cartRepository) FindBySessionId(c context.Context, session_id string) (domain.CartEntity, error) {
	var value domain.CartEntity
	query := "status = ? AND session_id = ?"
	_ = tr.database.Where(query, domain.CartOpen, session_id).First(&value)
	return value, nil
}

func (tr *cartRepository) FindByProductName(c context.Context, cart_id uint, product_name string) (domain.CartItemEntity, error) {
	var value domain.CartItemEntity
	query := " cart_id = ? and product_name  = ?"
	_ = tr.database.Where(query, cart_id, product_name).First(&value)
	return value, nil
}

func (tr *cartRepository) Delete(c context.Context, cart *domain.CartEntity) error {
	return nil
}
