CREATE DATABASE payuz;
\c payuz

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    fname VARCHAR(65),
    lname VARCHAR(65),
    balance INT DEFAULT 0
);