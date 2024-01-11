sudo -i -u postgres
psql
CREATE USER exemtest WITH SUPERUSER CREATEDB CREATEROLE;
exit
exit
psql -U exemtest postgres
CREATE DATABASE Movies;
\c movies

CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    age INT,
    salary INT
);


CREATE TABLE films (
    id SERIAL PRIMARY KEY,
    name VARCHAR(65),
    filming_year TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    genre VARCHAR(65),
    FOREIGN KEY (id) REFERENCES actors(id)
);

CREATE TABLE actors_films (
    actors_id INT,
    films_id INT,
    PRIMARY KEY (actors_id, films_id),
    CONSTRAINT "fk_actor" FOREIGN KEY (actors_id) REFERENCES actors(id),
    CONSTRAINT "fk_film" FOREIGN KEY (films_id) REFERENCES films(id)
);

INSERT INTO actors (name, age, salary) values ('Tim cruz', 50, 100000);
INSERT INTO actors (name, age, salary) values ('Jon Jakson', 25, 20000);
INSERT INTO actors (name, age, salary) values ('Flip Konor', 15, 10000);
INSERT INTO actors (name, age, salary) values ('Umit Aka', 35, 500);
INSERT INTO actors (name, age, salary) values ('Qutbidin Aka', 35, 500);

INSERT INTO films (name, genre) values ('Sening Omadi chopmagan dosting', 'Comedia');
INSERT INTO films (name, genre) values ('Baxorda yana uchrashguncha', 'Drama');
INSERT INTO films (name, genre) values ('Forjumanla', 'Action');

INSERT INTO actors_films(actors_id, films_id) values 
    (1, 2),
    (2, 2),
    (3, 3),
    (4, 1),
    (5, 1);

SELECT a.name, f.name FROM actors_films af JOIN actors a ON af.actors_id = a.id JOIN films f ON af.films_id = f.id;

SELECT DISTINCT f.name, f.filming_year FROM actors_films af JOIN actors a ON af.actors_id = a.id JOIN films f ON af.films_id = f.id;