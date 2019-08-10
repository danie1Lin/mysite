CREATE TABLE users
(
    id uuid DEFAULT uuid_generate_v4 (),
    email VARCHAR NOT NULL UNIQUE,
    pswd VARCHAR NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
