-- +goose Up
create table if not exists payments (
    order_uid varchar(50) primary key references orders(order_uid) on delete cascade, 
    transaction varchar(50) not null unique,
    request_id varchar(50),
    currency varchar(10) not null,
    provider varchar(50) not null,
    amount integer not null,
    payment_dt integer not null,
    bank varchar(50) not null,
    delivery_cost integer,
    goods_total integer not null, /*стоимость заказа без доставки*/
    custom_fee integer
);


-- +goose Down
drop table if exists payments;
