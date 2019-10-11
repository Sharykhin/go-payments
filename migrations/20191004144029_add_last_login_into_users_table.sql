-- +goose Up
-- SQL in this section is executed when the migration is applied.
alter table users add column last_login timestamp without time zone;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
alter table users drop column last_login;