-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE transactions(
  id serial not null,
  user_id integer not null,
  amount decimal(8,2) not null,
  description text not null,
  transaction_id varchar(60) not null,
  status varchar(80) not null,
  created_at timestamp(0) without time zone,
  CONSTRAINT transactions_pkey PRIMARY KEY (id)
);

ALTER TABLE transactions
    ADD CONSTRAINT fk_transactions_user_id_users_id FOREIGN KEY (user_id)
    REFERENCES public.users (id) MATCH SIMPLE
    ON UPDATE RESTRICT
    ON DELETE RESTRICT;
CREATE INDEX fki_fk_transactions_user_id_users_id
    ON transactions(user_id);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE transactions;
