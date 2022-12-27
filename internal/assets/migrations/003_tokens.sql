-- +migrate Up

create table tokens
(
    id                 bigserial primary key,
    account            text      not null default '',
    token_id           bigint    not null default 0,
    book_id            bigint    not null default 0,
    payment_id         bigint    not null default 0,
    metadata_hash      text      not null default 0 unique,
    signature          text      not null default '',
    status             int8      not null default 0
);

-- +migrate Down

drop table if exists tokens;
