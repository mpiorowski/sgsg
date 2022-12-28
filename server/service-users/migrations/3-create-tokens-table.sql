-- +migrate Up
CREATE TABLE
  tokens (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    created timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted boolean NOT NULL DEFAULT FALSE,
    code text NOT NULL,
    email text UNIQUE NOT NULL,
    expires timestamptz NOT NULL
  );

-- +migrate Up
CREATE TRIGGER set_timestamp BEFORE
UPDATE
  ON tokens FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp ();
