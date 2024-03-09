-- +goose Up
-- +goose StatementBegin
ALTER TABLE paragraphs ADD COLUMN title VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE paragraphs DROP COLUMN title;
-- +goose StatementEnd
