sudo -i -u postgres
psql
CREATE DATABASE users;
\c users


CREATE TABLE users (
    id UUID PRIMARY KEY,
    id_num SERIAL,
    name text,
    last_name text
);






select * from users order by id_num desc;

-- Queries
select * from users;
select * from users limit 10;
select * from users limit 10 OFFSET 1;
-- limit 10 page 1;

-- boshidan 10 ta
select * from users limit 10 offset 0;

-- boshidan 10 ta tashlab keyingi 10 ta
select * from users limit 10 offset 10;

-- boshidan 20 ta tashlab keyingi 10 ta
select * from users limit 10 offset 20;