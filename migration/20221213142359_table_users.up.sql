CREATE TABLE IF NOT EXISTS users
(
    `id`         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `name`       VARCHAR(255) NOT NULL,
    `username` VARCHAR(255),
    `email` VARCHAR(255),
    `score` INT DEFAULT 0,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL     DEFAULT NULL,

    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;