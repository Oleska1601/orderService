-- +goose Up
/*name, phone, city считаю не равным customer_name, customer_phone, customer_email, тк могут быть разные люди: получатель доставки и тот, кто ее оформил*/
create table if not exists deliveries (
    order_uid varchar(50) primary key references orders(order_uid) on delete cascade, 
    name varchar(100) not null,
    phone varchar(15) not null,
    zip varchar(20) not null, /*почтовый индекс*/
    city varchar(100) not null,
    address text not null,
    region varchar(100) not null,
    email varchar(255) not null
);


-- +goose Down
drop table if exists deliveries;

