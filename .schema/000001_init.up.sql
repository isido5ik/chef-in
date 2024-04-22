-- Создание таблицы пользователей
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы администраторов
CREATE TABLE admin (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id)
);

-- Создание таблицы клиентов
CREATE TABLE client (
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id)
);
