-- +goose Up
-- +goose StatementBegin
create table if not exists containers
(
    id      text not null,
    image   text not null,
    state   text not null,
    status  text not null,
    name    text not null,
    primary key (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists containers;
-- +goose StatementEnd
