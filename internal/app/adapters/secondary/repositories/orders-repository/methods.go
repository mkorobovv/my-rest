package orders_repository

import (
	"context"
	"time"

	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

func (repo *OrdersRepository) Create(ctx context.Context, _order order.Order) (uid string, err error) {
	query, args, err := getQueryCreate(_order)
	if err != nil {
		return "", err
	}

	err = repo.DB.GetContext(ctx, &uid, query, args...)
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (repo *OrdersRepository) Get(ctx context.Context, trackNumber string) (_order order.Order, err error) {
	query, args, err := getQueryGet(trackNumber)
	if err != nil {
		return order.Order{}, err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var dto getOrderDTO

	err = repo.DB.GetContext(ctx, &dto, query, args...)
	if err != nil {
		return order.Order{}, err
	}

	_order, err = dto.toEntity()
	if err != nil {
		return order.Order{}, err
	}

	return _order, nil
}

func (repo *OrdersRepository) Update(ctx context.Context, _order order.Order) (err error) {
	query, args, err := getQueryUpdate(_order)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	_, err = repo.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
