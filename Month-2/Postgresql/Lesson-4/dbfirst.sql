CREATE DATABASE library;
\c library
\password -- if you want change your postgres password or set password

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    pages INT,
    author_id INT REFERENCES authors(id)
);

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    birth_day INT
);