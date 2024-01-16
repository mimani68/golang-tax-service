package repository

import (
	"context"
	"fmt"
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
	query := fmt.Sprintf("status = '%s' AND session_id = '%s'", domain.CartOpen, session_id)
	_ = tr.database.Where(query).First(&value)
	return value, nil
}

func (tr *cartRepository) FindByProductName(c context.Context, cart_id uint, product_name string) (domain.CartItemEntity, error) {
	var value domain.CartItemEntity
	query := " cart_id = ? and product_name  = ?"
	_ = tr.database.Where(query, cart_id, product_name).First(&value)
	return value, nil
	// pu.cartRepository.Where(" cart_id = ? and product_name  = ?", cartdomain.ID, addItemForm.Product).First(&cartItemEntity)
}

func (tr *cartRepository) Delete(c context.Context, cart *domain.CartEntity) error {
	return nil
}
