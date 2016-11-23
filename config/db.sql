-- DROP DB IF EXISTS

DROP DATABASE IF EXISTS garbanzo;

-- CREATE DB

CREATE DATABASE garbanzo;

\c garbanzo

-- CREATE USER TABLE

CREATE TABLE "user" (
  id             SERIAL NOT NULL,
  name           VARCHAR(20) NOT NULL,
  surname        VARCHAR(20) NOT NULL,
  email          VARCHAR(20) NOT NULL,

  created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION update_modified_column()
  RETURNS TRIGGER AS $$
  BEGIN
    NEW.updated_at = now();
    return NEW;
  END
$$ language 'plpgsql';

CREATE TRIGGER user_updated_at_modify BEFORE UPDATE ON "user" FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
