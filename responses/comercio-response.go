package responses

type ComercioResponse struct {
	Merchant_id   int    `json:"merchant_id"`
	Merchant_name string `json:"merchant_name"`
	Commission    int    `json:"commission"`
	Created_at    int    `json:"created_at"`
	Updated_at    int    `json:"updated_at"`
}
