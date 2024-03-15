-- +goose Up
-- +goose StatementBegin
create table roles
(
  id serial primary key,
  name text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
