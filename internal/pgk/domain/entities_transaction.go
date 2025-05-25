package domain

type Transaction struct {
	TransactionID string `db:"transaction_id" json:"transaction_id"`
	PaymentID     string `db:"payment_id" json:"payment_id, omitempty"`
	UserID        string `db:"user_id" json:"user_id"`
}

func NewTransaction(transactionID, paymentID, userID string) *Transaction {
	return &Transaction{
		TransactionID: transactionID,
		PaymentID:     paymentID,
		UserID:        userID,
	}
}
