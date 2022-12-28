-- +migrate Up
CREATE TABLE
  files (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    created timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted boolean NOT NULL DEFAULT FALSE,
    "targetId" uuid NOT NULL,
    "name" text NOT NULL,
    "type" text NOT NULL
  );

-- +migrate Up
CREATE TRIGGER set_timestamp BEFORE
UPDATE
  ON files FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp ();

