-- +goose Up
-- +goose StatementBegin
insert into users (id, name) values ('e05fa11d-eec3-4fba-b223-d6516800a047', 'Alex1');
insert into users (id, name) values ('3966749e-45d4-460d-8e59-34235672f03b', 'Alex2');
insert into users (id, name) values ('b3f7c269-1e35-4139-b882-2ec0b6629f7e', 'Alex3');
insert into users (id, name) values ('70d77738-2f5f-447e-8fa3-c36b238d9301', 'Alex4');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from users where name like 'Alex%'
-- +goose StatementEnd
