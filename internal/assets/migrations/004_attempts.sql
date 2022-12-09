-- +migrate Up

alter table tasks add column created_at timestamp not null default CURRENT_TIMESTAMP;

-- +migrate Down

alter table tasks drop column created_at;
