-- Создание таблицы пользователей
CREATE TABLE t_users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE t_admin(
    admin_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id) 
);

CREATE TABLE t_client(
    client_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id)
);

CREATE TABLE t_roles(
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR(100) NOT NULL
);

CREATE TABLE t_users_roles (
    user_id INTEGER REFERENCES t_users(user_id),
    role_id INTEGER REFERENCES t_roles(role_id)
);

-- Создание таблицы постов
CREATE TABLE t_posts (
    post_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id),
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы комментариев
CREATE TABLE t_comments (
    comment_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id),
    post_id INTEGER REFERENCES t_posts(post_id),
    parent_id INTEGER REFERENCES t_comments(comment_id),
    comment_text TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы лайков
CREATE TABLE t_likes (
    like_id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES t_users(user_id),
    post_id INTEGER REFERENCES t_posts(post_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
