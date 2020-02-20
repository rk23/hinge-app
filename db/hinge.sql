CREATE DATABASE hinge;

\c hinge;

CREATE TABLE users (
    id varchar(36) NOT NULL PRIMARY KEY,
    first_name varchar(255),
    last_name varchar(255)
);

CREATE TABLE relationships (
    initiator_id varchar(36) REFERENCES users (id),
    receiver_id varchar(36) REFERENCES users (id),
    status_id int NOT NULL
);

CREATE TABLE status_types (
    status_id SERIAL PRIMARY KEY,
    status_type varchar(20),
    hide_profile boolean
);