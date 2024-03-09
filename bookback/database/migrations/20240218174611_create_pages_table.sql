-- +goose Up
-- +goose StatementBegin
CREATE TABLE pages (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
  deleted_at TIMESTAMP WITH TIME ZONE,
  text TEXT NOT NULL,
  chapter_id UUID NOT NULL,
  is_public BOOLEAN NOT NULL DEFAULT false,
  map_params_id UUID,
  FOREIGN KEY (chapter_id) REFERENCES chapters(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
