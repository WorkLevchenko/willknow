CREATE TABLE users  (
    id bigserial not null primary key,
    email varchar  unique,
    encrypted_password varchar not null
);

