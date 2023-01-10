-- +migrate Up

alter table tokens add unique(metadata_hash);

-- +migrate Down

