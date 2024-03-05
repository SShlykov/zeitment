-- +goose Up
-- +goose StatementBegin
BEGIN;

ALTER TABLE book_events ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE map_variables ALTER COLUMN id SET DEFAULT uuid_generate_v4();

ALTER TABLE book_events ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE book_events ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE map_variables ALTER COLUMN created_at SET DEFAULT now();

CREATE TRIGGER update_book_events_updated_at BEFORE UPDATE ON book_events FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
