-- +migrate Up
CREATE TABLE
  notes (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    created timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted timestamptz,
    "userId" uuid NOT NULL,
    title text NOT NULL,
    content text NOT NULL
  );

CREATE TRIGGER set_timestamp BEFORE
UPDATE
  ON notes FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp ();
