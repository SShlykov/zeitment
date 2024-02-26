-- +goose Up
-- +goose StatementBegin
-- Должен быть генератор uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

BEGIN;

ALTER TABLE books ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE chapters ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE map_variables ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE pages ALTER COLUMN id SET DEFAULT uuid_generate_v4();
ALTER TABLE paragraphs ALTER COLUMN id SET DEFAULT uuid_generate_v4();

-- Обновляем дефолтное значение для inserted_at для каждой таблицы
ALTER TABLE books ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE chapters ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE map_variables ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE pages ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE paragraphs ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;

-- Создаем общую функцию для обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Создаем триггеры для каждой таблицы, чтобы обновлять updated_at
CREATE TRIGGER update_books_updated_at BEFORE UPDATE ON books FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_chapters_updated_at BEFORE UPDATE ON chapters FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_map_variables_updated_at BEFORE UPDATE ON map_variables FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_pages_updated_at BEFORE UPDATE ON pages FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
CREATE TRIGGER update_paragraphs_updated_at BEFORE UPDATE ON paragraphs FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
