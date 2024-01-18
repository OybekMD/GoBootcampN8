CREATE DATABASE idx;
\c idx 

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) -- Adjust the length as per your requirements
);

-- Insert random data into the table
INSERT INTO employees (name)
SELECT
    md5(random()::text) -- Using MD5 hash of random text as a simple way to generate random strings
FROM
    generate_series(1, 11000000);

-- Optional: Analyze the table to update statistics
ANALYZE employees;


explain analyze select * from table_name where id = 10000;



-- Neccessary
SELECT sha256('hello world');
select array_to_string(array(SELECT chr((32 + random() * 94)::integer) FROM generate_series(1, 5)), ''); -- data output: {ASD&
SELECT array_to_string(array(SELECT chr(65 + floor(random() * 26)::integer) FROM generate_series(1, 5)), ''); -- data output: ASVEW

-- My commands ---------

CREATE TABLE employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255)
);

INSERT INTO employees(name)
SELECT array_to_string(array(SELECT chr(65 + floor(random() * 26)::integer) FROM generate_series(1, 5)), '')
FROM generate_series(1, 1000000);