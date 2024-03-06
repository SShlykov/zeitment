-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    owner UUID NOT NULL,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    description TEXT,
    is_public BOOLEAN NOT NULL,
    publication TIMESTAMP WITH TIME ZONE,
    image_link TEXT,
    map_link TEXT,
    map_params_id UUID,
    variables TEXT[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books;
-- +goose StatementEnd
