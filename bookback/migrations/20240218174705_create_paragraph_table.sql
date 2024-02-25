-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS paragraphs (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    text TEXT NOT NULL,
    is_public BOOLEAN NOT NULL,
    page_id UUID NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    FOREIGN KEY (page_id) REFERENCES pages(id) -- Assuming there is a pages table that paragraph relates to
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS paragraphs CASCADE;
-- +goose StatementEnd
