-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id uuid primary key not null,
    email varchar(255) not null,
    password varchar(16) not null
);

alter table tasks
    add constraint fk_userID foreign key (userID) references users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table tasks
    drop constraint fk_userID;

drop table if exists users;
-- +goose StatementEnd
