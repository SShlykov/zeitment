-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS book_events (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    book_id UUID NOT NULL,
    chapter_id UUID,
    page_id UUID,
    paragraph_id UUID,
    event_type VARCHAR(255) NOT NULL,
    is_public BOOLEAN NOT NULL,
    key VARCHAR(255) NOT NULL,
    value TEXT NOT NULL,
    link TEXT,
    link_text TEXT,
    link_type VARCHAR(255),
    link_image TEXT,
    description TEXT,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (chapter_id) REFERENCES chapters(id) ON DELETE SET NULL,
    FOREIGN KEY (page_id) REFERENCES pages(id) ON DELETE SET NULL,
    FOREIGN KEY (paragraph_id) REFERENCES paragraphs(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS book_events;
-- +goose StatementEnd
