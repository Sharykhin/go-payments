-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE user_passwords
(
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    password character varying(80) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE user_passwords
    ADD CONSTRAINT fk_user_passwords_user_id_users_id FOREIGN KEY (user_id)
    REFERENCES users (id) MATCH SIMPLE
    ON UPDATE RESTRICT
    ON DELETE RESTRICT;

CREATE INDEX fki_fk_user_passwords_user_id_users_id
    ON user_passwords(user_id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE user_passwords;
