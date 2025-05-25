package storage

import (
	"canchitas-libres-transaction/internal/pgk/domain"
	"context"
	"fmt"
)

func (p *Postgres) GetAll(ctx context.Context) ([]domain.Transaction, error) {
	fmt.Println("Ejecutando SELECT a la base de datos...")
	rows, err := p.DB.QueryContext(ctx,
		"SELECT transaction_id, payment_id, user_id FROM transactions")
	if err != nil {
		fmt.Println("Error en consulta:", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		var tx domain.Transaction
		if err := rows.Scan(&tx.TransactionID, &tx.PaymentID, &tx.UserID); err != nil {
			fmt.Println("Error escaneando fila:", err)
			return nil, err
		}
		fmt.Printf("Fila leída: %+v\n", tx) // 👈 Imprimimos cada fila leída
		transactions = append(transactions, tx)
	}

	fmt.Printf("Todas las transacciones leídas: %+v\n", transactions) // 👈 Imprimimos todo el array antes de devolverlo

	return transactions, nil
}

func (p *Postgres) GetByID(ctx context.Context, id string) (domain.Transaction, error) {
	var tx domain.Transaction
	err := p.DB.QueryRowContext(ctx, "SELECT transaction_id, payment_id, user_id FROM transactions WHERE transaction_id = $1::uuid", id).
		Scan(&tx.TransactionID, &tx.PaymentID, &tx.UserID)
	if err != nil {
		return domain.Transaction{}, err
	}
	return tx, nil
}

func (p *Postgres) Add(ctx context.Context, transaction domain.Transaction) error {
	_, err := p.DB.ExecContext(ctx,
		"INSERT INTO transactions (transaction_id, payment_id, user_id) VALUES ($1, $2, $3)",
		transaction.TransactionID,
		transaction.PaymentID,
		transaction.UserID,
	)
	return err
}

func (p *Postgres) Delete(ctx context.Context, id string) error {
	_, err := p.DB.ExecContext(ctx, "DELETE FROM transactions WHERE transaction_id = $1::uuid", id)
	return err
}

func (p *Postgres) Update(ctx context.Context, id string, transaction domain.Transaction) error {
	_, err := p.DB.ExecContext(ctx,
		"UPDATE transactions SET transaction_id = $1, payment_id = $2, user_id = $3 WHERE transaction_id = $4::uuid",
		transaction.TransactionID, transaction.PaymentID, transaction.UserID, id)
	return err
}
