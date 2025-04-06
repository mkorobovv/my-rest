package api_service

import (
	"context"
	"errors"
	"time"

	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

func (svc *ApiService) Create(ctx context.Context, _order order.Order) (uid string, err error) {
	uid, err = svc.ordersRepository.Create(ctx, _order)
	if err != nil {
		return "", err
	}

	err = svc.cache.Storage.Add(_order.TrackNumber, _order, 2*time.Hour)
	if err != nil {
		svc.logger.Error(err.Error(), "cache err")
	}

	return uid, nil
}

func (svc *ApiService) Get(ctx context.Context, trackNumber string) (order.Order, error) {
	getOrderFn := func() (any, error) {
		return svc.ordersRepository.Get(ctx, trackNumber)
	}

	result, err, _ := svc.cache.Memoize(trackNumber, getOrderFn)
	if err != nil {
		return order.Order{}, err
	}

	_order, ok := result.(order.Order)
	if !ok {
		return order.Order{}, errors.New("invalid result type")
	}

	return _order, nil
}

func (svc *ApiService) Update(ctx context.Context, _order order.Order) (err error) {
	err = svc.ordersRepository.Update(ctx, _order)
	if err != nil {
		return err
	}

	return nil
}
