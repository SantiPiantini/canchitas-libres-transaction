package domain

type Transaction struct {
	transactionID int
	paymentID     int
}

func NewTransaction(paymentID int) (*Transaction, error) {
	return &Transaction{
		paymentID: paymentID,
	}, nil
}
