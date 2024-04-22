package dtos

import "time"

type User struct {
	ID        int       `json:"-" db:"id"`
	Username  string    `json:"username" db:"username" binding:"required"`
	Email     string    `json:"email" db:"email" binding:"required"`
	Password  string    `json:"password" db:"password_hash" binding:"required"`
	CreatedAt time.Time `db:"created_at"`
}

type SignInInput struct {
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

// -- Создание таблицы пользователей
// CREATE TABLE users (
//     id SERIAL PRIMARY KEY,
//     username VARCHAR(255) NOT NULL,
//     email VARCHAR(255) NOT NULL,
//     password VARCHAR(255) NOT NULL,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// -- Создание таблицы администраторов
// CREATE TABLE admin (
//     id SERIAL PRIMARY KEY,
//     user_id INT UNIQUE REFERENCES users(id)
// );

// -- Создание таблицы клиентов
// CREATE TABLE client (
//     id SERIAL PRIMARY KEY,
//     user_id INT UNIQUE REFERENCES users(id)
// );
