package domain

import "context"

func (s *Service) GetByID(ctx context.Context, id string) (Transaction, error) {
	return s.StorageRepository.GetByID(ctx, id)
}
