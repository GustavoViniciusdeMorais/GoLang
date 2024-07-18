\c gopos;

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    age INT NOT NULL
);

INSERT INTO users (name, email, age) VALUES ('Admin User', 'admin@example.com', 30);
