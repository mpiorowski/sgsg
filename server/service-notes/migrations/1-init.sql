-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp ()
  RETURNS TRIGGER
  AS $$
BEGIN
  NEW.updated = NOW();
  RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd
