package repository

import (
	"context"
	"interview/domain"
)

type cartRepository struct {
	database   string
	collection string
}

func NewCartRepository(db string, collection string) domain.CartRepository {
	return &cartRepository{
		database:   db,
		collection: collection,
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
