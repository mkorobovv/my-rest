package order

import "time"

type Order struct {
	UID         string    `json:"uid"`
	TrackNumber string    `json:"track_number"`
	Locale      string    `json:"locale"`
	CustomerID  int64     `json:"customer_id"`
	CreatedDt   time.Time `json:"created_dt"`
	Payment     Payment   `json:"payment"`
	Delivery    Delivery  `json:"delivery"`
	Items       []Item    `json:"items"`
}

type Item struct {
	ChrtID     int64   `json:"chrt_id"`
	Price      float64 `json:"price"`
	Name       string  `json:"name"`
	Sale       int64   `json:"sale"`
	TotalPrice float64 `json:"total_price"`
	NmID       int64   `json:"nm_id"`
}

type Payment struct {
	TransactionID string  `json:"transaction_id"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	Provider      string  `json:"provider"`
	PaymentDt     string  `json:"payment_dt"`
	DeliveryCost  float64 `json:"delivery_cost"`
	GoodsTotal    float64 `json:"goods_total"`
	Bank          string  `json:"bank"`
}

type Delivery struct {
	RecipientName string `json:"recipient_name"`
	PhoneNumber   string `json:"phone_number"`
	ZipCode       string `json:"zip_code"`
	Address       string `json:"address"`
	Email         string `json:"email"`
}
