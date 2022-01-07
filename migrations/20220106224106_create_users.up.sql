create table users (
    id serial not null primary key,
    email varchar not null unique,
    encrypted_password varchar not null
);