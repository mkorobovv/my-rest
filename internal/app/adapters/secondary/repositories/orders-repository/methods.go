package orders_repository

import (
	"context"

	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

func (repo *OrdersRepository) Create(ctx context.Context, order order.Order) (err error) {
	return nil
}

func (repo *OrdersRepository) Get(ctx context.Context, trackNumber string) (order order.Order, err error) {
	//TODO implement me
	panic("implement me")
}

func (repo *OrdersRepository) Update(ctx context.Context, order order.Order) (err error) {
	return nil
}
