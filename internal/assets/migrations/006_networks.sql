-- +migrate Up

alter table tokens add column chain_id bigint not null default 0;

-- +migrate Down

alter table tokens drop column chain_id;
