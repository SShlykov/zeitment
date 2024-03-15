-- +goose Up
-- +goose StatementBegin
BEGIN TRANSACTION;
create table if not exists users
(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP WITH TIME ZONE,

    logged_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    confirmed_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),

    login              varchar(255),
    email              varchar(255),
    password_hash      varchar(255),

    deleted_by         integer,
    access_template_id integer,
    update_after       bigint
);

create unique index users_login_index   on users (login);
create unique index users_email_index   on users (email);
END TRANSACTION;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN TRANSACTION;
drop table if exists users;
drop index if exists users_login_index;
drop index if exists users_email_index;
END TRANSACTION;
-- +goose StatementEnd
