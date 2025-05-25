package domain

import (
	"canchitas-libres-transaction/internal/configuration"
	"context"
)

type StorageRepository interface {
	GetAll(ctx context.Context) ([]Transaction, error)
	GetByID(id int) (Transaction, error)
	Add(ctx context.Context, transaction Transaction) error
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id int, transaction Transaction) error
}

type Service struct {
	Config            *configuration.Configuration
	StorageRepository StorageRepository
}

func NewService(config *configuration.Configuration, storageRepository StorageRepository) *Service {
	return &Service{
		Config:            config,
		StorageRepository: storageRepository,
	}
}
