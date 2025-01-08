-- +goose Up
-- +goose StatementBegin
insert into users (id, email, password) values ('e05fa11d-eec3-4fba-b223-d6516800a047', 'common@example.com', 'password');
insert into users (id, email, password) values ('3966749e-45d4-460d-8e59-34235672f03b', 'common@example.com', 'password');
insert into users (id, email, password) values ('b3f7c269-1e35-4139-b882-2ec0b6629f7e', 'common@example.com', 'password');
insert into users (id, email, password) values ('70d77738-2f5f-447e-8fa3-c36b238d9301', 'common@example.com', 'password');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from users where name like 'Alex%'
-- +goose StatementEnd
