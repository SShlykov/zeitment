-- +goose Up
-- +goose StatementBegin
CREATE TABLE chapters (
   id UUID PRIMARY KEY,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
   updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
   deleted_at TIMESTAMP WITH TIME ZONE,
   title VARCHAR(255) NOT NULL,
   number INTEGER NOT NULL,
   text TEXT NOT NULL,
   book_id UUID NOT NULL,
   is_public BOOLEAN NOT NULL DEFAULT false,
   FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chapters;
-- +goose StatementEnd
