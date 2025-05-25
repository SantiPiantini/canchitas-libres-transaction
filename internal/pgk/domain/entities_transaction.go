package domain

type Transaction struct {
	TransactionID string `db:"transaction_id" json:"transaction_id,omitempty"`
	PaymentID     string `db:"payment_id" json:"payment_id"`
	UserID        string `db:"user_id" json:"user_id"`
}

func NewTransaction(paymentID, userID string) *Transaction {
	return &Transaction{
		PaymentID: paymentID,
		UserID:    userID,
	}
}
