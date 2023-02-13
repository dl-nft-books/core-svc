-- +migrate Up

alter table tokens add column is_token_payment boolean default true;

-- +migrate Down

alter table tokens drop column is_token_payment;
