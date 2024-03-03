-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS map_variables (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    book_id UUID NOT NULL,
    chapter_id UUID,
    page_id UUID,
    paragraph_id UUID,
    map_link TEXT NOT NULL,
    lat DOUBLE PRECISION NOT NULL,
    lng DOUBLE PRECISION NOT NULL,
    zoom INTEGER,
    date varchar(255),
    description TEXT,
    link TEXT,
    link_text TEXT,
    link_type VARCHAR(255),
    link_image TEXT,
    image TEXT,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (chapter_id) REFERENCES chapters(id) ON DELETE SET NULL,
    FOREIGN KEY (page_id) REFERENCES pages(id) ON DELETE SET NULL,
    FOREIGN KEY (paragraph_id) REFERENCES paragraphs(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
