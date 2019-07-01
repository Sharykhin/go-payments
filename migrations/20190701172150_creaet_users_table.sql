-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE users(
  id serial not null,
  first_name varchar(80) not null,
  last_name varchar(80),
  email varchar(80) not null,
  created_at timestamp(0) without time zone,
  deleted_at timestamp(0) without time zone,
  CONSTRAINT users_pkey PRIMARY KEY (id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;
