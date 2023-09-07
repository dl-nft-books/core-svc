-- +migrate Up

create table nft_requests
(
    id bigserial primary key,
    requester char(42) not null default '',
    marketplace_request_id bigint not null default 0,
    nft_address char(42) not null default '',
    nft_id bigint not null default 0,
    book_id bigint not null default 0,
    chain_id bigint not null default 0,
    status int8 not null default 0,
    created_at timestamp not null default CURRENT_TIMESTAMP,
    last_updated_at timestamp not null default CURRENT_TIMESTAMP
);

-- +migrate Down

drop table if exists nft_requests;