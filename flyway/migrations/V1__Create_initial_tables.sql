CREATE SCHEMA IF NOT EXISTS orders;

CREATE TABLE IF NOT EXISTS orders.orders (
    uid TEXT PRIMARY KEY,
    track_number VARCHAR(20) NOT NULL,
    locale TEXT NOT NULL,
    customer_id BIGINT NOT NULL,
    created_dt TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    transaction_id TEXT NOT NULL,
    currency VARCHAR(3) NOT NULL,
    AMOUNT NUMERIC(15,2) NOT NULL,
    provider TEXT NOT NULL,
    bank TEXT NOT NULL,
    payment_dt TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    delivery_cost NUMERIC(15,2) NOT NULL,
    goods_total NUMERIC(15,2) NOT NULL,
    recipient_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    zip_code TEXT NOT NULL,
    address TEXT NOT NULL,
    email TEXT
);

CREATE TABLE IF NOT EXISTS orders.items (
    id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_uid TEXT,
    chrt_id BIGINT NOT NULL,
    price NUMERIC(15,2) NOT NULL,
    name TEXT NOT NULL,
    sale INTEGER,
    total_price NUMERIC(15,2),
    nm_id BIGINT,
    CONSTRAINT fk_orders_items
        FOREIGN KEY (order_uid)
        REFERENCES orders.orders(uid)
);