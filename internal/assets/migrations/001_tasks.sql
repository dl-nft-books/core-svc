-- +migrate Up

create table tasks
(
    id        bigserial primary key,
    token_id  bigint    not null default 0,
    book_id   bigint    not null default 0,
    account   text      not null default '',
    signature text      not null default '',
    ipfs_hash text      not null default '',
    status    int8      not null default 0
);

-- +migrate Down

drop table if exists tasks;