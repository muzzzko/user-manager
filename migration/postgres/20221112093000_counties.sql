-- +goose Up
insert into country (code) values
    ('UK'), ('FR'), ('US');

-- +goose Down
truncate country;


