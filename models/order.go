package models

type Order struct {
	order_uid          string
	track_number       string
	entry              string
	delivery           Delivery
	payment            Payment
	items              []Item
	locale             string
	internal_signature string
	customer_id        string
	delivery_service   string
	shardkey           string
	sm_id              int
	date_created       string
	oof_shard          string
}
