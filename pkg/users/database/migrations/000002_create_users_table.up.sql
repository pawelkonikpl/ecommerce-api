CREATE TABLE IF NOT EXISTS users
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    birth_date TIMESTAMP WITHOUT TIME ZONE,
    street VARCHAR(255),
    products jsonb
)