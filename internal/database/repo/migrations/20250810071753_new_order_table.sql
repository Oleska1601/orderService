-- +goose Up
create table if not exists orders (
    order_uid varchar(50) primary key, /*Transaction ID for assembly orders grouping. Orders in the same buyer's cart will have the same orderUid*/
    track_number varchar(50) not null unique, /*трек-номер для отслеживания посылки - принимаем, что заказ отп*/
    entry varchar(10) not null,
    locale varchar(10),
    internal_signature varchar(100),
    customer_id varchar(50) not null,
    delivery_service varchar(50),
    shardkey integer not null,
    sm_id integer not null,
    date_created timestamp with time zone not null,
    oof_shard varchar(10) not null
);


-- +goose Down
drop table if exists orders;

