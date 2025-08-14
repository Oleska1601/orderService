-- +goose Up
create table if not exists items (
    item_id serial primary key,
    chrt_id integer not null, /*product size ID in the WB systems*/
    price integer not null,
    name text not null,
    nm_id integer not null, /*WB article*/
    brand text not null
);


-- +goose Down
drop table if exists items;

