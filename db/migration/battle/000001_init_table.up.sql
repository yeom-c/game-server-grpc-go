CREATE TABLE `battle_result` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) unsigned NOT NULL,
    `match_account_user_id` int(10) unsigned NOT NULL DEFAULT 0,
    `channel_id` varchar(255) NOT NULL DEFAULT '',
    `deck_id` int(10) unsigned NOT NULL DEFAULT 0,
    `result` int(5) unsigned NOT NULL DEFAULT 0,
    `battle_start_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `confirmed_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uniq_account_user_id_channel_id` (`account_user_id`,`channel_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `account_user_id` int(10) NOT NULL,
    `match_point` int(10) unsigned NOT NULL DEFAULT 0,
    `match_win` int(10) unsigned NOT NULL DEFAULT 0,
    `match_lose` int(10) unsigned NOT NULL DEFAULT 0,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `idx_account_user_id` (`account_user_id`) USING BTREE,
    KEY `idx_match_point` (`match_point`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
