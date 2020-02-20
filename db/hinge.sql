CREATE DATABASE hinge;

\c hinge;

CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    api_key varchar(36),
    first_name varchar(255),
    last_name varchar(255)
);

CREATE TABLE relationships (
    initiator_id INTEGER REFERENCES users (id),
    receiver_id INTEGER REFERENCES users (id),
    status_id INTEGER NOT NULL,
    last_updated DATE
);

CREATE TABLE status_types (
    status_id SERIAL PRIMARY KEY,
    status_type varchar(20),
    hide_profile boolean,
    is_recommended boolean
);

-- Test Data
INSERT INTO users (first_name, last_name, api_key) VALUES ('Jon', 'Snow', 'hinge');
INSERT INTO users (first_name, last_name, api_key) VALUES ('Daenerys', 'Targaryen', 'hinge');
INSERT INTO users (first_name, last_name, api_key) VALUES ('Ygritte', '', 'hinge');
INSERT INTO users (first_name, last_name, api_key) VALUES ('Arya', 'Stark', 'hinge');
INSERT INTO users (first_name, last_name, api_key) VALUES ('Tyrion', 'Lannister', 'hinge');
INSERT INTO users (first_name, last_name, api_key) VALUES ('Khal', 'Drogo', 'hinge');
INSERT INTO users (first_name, last_name, api_key) VALUES ('Jorah', 'Mormont', 'hinge');

INSERT INTO status_types (status_type, hide_profile, is_recommended) VALUES ('like', FALSE, FALSE);
INSERT INTO status_types (status_type, hide_profile, is_recommended) VALUES ('dislike', TRUE, FALSE);
INSERT INTO status_types (status_type, hide_profile, is_recommended) VALUES ('match', FALSE, FALSE);
INSERT INTO status_types (status_type, hide_profile, is_recommended) VALUES ('block', TRUE, FALSE);
INSERT INTO status_types (status_type, hide_profile, is_recommended) VALUES ('report', TRUE, FALSE);

-- Assuming sequential order from above inserts
INSERT INTO relationships VALUES (1, 3, 3, NOW());
INSERT INTO relationships VALUES (2, 1, 1, NOW());
INSERT INTO relationships VALUES (5, 2, 1, NOW());
INSERT INTO relationships VALUES (6, 2, 1, NOW());
INSERT INTO relationships VALUES (7, 2, 1, NOW());