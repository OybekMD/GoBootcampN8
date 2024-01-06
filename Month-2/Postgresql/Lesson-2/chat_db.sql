CREATE DATABASE telegram;
\c telegram
CREATE TABLE users(
    id serial, 
    username varchar(15),
    fullname varchar(30)
);

CREATE TABLE chats(
    id serial,
    groupname varchar(30)
    members fk
)