CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE user_values (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    key VARCHAR(255) NOT NULL,
    value VARCHAR(255) NOT NULL
);

INSERT INTO users (email, password) VALUES
('user10@test.com', 'test'),
('user25@test.com', 'test'),
('user50@test.com', 'test'),
('user100@test.com', 'test'),
('user200@test.com', 'test');