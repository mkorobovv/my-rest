package orders_repository

import (
	"encoding/json"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
	"time"
)

type getOrderDTO struct {
	UID           string  `db:"uid"`
	TrackNumber   string  `db:"track_number"`
	Locale        string  `db:"locale"`
	CustomerID    int64   `db:"customer_id"`
	CreatedDt     string  `db:"created_dt"`
	TransactionID string  `db:"transaction_id"`
	Currency      string  `db:"currency"`
	Amount        float64 `db:"amount"`
	Provider      string  `db:"provider"`
	PaymentDt     string  `db:"payment_dt"`
	IsDeleted     bool    `db:"is_deleted"`
	DeliveryCost  float64 `db:"delivery_cost"`
	GoodsTotal    float64 `db:"goods_total"`
	Bank          string  `db:"bank"`
	RecipientName string  `db:"recipient_name"`
	PhoneNumber   string  `db:"phone_number"`
	ZipCode       string  `db:"zip_code"`
	Address       string  `db:"address"`
	Email         *string `db:"email"`
	Items         []byte  `db:"items"`
}

func (dto getOrderDTO) toEntity() (order.Order, error) {
	createdDt, err := time.Parse(time.RFC3339, dto.CreatedDt)
	if err != nil {
		return order.Order{}, err
	}

	paymentDt, err := time.Parse(time.RFC3339, dto.PaymentDt)
	if err != nil {
		return order.Order{}, err
	}

	items := make([]order.Item, 0)

	if len(dto.Items) != 0 {
		err = json.Unmarshal(dto.Items, &items)
		if err != nil {
			return order.Order{}, err
		}
	}

	return order.Order{
		UID:         dto.UID,
		TrackNumber: dto.TrackNumber,
		Locale:      dto.Locale,
		CustomerID:  dto.CustomerID,
		CreatedDt:   createdDt,
		IsDeleted:   dto.IsDeleted,
		Payment: order.Payment{
			TransactionID: dto.TransactionID,
			Currency:      dto.Currency,
			Amount:        dto.Amount,
			Provider:      dto.Provider,
			PaymentDt:     paymentDt,
			DeliveryCost:  dto.DeliveryCost,
			GoodsTotal:    dto.GoodsTotal,
			Bank:          dto.Bank,
		},
		Delivery: order.Delivery{
			RecipientName: dto.RecipientName,
			PhoneNumber:   dto.PhoneNumber,
			ZipCode:       dto.ZipCode,
			Address:       dto.Address,
			Email:         dto.Email,
		},
		Items: items,
	}, nil
}
