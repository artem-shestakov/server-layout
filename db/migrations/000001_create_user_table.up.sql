CREATE TABLE IF NOT EXISTS roles(
    id SERIAL PRIMARY KEY,
    name VARCHAR(15) UNIQUE
);

INSERT INTO roles (name) VALUES ('Administrator');
INSERT INTO roles (name) VALUES ('User');

CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    firstname VARCHAR(255),
    lastname VARCHAR(255),
    email TEXT NOT NULL UNIQUE,
    role VARCHAR(15) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_login TIMESTAMP,
    last_password_change TIMESTAMP,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT fk_roless FOREIGN KEY(role) REFERENCES roles(name)
);
