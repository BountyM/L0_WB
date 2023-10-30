package models

type Payment struct {
	transaction   string
	request_id    string
	currency      string
	provider      string
	amount        int
	payment_dt    int
	bank          string
	delivery_cost int
	goods_total   int
	custom_fee    int
}
