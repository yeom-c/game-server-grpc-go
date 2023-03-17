CREATE TABLE `account` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uuid` char(36) NOT NULL,
    `world_id` int(10) NOT NULL DEFAULT 0,
    `profile_idx` int(5) unsigned DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_uuid` (`uuid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `account_user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_id` int(10) unsigned NOT NULL,
    `game_db` int(5) unsigned NOT NULL DEFAULT 0,
    `nickname` varchar(50) NOT NULL,
    `signed_in_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_nickname` (`nickname`) USING BTREE,
    KEY `idx_account_id` (`account_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `sharding` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `game_db` int(5) unsigned NOT NULL,
    `count` int(10) unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_game_db` (`game_db`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `shop` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `type` int(5) unsigned NOT NULL,
    `visible` int(2) unsigned NOT NULL DEFAULT 0,
    `name` varchar(255) NOT NULL DEFAULT '',
    `desc` text NOT NULL DEFAULT '',
    `start_at` timestamp NULL DEFAULT NULL,
    `end_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `shop_category` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `shop_id` int(10) unsigned NOT NULL,
    `visible` int(2) unsigned NOT NULL DEFAULT 0,
    `name` varchar(255) NOT NULL,
    `order` int(10) unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_shop_id` (`shop_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `shop_goods` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `shop_category_id` int(10) unsigned NOT NULL,
    `type` int(5) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `info` varchar(255) DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `desc` text NOT NULL,
    `cost_type` int(5) unsigned NOT NULL DEFAULT 0,
    `cost_enum_id` varchar(255) NOT NULL DEFAULT '',
    `cost` int(10) unsigned NOT NULL DEFAULT 0,
    `original_cost` int(10) unsigned NOT NULL DEFAULT 0,
    `count` int(10) unsigned NOT NULL DEFAULT 0,
    `visible` int(2) unsigned NOT NULL DEFAULT 0,
    `start_at` timestamp NULL DEFAULT NULL,
    `end_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_shop_category_id` (`shop_category_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
