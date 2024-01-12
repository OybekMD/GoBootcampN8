CREATE DATABASE library;
\c library
\password -- if you want change your postgres password or set password

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    birth_day INT
);

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    pages INT
);

CREATE TABLE authors_books (
    id SERIAL PRIMARY KEY,
    author_id INT REFERENCES authors(id),
    books_id INT REFERENCES books(id)
);