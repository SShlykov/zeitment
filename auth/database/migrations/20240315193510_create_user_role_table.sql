-- +goose Up
-- +goose StatementBegin
create table user_role
(
  id serial primary key,
  user_id uuid not null references users(id) on delete cascade,
  role_id integer not null references roles(id) on delete cascade,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user_role;
-- +goose StatementEnd
