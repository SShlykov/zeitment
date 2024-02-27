-- +goose Up
-- +goose StatementBegin
alter table books alter column variables set default '{}';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE books ALTER COLUMN variables DROP DEFAULT;
-- +goose StatementEnd
