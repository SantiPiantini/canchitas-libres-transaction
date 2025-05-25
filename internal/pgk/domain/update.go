package domain

import "context"

func (s *Service) Update(ctx context.Context, id string, transaction Transaction) error {
	return s.StorageRepository.Update(ctx, id, transaction)
}
