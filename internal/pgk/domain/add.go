package domain

import "context"

func (s *Service) Add(ctx context.Context, transaction Transaction) (string, error) {
	return s.StorageRepository.Add(ctx, transaction)
}
