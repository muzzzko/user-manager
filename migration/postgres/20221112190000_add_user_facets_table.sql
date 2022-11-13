-- +goose Up
create table user_facets (
    id char(36) PRIMARY KEY not null,
    tsv tsvector not null
);

CREATE INDEX user_facets_tsv_idx ON user_facets USING GIN (tsv);

-- +goose Down
drop table user_facets;
