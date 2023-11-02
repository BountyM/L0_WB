DROP TABLE deliveries, payments, items, orders;

CREATE TABLE IF NOT EXISTS public.orders
(
    "order_uid" TEXT PRIMARY KEY NOT NULL,
    "track_number" TEXT,
    "entry" TEXT,
    "locale" TEXT,
    "internal_signature" TEXT,
    "customer_id" TEXT,
    "delivery_service" TEXT,
    "shardkey" TEXT,
    "sm_id" INT,
    "date_created" TEXT,
    "oof_shard" TEXT
);

CREATE TABLE IF NOT EXISTS public.deliveries
(
    "order_uid" TEXT NOT NULL,
    "name" TEXT PRIMARY KEY NOT NULL,
    "phone" TEXT,
    "zip" TEXT,
    "city" TEXT,
    "address" TEXT,
    "region" TEXT,
    "email" TEXT
);

CREATE TABLE IF NOT EXISTS public.payments
(
    "transaction" TEXT NOT NULL,
    "request_id" TEXT PRIMARY KEY NOT NULL,
    "currency" TEXT,
    "provider" TEXT,
    "amount" INT,
    "payment_dt" INT,
    "bank" TEXT,
    "delivery_cost" INT,
    "goods_total" INT,
    "custom_fee" INT
);
CREATE TABLE IF NOT EXISTS public.items
(
    "chrt_id" INT PRIMARY KEY NOT NULL,
    "track_number" TEXT NOT NULL,
    "price" INT,
    "rid" TEXT,
    "name" TEXT,
    "sale" INT,
    "size" TEXT,
    "total_price" INT,
    "nm_id" INT,
    "brand" TEXT,
    "status" INT
);

