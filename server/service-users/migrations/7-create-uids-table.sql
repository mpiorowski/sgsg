-- +migrate Up
CREATE TABLE
  uids (
    uid text PRIMARY KEY NOT NULL,
    created timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted timestamptz,

    "userId" uuid REFERENCES users (id) ON DELETE RESTRICT NOT NULL
  );

CREATE TRIGGER set_timestamp BEFORE
UPDATE
  ON uids FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp ();
