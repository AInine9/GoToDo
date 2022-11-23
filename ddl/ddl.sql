CREATE TABLE `items`
(
    `id`         bigint(20) NOT NULL AUTO_INCREMENT,
    `name`       varchar(255),
    `status`     int      NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
