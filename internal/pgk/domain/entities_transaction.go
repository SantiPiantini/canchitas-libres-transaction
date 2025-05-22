package domain

type Transaction struct {
	id        int
	paymentID int
}

func NewTransaction(paymentID int) (*Transaction, error) {
	return &Transaction{
		paymentID: paymentID,
	}, nil
}
