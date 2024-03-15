-- +goose Up
-- +goose StatementBegin
create table if not exists user_tokens
(
    id          uuid primary key default uuid_generate_v4(),
    user_id     uuid references users(id) on delete cascade,
    token       bytea        not null,
    context     varchar(255) not null,
    sent_to     varchar(255),
    inserted_at timestamp(0) not null
);

create index users_tokens_user_id_index
    on user_tokens (user_id);

create unique index users_tokens_context_token_index
    on user_tokens (context, token);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
