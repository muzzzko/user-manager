-- +goose Up
create table user_profile (
    id char(36) PRIMARY KEY not null,
    first_name varchar(256) not null,
    last_name varchar(256) not null,
    nickname varchar(256) not null,
    email varchar(256) not null,
    password_hash char(128) not null,
    country_id integer not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);

create unique index uidx_email_user_profile on user_profile (email);

create table country (
    id serial PRIMARY KEY not null,
    code char(2)
);

-- +goose Down
drop table user_profile;
drop table country;


