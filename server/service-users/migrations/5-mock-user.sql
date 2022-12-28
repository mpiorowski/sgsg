-- +migrate Up
insert into users (id, email, role) values ('39eedc66-d189-4d01-bf6a-167d6ff270b5', '{mat@gmail.com}', 'ROLE_ADMIN');
