-- +migrate Up
alter table users
  add column if not exists "uid" varchar not null default '';
