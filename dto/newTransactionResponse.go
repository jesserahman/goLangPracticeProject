package dto

type NewTransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	Balance       float64 `json:"balance"`
}
