package repository

import (
	"context"
	"interview/db"
	"interview/domain"
)

type cartRepository struct {
	database db.Database
}

func NewCartRepository(db db.Database) domain.CartRepository {
	return &cartRepository{
		database: db,
	}
}

func (tr *cartRepository) Create(c context.Context, cart *domain.CartEntity) error {
	return nil
}

func (tr *cartRepository) Where(c context.Context, criteria string) []domain.CartEntity {
	return []domain.CartEntity{}
}

func (tr *cartRepository) Delete(c context.Context, cart *domain.CartEntity) error {
	return nil
}
