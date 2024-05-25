package model

type Comercio struct {
	Merchant_id   uint   `json:"id" gorm:"primary_key"`
	Merchant_name string `json:"merchant_name"`
	Commission    string `json:"commission,"`
	Created_at    int    `json:"created_at,"`
	Updated_at    int    `json:"updated_at,"`
}
