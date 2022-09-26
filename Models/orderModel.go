package models

type Order struct {
	Type        string  `json:"type"`
	Quantity    int     `json:"quantity"`
	MarketPrice float64 `json:"marketPrice"`
	LimitPrice  float64 `json:"limitPrice"`
}

type Queue struct {
	Order []Order
	Size  int
}
