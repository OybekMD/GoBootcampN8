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











insert into authors (name, birth_day) values ('Harriet Lytle', 1984);
insert into authors (name, birth_day) values ('Rosana Roycraft', 1998);
insert into authors (name, birth_day) values ('Peria Clace', 2009);
insert into authors (name, birth_day) values ('Quint Saunier', 2001);
insert into authors (name, birth_day) values ('Nisse Klemz', 1995);

