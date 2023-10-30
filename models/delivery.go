package models

type Delivery struct {
	name    string `json:"name"`
	phone   string `json:"phone"`
	zip     string `json:"zip"`
	city    string `json:"city"`
	address string `json:"address"`
	region  string `json:"region"`
	email   string `json:"email"`
}
