CREATE TABLE IF NOT EXISTS posts (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    title LONGTEXT,
    content LONGTEXT,
    created_at DATETIME,
    updated_at DATETIME,
    deleted_at DATETIME
);