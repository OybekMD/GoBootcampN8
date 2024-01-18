CREATE DATABASE shopify;
\c shopify

CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    fname VARCHAR(65),
    lname VARCHAR(65),
    balance INT DEFAULT 0
);

CREATE TABLE products(
    id SERIAL PRIMARY KEY,
    name VARCHAR(35),
    price INT,
    created_at TIMESTAMP DEFAULT NOW(),
    isactive BOOLEAN DEFAULT true
);

CREATE TABLE users_products(
    user_id INT REFERENCES users(id),
    product_id INT REFERENCES products(id)
);