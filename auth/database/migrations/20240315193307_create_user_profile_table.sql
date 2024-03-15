-- +goose Up
-- +goose StatementBegin
create table user_profile
(
    id      uuid primary key default uuid_generate_v4(),
    user_id uuid references users(id) on delete cascade,
    first_name         varchar(255),
    last_name          varchar(255),
    second_name        varchar(255),
    locale             varchar(255),
    gender             integer,
    phone              varchar(255),
    notes              varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_profile;
-- +goose StatementEnd
