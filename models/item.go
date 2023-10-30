package models

type Item struct {
	chrt_id      int    `json:"chrt_id"`
	track_number string `json:"track_number"`
	price        int    `json:"price"`
	rid          string `json:"rid"`
	name         string `json:"name"`
	sale         int    `json:"sale"`
	size         string `json:"size"`
	total_price  int    `json:"total_price"`
	nm_id        int    `json:"nm_id"`
	brand        string `json:"brand"`
	status       int    `json:"status"`
}
