CREATE DATABASE alprov;
\c alprov
CREATE TABLE students(
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(100),
    username VARCHAR(35) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE teachers(
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(100),
    username VARCHAR(35) UNIQUE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE courses(
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    teacher INT REFERENCES teachers(id),
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE students_courses(
    student_id INT,
    course_id INT
);
