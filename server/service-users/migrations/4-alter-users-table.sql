-- +migrate Up
alter table users 
    alter email drop default,
    alter email type text[] using array[email],
    alter email set default '{}';
