-- +goose Up
-- +goose StatementBegin
ALTER table pages ADD COLUMN title text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE pages DROP COLUMN title;
-- +goose StatementEnd
