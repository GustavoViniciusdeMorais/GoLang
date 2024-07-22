\c gopos;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    age INT NOT NULL,
    passwrod VARCHAR(255) NOT NULL
);

INSERT INTO users (name, email, age, password) VALUES ('Admin User', 'admin@example.com', 30, 'admin123');
