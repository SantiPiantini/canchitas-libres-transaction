package domain

import "context"

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.StorageRepository.Delete(ctx, id)
}
