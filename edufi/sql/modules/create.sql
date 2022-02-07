-- CREATE tables in the edufi database

CREATE TABLE modules(
    id SERIAL PRIMARY KEY,
    name VARCHAR(64) UNIQUE NOT NULL,
    synopis VARCHAR(128) NOT NULL
);

