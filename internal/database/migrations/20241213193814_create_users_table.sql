-- +goose Up
-- +goose StatementBegin
create table users (
    id char(36) primary key,
    name varchar(255) not null,
    email varchar(50) unique not null,
    phone varchar(20) unique not null,
    password varchar(50) not null,
    role varchar(50) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
