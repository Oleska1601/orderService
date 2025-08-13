-- +goose Up
create table if not exists items (
    item_id serial primary key,
    chrt_id integer not null, /*product size ID in the WB systems*/
    price integer not null,
    name text,
    nm_id integer not null, /*WB article*/
    brand text
);


-- +goose Down
drop table if exists items;

