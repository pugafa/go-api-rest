package model

type Transaction struct {
	Transaction_id uint    `json:"Transaction_id" gorm:"primary_key"`
	Merchant_id    int     `json:"merchant_id"`
	Amount         float64 `json:"amount,"`
	Commission     string  `json:"commission,"`
	Fee            string  `json:"fee,"`
	Created_at     int     `json:"created_at,"`
}
