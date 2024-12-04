CREATE TABLE IF NOT EXISTS user (
    id VARCHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    nickname VARCHAR(50) NOT NULL UNIQUE,
    status ENUM('enabled', 'disabled') NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
