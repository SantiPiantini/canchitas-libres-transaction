package domain

import (
	"canchitas-libres-transaction/internal/configuration"
	"context"
)

type StorageRepository interface {
	GetAll(ctx context.Context) ([]Transaction, error)
	GetByID(ctx context.Context, id string) (Transaction, error)
	Add(ctx context.Context, transaction Transaction) (string, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, transaction Transaction) error
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
