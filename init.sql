CREATE TABLE IF NOT EXISTS benchmark_users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    backend_tag VARCHAR(50) NOT NULL,
    values_count INTEGER NOT NULL,
    UNIQUE (email, backend_tag)
);

INSERT INTO benchmark_users (email, password, backend_tag, values_count) VALUES
('user10@test.com', 'test', 'go', 10),
('user25@test.com', 'test', 'go', 25),
('user50@test.com', 'test', 'go', 50),
('user100@test.com', 'test', 'go', 100),
('user200@test.com', 'test', 'go', 200),
('user10@test.com', 'test', 'spring', 10),
('user25@test.com', 'test', 'spring', 25),
('user50@test.com', 'test', 'spring', 50),
('user100@test.com', 'test', 'spring', 100),
('user200@test.com', 'test', 'spring', 200),
('user10@test.com', 'test', 'django_ninja', 10),
('user25@test.com', 'test', 'django_ninja', 25),
('user50@test.com', 'test', 'django_ninja', 50),
('user100@test.com', 'test', 'django_ninja', 100),
('user200@test.com', 'test', 'django_ninja', 200)
ON CONFLICT (email, backend_tag) DO NOTHING;