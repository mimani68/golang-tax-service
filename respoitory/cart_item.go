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

func (tr *cartItemRepository) FindById(c context.Context, id int) (domain.CartItemEntity, error) {
	return domain.CartItemEntity{}, nil
}

func (tr *cartItemRepository) FindByCartId(c context.Context, id int) ([]domain.CartItemEntity, error) {
	// result = pu.cartItemRepository.Where(fmt.Sprintf("cart_id = %d", cartEntity.ID)).Find(&cartItems)
	return []domain.CartItemEntity{}, nil
}

func (tr *cartItemRepository) Delete(c context.Context, id int) error {
	return nil
}

func (tr *cartItemRepository) Save(c context.Context, cartItemEntity domain.CartItemEntity) error {
	return nil
}
