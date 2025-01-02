-- +goose Up
-- +goose StatementBegin
create table if not exists tasks
(
    taskID uuid primary key not null,
    userID uuid not null,
    type text,
    name text,
    description text not null,
    status smallint not null,
    createdAt timestamptz default now(),
    updatedAt timestamptz default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd
