
CREATE DATABASE information_service;

Create table orders (
    order_uid varchar(50),
    track_number varchar(50),
    "entry" varchar(50),
    locale varchar(50),
    internal_signature varchar(50),
    customer_id varchar(50),
    delivery_service varchar(50),
    shardkey varchar(50),
    sm_id int,
    date_created varchar(50),
    oof_shard varchar(50)
);

Create table delivery(
    order_uid varchar(50),
    name varchar(50),
    phone varchar(50),
    zip varchar(50),
    city varchar(50),
    "address" varchar(50),
    region varchar(50),
    email varchar(50)
    CONSTRAINT fk_delivery_order_uid FOREIGN KEY (order_uid) REFERENCES Orders(order_uid)
);

Create table payment (
    transaction varchar(50),
    request_id varchar(50),
    currency varchar(50),
    "provider" varchar(50),
    amount int,
    payment_dt int,
    bank varchar(50),
    delivery_cost int,
    goods_total int,
    custom_fee int
    CONSTRAINT fk_payment_transaction FOREIGN KEY (transaction) REFERENCES Orders(order_uid)
);

Create table orderitems (
    order_uid varchar(50),
    chrt_id int
);

Create table items (
    chrt_id int,
    track_number varchar(50),
    price int,
    rid varchar(50),
    "name" varchar(50),
    sale int,
    "size" varchar(50),
    total_price int,
    nm_id int,
    brand varchar(50),
    "status" int
    CONSTRAINT fk_items_track_number FOREIGN KEY (track_number) REFERENCES Orders(track_number)
);