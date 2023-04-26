-- +migrate Up

create table promocodes
(
    id                 bigserial primary key,
    promocode          text      not null,
    discount           float      not null default 0,
    initial_usages     bigint    not null default 0,
    usages        bigint    not null default 0,
    expiration_date    timestamp not null default CURRENT_TIMESTAMP,
    state              int8      not null default 1
);

create table promocodes_books
(
    promocode_id  bigserial references promocodes(id) ON DELETE CASCADE ,
    book_id       bigserial,
    primary key (promocode_id, book_id)
);
-- +migrate Down

drop table if exists promocodes_books;
drop table if exists promocodes;