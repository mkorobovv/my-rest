package orders_repository

import (
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/mkorobovv/my-rest/internal/app/domain/order"
)

func getQueryCreate(_order order.Order) (string, []interface{}, error) {
	itemsJSON, err := json.Marshal(_order.Items)
	if err != nil {
		return "", nil, err
	}

	query := `
    WITH insert_order_cte AS (
        INSERT INTO orders.orders (
            track_number,
            locale,
            customer_id,
            transaction_id,
            currency,
            amount,
            provider,
            bank,
            delivery_cost,
            goods_total,
            recipient_name,
            phone_number,
            zip_code,
            address,
			payment_dt,
            email
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
        ) RETURNING uid
    ),
    items_cte AS (
        SELECT 
            io.uid AS order_uid,
            i.chrt_id,
            i.price,
            i.name,
            i.sale,
            i.total_price,
            i.nm_id
        FROM insert_order_cte io
        CROSS JOIN JSONB_TO_RECORDSET($17) AS i (
            chrt_id bigint,
            price numeric(15,2),
            name text,
            sale integer,
            total_price numeric(15,2),
            nm_id bigint
        )
    ),
	insert_items_cte AS (
        INSERT INTO orders.items (
            order_uid,
            chrt_id,
            price,
            name,
            sale,
            total_price,
            nm_id
        ) SELECT
              c.order_uid,
              c.chrt_id,
              c.price,
              c.name,
              c.sale,
              c.total_price,
              c.nm_id
        FROM items_cte c
    )
    SELECT c.uid FROM insert_order_cte c;`

	args := []interface{}{
		_order.TrackNumber,
		_order.Locale,
		_order.CustomerID,
		_order.Payment.TransactionID,
		_order.Payment.Currency,
		_order.Payment.Amount,
		_order.Payment.Provider,
		_order.Payment.Bank,
		_order.Payment.DeliveryCost,
		_order.Payment.GoodsTotal,
		_order.Delivery.RecipientName,
		_order.Delivery.PhoneNumber,
		_order.Delivery.ZipCode,
		_order.Delivery.Address,
		_order.Payment.PaymentDt,
	}

	if _order.Delivery.Email != nil {
		args = append(args, *_order.Delivery.Email)
	} else {
		args = append(args, nil)
	}

	args = append(args, itemsJSON)

	return query, args, nil
}

func getQueryGet(trackNumber string) (query string, args []interface{}, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sb := psql.
		Select(
			"o.uid",
			"o.track_number",
			"o.locale",
			"o.customer_id",
			"o.created_dt",
			"o.transaction_id",
			"o.currency",
			"o.amount",
			"o.provider",
			"o.bank",
			"o.is_deleted",
			"o.payment_dt",
			"o.delivery_cost",
			"o.goods_total",
			"o.recipient_name",
			"o.phone_number",
			"o.zip_code",
			"o.address",
			"o.email",
		).
		From("orders.orders o").
		Where(sq.Eq{"o.track_number": trackNumber}).
		Where(sq.Eq{"o.is_deleted": false})

	sbItems := psql.
		Select("json_agg(to_json(i.*)) AS items").
		From("orders.items i").
		Where("o.uid = i.order_uid")

	queryItems, _, err := sbItems.ToSql()
	if err != nil {
		return "", nil, err
	}

	queryItems = fmt.Sprintf("LATERAL (%s) i ON true", queryItems)

	sb = sb.Column("i.items").LeftJoin(queryItems)

	query, args, err = sb.ToSql()
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

func getQueryUpdate(_order order.Order) (query string, args []interface{}, err error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	ub := psql.
		Update("orders.orders o").
		Set("locale", _order.Locale).
		Set("customer_id", _order.CustomerID).
		Set("transaction_id", _order.Payment.TransactionID).
		Set("currency", _order.Payment.Currency).
		Set("amount", _order.Payment.Amount).
		Set("provider", _order.Payment.Provider).
		Set("bank", _order.Payment.Bank).
		Set("is_deleted", _order.IsDeleted).
		Set("delivery_cost", _order.Payment.DeliveryCost).
		Set("goods_total", _order.Payment.GoodsTotal).
		Set("recipient_name", _order.Delivery.RecipientName).
		Set("phone_number", _order.Delivery.PhoneNumber).
		Set("zip_code", _order.Delivery.ZipCode).
		Set("address", _order.Delivery.Address)

	if _order.Delivery.Email != nil {
		ub = ub.
			Set("email", *_order.Delivery.Email)
	}

	ub = ub.Where(sq.Eq{"o.track_number": _order.TrackNumber})

	query, args, err = ub.ToSql()
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}
