\c gopos;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    birthday DATE NULL,
    password VARCHAR(255) NOT NULL,
    active BOOLEAN DEFAULT false,
);

INSERT INTO users (name, email, birthday, password, active)
VALUES ('Admin User', 'admin@example.com', '1980-01-01', '$2a$14$ozETnrP2tD5o1dU85MHT6.nHZmvN1Z7DsxWvYBnO4V5rAo4RZNwzG', TRUE);
