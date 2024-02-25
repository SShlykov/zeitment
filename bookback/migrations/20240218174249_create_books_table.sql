-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    owner INTEGER NOT NULL,
    description TEXT,
    is_public BOOLEAN NOT NULL DEFAULT false,
    publication TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
