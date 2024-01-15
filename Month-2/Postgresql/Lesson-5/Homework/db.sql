CREATE DATABASE najottalimhomew;
\c najottalimhomew

CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    birthyear INT
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(35),
    teacher VARCHAR(65)
);

CREATE TABLE student_group (
    id SERIAL PRIMARY KEY,
    student_id INT REFERENCES students(id),
    group_id INT REFERENCES groups(id)
);