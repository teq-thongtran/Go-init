CREATE TABLE IF NOT EXISTS cards
(
    `id`         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    `name_card`       VARCHAR(255) NOT NULL,
    `card_type` VARCHAR(255),
    `user_id` BIGINT(20),
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL     DEFAULT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
