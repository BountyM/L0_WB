package models

type Order struct {
	Order_uid          string   `json:"order_uid" db:"order_uid"`
	Track_number       string   `json:"track_number" db:"track_number"`
	Entry              string   `json:"entry" db:"entry"`
	Delivery           Delivery `json:"delivery"`
	Payment            Payment  `json:"payment"`
	Items              []Item   `json:"items"`
	Locale             string   `json:"locale" db:"chrt_id"`
	Internal_signature string   `json:"internal_signature" db:"chrt_id"`
	Customer_id        string   `json:"customer_id" db:"chrt_id"`
	Delivery_service   string   `json:"delivery_service" db:"chrt_id"`
	Shardkey           string   `json:"shardkey" db:"chrt_id"`
	Sm_id              int      `json:"sm_id" db:"chrt_id"`
	Date_created       string   `json:"date_created" db:"chrt_id"`
	Oof_shard          string   `json:"oof_shard" db:"chrt_id"`
}
