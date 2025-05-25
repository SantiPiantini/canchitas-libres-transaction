package domain

import "context"

func (s *Service) GetAll(ctx context.Context) ([]Transaction, error) {
	return s.StorageRepository.GetAll(ctx)
}
