package storage

import (
	"canchitas-libres-transaction/internal/pgk/domain"
	"context"
)

func (p *Postgres) GetAll() ([]domain.Transaction, error) {
	return []domain.Transaction{}, nil
}

func (p *Postgres) GetByID(id int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (p *Postgres) Add(ctx context.Context, transaction domain.Transaction) error {
	return nil
}

func (p *Postgres) Delete(ctx context.Context, id int) error {
	return nil
}

func (p *Postgres) Update(ctx context.Context, id int, transaction domain.Transaction) error {
	return nil
}
