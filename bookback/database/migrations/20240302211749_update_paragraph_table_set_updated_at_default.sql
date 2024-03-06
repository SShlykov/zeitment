-- +goose Up
-- +goose StatementBegin
ALTER TABLE paragraphs ALTER COLUMN updated_at SET DEFAULT now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE paragraphs ALTER COLUMN updated_at DROP DEFAULT;
-- +goose StatementEnd
