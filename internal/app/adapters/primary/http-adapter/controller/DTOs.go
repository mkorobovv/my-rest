package controller

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

type responseError struct {
	Kind   string `json:"kind"`
	Detail string `json:"detail"`

	status int
}

func (err responseError) Error() string {
	return err.Detail
}

type OrderDTO struct {
	UID         *string      `json:"uid" validate:"required"`
	TrackNumber *string      `json:"track_number" validate:"required"`
	Locale      *string      `json:"locale" validate:"required"`
	CustomerID  *int64       `json:"customer_id" validate:"required"`
	CreatedDt   *time.Time   `json:"created_dt" validate:"required"`
	Payment     *PaymentDTO  `json:"payment" validate:"required"`
	Delivery    *DeliveryDTO `json:"delivery" validate:"required"`
	Items       []ItemDTO    `json:"items" validate:"dive"`
}

func (dto OrderDTO) validate() error {
	validate := validator.New()

	err := validate.Struct(dto)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			respErr := responseError{
				Kind:   "validation",
				status: http.StatusBadRequest,
				Detail: "failed struct validation: ",
			}

			for _, validationError := range validationErrors {
				respErr.Detail = respErr.Detail + fmt.Sprintf("field: %s, reason: %s; ", validationError.Field(), validationError.Tag())
			}

			return respErr
		}

		return err
	}

	return nil
}

func (dto OrderDTO) toRequest() order.Order {
	return order.Order{
		UID:         *dto.UID,
		TrackNumber: *dto.TrackNumber,
		Locale:      *dto.Locale,
		CustomerID:  *dto.CustomerID,
		CreatedDt:   *dto.CreatedDt,
		Payment:     dto.Payment.toRequest(),
		Delivery:    dto.Delivery.toRequest(),
		Items:       toItemRequests(dto.Items),
	}
}

type ItemDTO struct {
	ChrtID     *int64   `json:"chrt_id" validate:"required"`
	Price      *float64 `json:"price" validate:"gt=0"`
	Name       *string  `json:"name" validate:"required"`
	Sale       *int64   `json:"sale"`
	TotalPrice *float64 `json:"total_price" validate:"gt=0"`
	NmID       *int64   `json:"nm_id"`
}

func (dto ItemDTO) toRequest() order.Item {
	return order.Item{
		ChrtID:     *dto.ChrtID,
		Price:      *dto.Price,
		Name:       *dto.Name,
		Sale:       dto.Sale,
		TotalPrice: *dto.TotalPrice,
		NmID:       dto.NmID,
	}
}

func toItemRequests(dtos []ItemDTO) []order.Item {
	items := make([]order.Item, 0, len(dtos))

	for _, dto := range dtos {
		item := dto.toRequest()

		items = append(items, item)
	}

	return items
}

type PaymentDTO struct {
	TransactionID *string    `json:"transaction_id" validate:"required"`
	Currency      *string    `json:"currency" validate:"required"`
	Amount        *float64   `json:"amount" validate:"gt=0"`
	Provider      *string    `json:"provider" validate:"required"`
	PaymentDt     *time.Time `json:"payment_dt" validate:"required"`
	DeliveryCost  *float64   `json:"delivery_cost" validate:"gte=0"`
	GoodsTotal    *float64   `json:"goods_total" validate:"gt=0"`
	Bank          *string    `json:"bank" validate:"required"`
}

func (dto PaymentDTO) toRequest() order.Payment {
	return order.Payment{
		TransactionID: *dto.TransactionID,
		Currency:      *dto.Currency,
		Amount:        *dto.Amount,
		Provider:      *dto.Provider,
		PaymentDt:     *dto.PaymentDt,
		DeliveryCost:  *dto.DeliveryCost,
		GoodsTotal:    *dto.GoodsTotal,
		Bank:          *dto.Bank,
	}
}

type DeliveryDTO struct {
	RecipientName *string `json:"recipient_name" validate:"required"`
	PhoneNumber   *string `json:"phone_number" validate:"required,len=11"`
	ZipCode       *string `json:"zip_code" validate:"required"`
	Address       *string `json:"address" validate:"required"`
	Email         *string `json:"email" validate:"email"`
}

func (dto DeliveryDTO) toRequest() order.Delivery {
	return order.Delivery{
		RecipientName: *dto.RecipientName,
		PhoneNumber:   *dto.PhoneNumber,
		ZipCode:       *dto.ZipCode,
		Address:       *dto.Address,
		Email:         dto.Email,
	}
}
