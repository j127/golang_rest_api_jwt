-- This SQL file is only run if there is no data directory already.
-- See here for more information: https://hub.docker.com/_/postgres
--
--      Warning: scripts in /docker-entrypoint-initdb.d are only run if
--      you start the container with a data directory that is empty; any
--      pre-existing database will be left untouched on container startup. One
--      common problem is that if one of your /docker-entrypoint-initdb.d
--      scripts fails (which will cause the entrypoint script to exit) and your
--      orchestrator restarts the container with the already initialized data
--      directory, it will not continue on with your scripts.

-- This table holds the messages from Amazon SNS that we want to save for a while.
CREATE TABLE users (
    id serial NOT NULL PRIMARY KEY,
    email text NOT NULL unique,
    password text NOT NULL
);

INSERT INTO users (email, password) VALUES ('alice@example.com', '12345');
INSERT INTO users (email, password) VALUES ('bob@example.com', '54321');
