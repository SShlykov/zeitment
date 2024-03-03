-- +goose Up
-- +goose StatementBegin
ALTER TABLE map_variables ADD COLUMN updated_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE map_variables DROP COLUMN updated_at;
-- +goose StatementEnd
