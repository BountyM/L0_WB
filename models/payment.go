package models

type Payment struct {
	transaction   string `json:"transaction"`
	request_id    string `json:"request_id"`
	currency      string `json:"currency"`
	provider      string `json:"provider"`
	amount        int    `json:"amount"`
	payment_dt    int    `json:"payment_dt"`
	bank          string `json:"bank"`
	delivery_cost int    `json:"delivery_cost"`
	goods_total   int    `json:"goods_total"`
	custom_fee    int    `json:"custom_fee"`
}
