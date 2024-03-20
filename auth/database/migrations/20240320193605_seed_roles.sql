-- +goose Up
-- +goose StatementBegin
BEGIN TRANSACTION;
INSERT INTO roles (name) VALUES ('admin'), ('user'), ('guest');
COMMIT TRANSACTION;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
begin transaction;
delete from roles where name in ('admin', 'user', 'guest');
commit transaction;
-- +goose StatementEnd
