package responses

type TransactionResponse struct {
	Transaction_id int    `json:"transaction_id"`
	Merchant_id    string `json:"merchant_id"`
	Amount         int    `json:"amount,"`
	Commission     int    `json:"commission,"`
	Fee            int    `json:"fee,"`
	Created_at     int    `json:"created_at,"`
}
