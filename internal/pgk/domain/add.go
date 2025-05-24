package domain

import "context"

func (s *Service) Add(ctx context.Context, transaction Transaction) error {
	return s.StorageRepository.Add(ctx, transaction)
}
