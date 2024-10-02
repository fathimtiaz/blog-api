CREATE TABLE IF NOT EXISTS user_ (
    id SERIAL PRIMARY KEY,
    name_ VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS post_ (
    id SERIAL PRIMARY KEY,
    author_id BIGINT UNSIGNED,
    content TEXT NOT NULL,
    deleted BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (author_id) REFERENCES user_ (id)
);

CREATE TABLE IF NOT EXISTS post_comment_ (
    id SERIAL PRIMARY KEY,
    post_id BIGINT UNSIGNED,
    author_name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    FOREIGN KEY (post_id) REFERENCES post_ (id)
);