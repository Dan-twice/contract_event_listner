-- +migrate Up

CREATE TABLE events(
  id bigserial not null primary key,
  works varchar not null
);

-- +migrate Down

DROP TABLE IF EXISTS events;
