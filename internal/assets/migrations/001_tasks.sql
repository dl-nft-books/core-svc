-- +migrate Up

create table tasks
(
    id                 bigserial primary key,
    token_id           bigint    not null default 0,
    book_id            bigint    not null default 0,
    token_name         text      not null default '',
    account            text      not null default '',
    banner_ipfs_hash   text      not null default '',
    metadata_ipfs_hash text      not null default '',
    book_uri           text      not null default '',
    banner             bytea     NOT NULL,
    status             int8      not null default 0,
    created_at         timestamp not null default CURRENT_TIMESTAMP,
    chain_id            bigint    not null default 0
);

-- +migrate Down

drop table if exists tasks;
