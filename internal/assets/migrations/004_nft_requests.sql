-- +migrate Up

create table nft_requests
(
    id                 bigserial primary key,
    payer_address          text      not null,
    collection_address     float      not null default 0,
    nft_id     bigint      not null default 0,
    book_id     bigint      not null default 0,
    status             int8      not null default 0,
    created_at    timestamp not null default CURRENT_TIMESTAMP,
    last_updated_at    timestamp not null default CURRENT_TIMESTAMP
);

-- +migrate Down

drop table if exists nft_requests;