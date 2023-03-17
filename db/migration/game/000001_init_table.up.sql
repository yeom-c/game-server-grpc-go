CREATE TABLE `asset` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `type` int(5) unsigned NOT NULL DEFAULT 0,
    `balance` bigint(20) unsigned NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_enum_id` (`account_user_id`,`enum_id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_type` (`account_user_id`,`type`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `character` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `exp` int(10) unsigned NOT NULL DEFAULT 0,
    `equipment_level` int(10) unsigned NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `character_broadcast` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `character_enum_id` varchar(255) NOT NULL,
    `timeline_enum_id` varchar(255) NOT NULL,
    `type` int(5) unsigned NOT NULL DEFAULT 0,
    `on_air` int(2) unsigned NOT NULL DEFAULT 0,
    `complete` int(2) unsigned NOT NULL DEFAULT 0,
    `broadcasted_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_character_enum_id_timeline_enum_id` (`account_user_id`,`character_enum_id`,`timeline_enum_id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE,
    KEY `idx_character_enum_id` (`character_enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `character_collection` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `character_enum_id` varchar(255) NOT NULL,
    `affection_exp` int(10) NOT NULL DEFAULT 0,
    `count` int(10) unsigned NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_character_enum_id` (`account_user_id`,`character_enum_id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `costume` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `character_enum_id` varchar(255) NOT NULL,
    `state` int(10) unsigned NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_enum_id` (`account_user_id`,`enum_id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`),
    KEY `idx_account_user_id_character_enum_id` (`account_user_id`,`character_enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `deck` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `index` int(5) NOT NULL DEFAULT 0,
    `name` varchar(50) NOT NULL DEFAULT '',
    `character_id_0` int(10) unsigned DEFAULT NULL,
    `character_id_1` int(10) unsigned DEFAULT NULL,
    `character_id_2` int(10) unsigned DEFAULT NULL,
    `character_id_3` int(10) unsigned DEFAULT NULL,
    `character_id_4` int(10) unsigned DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_index` (`account_user_id`,`index`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `fate_card` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `character_enum_id` varchar(255) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`),
    KEY `idx_account_user_id_character_enum_id` (`account_user_id`,`character_enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `gacha_log` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `character_enum_id` varchar(255) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE,
    KEY `idx_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `item` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `enum_id` varchar(255) NOT NULL,
    `count` int(10) unsigned NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_enum_id` (`account_user_id`,`enum_id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `mail` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `sender` varchar(255) NOT NULL,
    `type` int(5) unsigned NOT NULL DEFAULT 0,
    `status` int(5) unsigned NOT NULL DEFAULT 0,
    `delete_all` int(2) unsigned NOT NULL DEFAULT 0,
    `attachment` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`attachment`)),
    `title` varchar(255) NOT NULL DEFAULT '',
    `message` varchar(255) DEFAULT NULL,
    `expired_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) NOT NULL,
    `story_index` int(10) unsigned NOT NULL DEFAULT 0,
    `tutorial_info` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '[]' CHECK (json_valid(`tutorial_info`)),
    `shop_info` longtext DEFAULT NULL CHECK (json_valid(`shop_info`)),
    `daily_reset_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `broadcast_reset_at` timestamp NOT NULL DEFAULT '2023-01-01 00:00:00',
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_account_user_id` (`account_user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
