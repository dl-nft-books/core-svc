-- +migrate Up

create table promocodes
(
    id                 bigserial primary key,
    promocode          text      not null,
    discount           float      not null default 0,
    initial_usages     bigint    not null default 0,
    left_usages        bigint    not null default 0,
    expiration_date    timestamp not null default CURRENT_TIMESTAMP,
    state              int8      not null default 1
);

-- +migrate Down

drop table if exists promocodes;
