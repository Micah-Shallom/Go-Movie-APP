-- This is just an example of raw SQL that does the equivilent of the first 30 lines of models.go
-- Create Director table
CREATE TABLE IF NOT EXISTS directors (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    INDEX (deleted_at)
);

-- Create Movie table
CREATE TABLE IF NOT EXISTS movies (
    id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    isbn VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    director_id INT UNSIGNED,
    FOREIGN KEY (director_id) REFERENCES directors(id) ON DELETE CASCADE,
    INDEX (deleted_at)
);
