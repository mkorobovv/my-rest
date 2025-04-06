# REST API EXAMPLE IN GOLANG

orders - core service for working with orders

Contents:
-
**[PART 1: API](#1)**    
&emsp;**[1.1 Update order](#1.1)**    
&emsp;**[1.2 Get order](#1.2)**    
**[PART 2: Background processes](#2)**    
&emsp;**[2.1 Publishing orders to Kafka](#2.1)**      

---

## <a name="1">PART 1: API:</a>

## <a name="1.1">1.1 Update order:</a>

&emsp;`http://localhost:8080/api/v1/orders/{trackNumber}`  
**METHOD:** PUT

**Request body**
```json5
{
  "uid": "ffb39d3c-7c6c-4b67-99c5-e7dcd18ed8c7",
  "track_number": "9VZ51UKOG7PUHGFQ18WQ",
  "locale": "Europe/Germany",
  "customer_id": 3400783320913439000,
  "created_dt": "2025-04-06 16:00:48",
  "payment": {
    "transaction_id": "TXEJGP2RLC43WDOTXIRP",
    "currency": "USD",
    "amount": 469439,
    "provider": "wbpay",
    "payment_dt": "2024-12-18 08:36:24",
    "delivery_cost": 5471,
    "goods_total": 474910,
    "bank": "WBBank"
  },
  "delivery": {
    "recipient_name": "Jane Brown",
    "phone_number": "+77232122870",
    "zip_code": "777720",
    "address": "990 Oak St, San Antonio, PA 777720",
    "email": "dvis@aol.com"
  }
}
```

**Response body**

Case OK:

`HTTP/1.1 200 OK`

Case error:

```json5
{
  "kind": "kind of error",
  "detail": "error details"
}
```

## <a name="1.2">1.2 Get order:</a>

&emsp;`http://localhost:8080/api/v1/orders/{trackNumber}`  
**METHOD:** GET

**Response body**

Case OK:

```json5
{
  "uid": "ffb39d3c-7c6c-4b67-99c5-e7dcd18ed8c7",
  "track_number": "9VZ51UKOG7PUHGFQ18WQ",
  "locale": "Europe/Germany",
  "customer_id": 3400783320913439000,
  "created_dt": "2025-04-06 16:00:48",
  "payment": {
    "transaction_id": "TXEJGP2RLC43WDOTXIRP",
    "currency": "USD",
    "amount": 469439,
    "provider": "wbpay",
    "payment_dt": "2024-12-18 08:36:24",
    "delivery_cost": 5471,
    "goods_total": 474910,
    "bank": "WBBank"
  },
  "delivery": {
    "recipient_name": "Jane Brown",
    "phone_number": "+77232122870",
    "zip_code": "777720",
    "address": "990 Oak St, San Antonio, PA 777720",
    "email": "dvis@aol.com"
  },
  "items": [
    {
      "chrt_id": 183614,
      "price": 77224,
      "name": "Philips Speaker Series with 2-Year Warranty",
      "sale": 0,
      "total_price": 77224,
      "nm_id": 396091
    },
    {
      "chrt_id": 803070,
      "price": 61487,
      "name": "Apple Watch Series with Bluetooth",
      "sale": 37,
      "total_price": 38736.81,
      "nm_id": 973896
    }
  ]
}
```

Case error:

```json5
{
  "kind": "kind of error",
  "detail": "error details"
}
```

## <a name="2">PART 2: Background processes:</a>

## <a name="2.1">2.1 Publishing orders to Kafka:</a>

```json5
{
  "uid": "ffb39d3c-7c6c-4b67-99c5-e7dcd18ed8c7",
  "track_number": "9VZ51UKOG7PUHGFQ18WQ",
  "locale": "Europe/Germany",
  "customer_id": 3400783320913439000,
  "created_dt": "2025-04-06 16:00:48",
  "payment": {
    "transaction_id": "TXEJGP2RLC43WDOTXIRP",
    "currency": "USD",
    "amount": 469439,
    "provider": "wbpay",
    "payment_dt": "2024-12-18 08:36:24",
    "delivery_cost": 5471,
    "goods_total": 474910,
    "bank": "WBBank"
  },
  "delivery": {
    "recipient_name": "Jane Brown",
    "phone_number": "+77232122870",
    "zip_code": "777720",
    "address": "990 Oak St, San Antonio, PA 777720",
    "email": "dvis@aol.com"
  },
  "items": [
    {
      "chrt_id": 183614,
      "price": 77224,
      "name": "Philips Speaker Series with 2-Year Warranty",
      "sale": 0,
      "total_price": 77224,
      "nm_id": 396091
    },
    {
      "chrt_id": 803070,
      "price": 61487,
      "name": "Apple Watch Series with Bluetooth",
      "sale": 37,
      "total_price": 38736.81,
      "nm_id": 973896
    }
  ]
}
```