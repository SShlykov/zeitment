-- +goose Up
-- +goose StatementBegin
ALTER TABLE pages ADD COLUMN number INT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE pages DROP COLUMN number;
-- +goose StatementEnd
