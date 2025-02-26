package dto

type TransactionDTO struct {
	Amount          string `json:"amount"`
	FrequencyID     string `json:"frequency"`
	CategoryID      string `json:"category"`
	TransactionType string `json:"transaction_type"`
}
