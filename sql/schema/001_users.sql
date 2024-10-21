-- +goose up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    fullname VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


-- +goose down
DROP TABLE users;