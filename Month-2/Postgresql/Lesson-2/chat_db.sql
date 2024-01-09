sudo -i -u postgres
psql
CREATE DATABASE telegram;
\c telegram
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(15) NOT NULL,
    fullname VARCHAR(30)
);

CREATE TABLE chats(
    id SERIAL PRIMARY KEY,
    groupname VARCHAR(30) NOT NULL
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    content VARCHAR(1000) NOT NULL,
    sent_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    sender_id INT,
    FOREIGN KEY (sender_id) REFERENCES users(id)
);

CREATE TABLE users_chats (
    user_id INT,
    chat_id INT,
    PRIMARY KEY (user_id, chat_id),
    CONSTRAINT "fk_user" FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT "fk_chat" FOREIGN KEY (chat_id) REFERENCES chats(id)
);

CREATE TABLE user_conversation_messages (
    user_id INT,
    chat_id INT,
    message_id INT,
    PRIMARY KEY (user_id, chat_id, message_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (chat_id) REFERENCES chats(id),
    FOREIGN KEY (message_id) REFERENCES messages(id)
);

-- Insert to users START

insert into users (username, fullname) values ('djelliman0', 'Damien Jelliman');
insert into users (username, fullname) values ('meppson1', 'Marjy Eppson');
insert into users (username, fullname) values ('sgotcher2', 'Smitty Gotcher');
insert into users (username, fullname) values ('sfoxall3', 'Sallyann Foxall');
insert into users (username, fullname) values ('bcometson4', 'Brendis Cometson');
insert into users (username, fullname) values ('cpinckney5', 'Ches Pinckney');
insert into users (username, fullname) values ('swainscot6', 'Sonnie Wainscot');
insert into users (username, fullname) values ('sbrainsby7', 'Sybyl Brainsby');
insert into users (username, fullname) values ('alomax8', 'Avrom Lomax');
insert into users (username, fullname) values ('kcasiero9', 'Kingston Casiero');
insert into users (username, fullname) values ('lblackwella', 'Lawrence Blackwell');
insert into users (username, fullname) values ('dmcmeartyb', 'Dianemarie McMearty');
insert into users (username, fullname) values ('ffinkc', 'Foss Fink');
insert into users (username, fullname) values ('lpenwrightd', 'Lyn Penwright');
insert into users (username, fullname) values ('gwordene', 'Gwennie Worden');
insert into users (username, fullname) values ('ibonhamf', 'Ivan Bonham');
insert into users (username, fullname) values ('mflatherg', 'Mignonne Flather');
insert into users (username, fullname) values ('oharesnapeh', 'Oliy Haresnape');
insert into users (username, fullname) values ('kleggani', 'Keven Leggan');
insert into users (username, fullname) values ('dgrassotj', 'Dniren Grassot');

-- Get all usernames 
SELECT * FROM users;

-- Insert to users END


-- Insert to chats START

insert into chats (groupname) values ('Gocoders');
insert into chats (groupname) values ('phplovers');
insert into chats (groupname) values ('outside fights');
insert into chats (groupname) values ('psql is better');

-- Insert to chats END

-- Get all usernames 
SELECT * FROM chats;



-- Insert to Users enter to chat START

insert into users_chats (user_id, chat_id) values (1, 1);
insert into users_chats (user_id, chat_id) values (2, 1);
insert into users_chats (user_id, chat_id) values (3, 1);
insert into users_chats (user_id, chat_id) values (4, 1);
insert into users_chats (user_id, chat_id) values (5, 1);
insert into users_chats (user_id, chat_id) values (6, 1);
insert into users_chats (user_id, chat_id) values (7, 1);
insert into users_chats (user_id, chat_id) values (8, 1);
insert into users_chats (user_id, chat_id) values (9, 2);
insert into users_chats (user_id, chat_id) values (10, 2);
insert into users_chats (user_id, chat_id) values (11, 2);
insert into users_chats (user_id, chat_id) values (12, 2);
insert into users_chats (user_id, chat_id) values (13, 2);
insert into users_chats (user_id, chat_id) values (14, 3);
insert into users_chats (user_id, chat_id) values (15, 3);
insert into users_chats (user_id, chat_id) values (16, 3);
insert into users_chats (user_id, chat_id) values (17, 4);
insert into users_chats (user_id, chat_id) values (18, 4);
insert into users_chats (user_id, chat_id) values (19, 4);
insert into users_chats (user_id, chat_id) values (20, 4);

-- Insert to Users enter to chat END

-- Get all users joined chanes
SELECT * FROM users_chats;

-- Get joined users
SELECT u.fullname, c.groupname FROM users_chats uc JOIN users u ON uc.user_id = u.id JOIN chats c ON uc.chat_id = c.id;



-- Insert to Messages START

insert into messages (content, sender_id) values ('Go is best of the best', 1);
insert into messages (content, sender_id) values ('Yes I agree to you Damien', 2);
insert into messages (content, sender_id) values ('php is better than Go guys!', 9);
insert into messages (content, sender_id) values ('No Go is better than php', 3);
insert into messages (content, sender_id) values ('but both use postgres lol )', 19);

-- Insert to Messages END

-- Get all messages
SELECT * FROM messages;



-- Insert to START

INSERT INTO user_conversation_messages (user_id, chat_id, message_id)
VALUES (1, 1, 1),
(2, 1, 2),
(9, 1, 3),
(3, 1, 4),
(19, 1, 5);

-- Insert to END

-- Get all user conversation messages
SELECT * FROM user_conversation_messages;

-- Get all user conversation messages which message u
SELECT m.id ,u.fullname, c.groupname, m.content FROM user_conversation_messages ucm JOIN users u ON ucm.user_id = u.id JOIN chats c ON ucm.chat_id = c.id JOIN messages m ON ucm.message_id = m.id;
