/*
 Navicat Premium Data Transfer

 Source Server         : quasar_dev
 Source Server Type    : MariaDB
 Source Server Version : 100605 (10.6.5-MariaDB-log)
 Source Host           : db-quasar-dev.c380l5msf2mu.ap-northeast-2.rds.amazonaws.com:3306
 Source Schema         : static_data

 Target Server Type    : MariaDB
 Target Server Version : 100605 (10.6.5-MariaDB-log)
 File Encoding         : 65001

 Date: 18/01/2023 14:09:46
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for abyss
-- ----------------------------
DROP TABLE IF EXISTS `abyss`;
CREATE TABLE `abyss` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `map_file` varchar(255) NOT NULL,
  `size` varchar(255) NOT NULL,
  `size_x` int(10) unsigned NOT NULL,
  `size_y` int(10) unsigned NOT NULL,
  `rule` varchar(255) NOT NULL,
  `abyss_title_resource` varchar(255) NOT NULL,
  `boss` varchar(255) NOT NULL,
  `minor` varchar(255) NOT NULL,
  `elite` varchar(255) NOT NULL,
  `tower` varchar(255) NOT NULL,
  `time_limit` double NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for abyss_pool
-- ----------------------------
DROP TABLE IF EXISTS `abyss_pool`;
CREATE TABLE `abyss_pool` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `abyss` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for account_growth
-- ----------------------------
DROP TABLE IF EXISTS `account_growth`;
CREATE TABLE `account_growth` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `exp_plus` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for ai_deck
-- ----------------------------
DROP TABLE IF EXISTS `ai_deck`;
CREATE TABLE `ai_deck` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ai_file` varchar(255) NOT NULL,
  `ai_match_score` int(10) unsigned NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `deck_slot` text NOT NULL,
  `character_level` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for ai_player
-- ----------------------------
DROP TABLE IF EXISTS `ai_player`;
CREATE TABLE `ai_player` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `nickname` varchar(255) NOT NULL,
  `deck` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for ai_pool
-- ----------------------------
DROP TABLE IF EXISTS `ai_pool`;
CREATE TABLE `ai_pool` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ai` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for asset
-- ----------------------------
DROP TABLE IF EXISTS `asset`;
CREATE TABLE `asset` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_asset` varchar(255) NOT NULL,
  `ce_asset_grade` varchar(255) NOT NULL,
  `grade_resource_thumbnail` varchar(255) NOT NULL,
  `grade_resource` varchar(255) NOT NULL,
  `image_reference` varchar(255) NOT NULL,
  `pictogram_reference` varchar(255) NOT NULL,
  `grade_reference` varchar(255) NOT NULL,
  `image_tile` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for balance
-- ----------------------------
DROP TABLE IF EXISTS `balance`;
CREATE TABLE `balance` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `content_type` varchar(255) NOT NULL,
  `fail_value` double NOT NULL,
  `balance` text NOT NULL,
  `rate` text NOT NULL,
  `min_value` text NOT NULL,
  `max_value` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character
-- ----------------------------
DROP TABLE IF EXISTS `character`;
CREATE TABLE `character` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `character_root` varchar(255) NOT NULL,
  `ce_character_species` varchar(255) NOT NULL,
  `crew` varchar(255) NOT NULL,
  `ce_character_class` varchar(255) NOT NULL,
  `ce_character_property` varchar(255) NOT NULL,
  `mbti` varchar(255) NOT NULL,
  `hobby` varchar(255) NOT NULL,
  `partner` tinyint(1) NOT NULL,
  `ce_character_grade` varchar(255) NOT NULL,
  `base_stat` varchar(255) NOT NULL,
  `level_stat` varchar(255) NOT NULL,
  `resource_list` varchar(255) NOT NULL,
  `skill_set` varchar(255) NOT NULL,
  `item_preference` varchar(255) NOT NULL,
  `signature_weapon` varchar(255) NOT NULL,
  `costume_bundle` varchar(255) NOT NULL,
  `second_name` varchar(255) NOT NULL,
  `library` tinyint(1) NOT NULL,
  `profile` varchar(255) NOT NULL,
  `unit_size` double NOT NULL,
  `basic` varchar(255) NOT NULL,
  `active` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_base_stat
-- ----------------------------
DROP TABLE IF EXISTS `character_base_stat`;
CREATE TABLE `character_base_stat` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `base_stat` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_costume_bundle
-- ----------------------------
DROP TABLE IF EXISTS `character_costume_bundle`;
CREATE TABLE `character_costume_bundle` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `costume` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_get_dialogue
-- ----------------------------
DROP TABLE IF EXISTS `character_get_dialogue`;
CREATE TABLE `character_get_dialogue` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `character_enum_id` varchar(255) NOT NULL,
  `animation` text NOT NULL,
  `day_dialogue` text NOT NULL,
  `night_dialogue` text NOT NULL,
  `day_voice` text NOT NULL,
  `night_voice` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_growth
-- ----------------------------
DROP TABLE IF EXISTS `character_growth`;
CREATE TABLE `character_growth` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `exp_plus` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_level_stat
-- ----------------------------
DROP TABLE IF EXISTS `character_level_stat`;
CREATE TABLE `character_level_stat` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `level_stat` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_resource
-- ----------------------------
DROP TABLE IF EXISTS `character_resource`;
CREATE TABLE `character_resource` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `spine` varchar(255) NOT NULL,
  `eff_landing` varchar(255) NOT NULL,
  `character_weapon` varchar(255) NOT NULL,
  `phone_case_close` varchar(255) NOT NULL,
  `phone_case_open` varchar(255) NOT NULL,
  `day_ld_reference` varchar(255) NOT NULL,
  `night_ld_reference` varchar(255) NOT NULL,
  `night_sd_reference` varchar(255) NOT NULL,
  `portrait_night_reference` varchar(255) NOT NULL,
  `portrait_daytime_reference` varchar(255) NOT NULL,
  `thumbnail_reference` varchar(255) NOT NULL,
  `thumbnail_sd_reference` varchar(255) NOT NULL,
  `background_reference` varchar(255) NOT NULL,
  `card_reference` varchar(255) NOT NULL,
  `voice_skill` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for character_type_resource
-- ----------------------------
DROP TABLE IF EXISTS `character_type_resource`;
CREATE TABLE `character_type_resource` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `common_enum_value` varchar(255) NOT NULL,
  `image_ref` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for characterskillset
-- ----------------------------
DROP TABLE IF EXISTS `characterskillset`;
CREATE TABLE `characterskillset` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `skill` text NOT NULL,
  `ce_skill_type` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `comment_writer` text NOT NULL,
  `comment_contents` text NOT NULL,
  `comment_reference` text NOT NULL,
  `comment_like` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for costume
-- ----------------------------
DROP TABLE IF EXISTS `costume`;
CREATE TABLE `costume` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `skin_name` varchar(255) NOT NULL,
  `character` varchar(255) NOT NULL,
  `ce_costume_condition` varchar(255) NOT NULL,
  `condition_value` varchar(255) NOT NULL,
  `illust_reference` varchar(255) NOT NULL,
  `portrait_reference` varchar(255) NOT NULL,
  `voice_appear` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for crew
-- ----------------------------
DROP TABLE IF EXISTS `crew`;
CREATE TABLE `crew` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `crew_introduction` varchar(255) NOT NULL,
  `crew_reference` varchar(255) NOT NULL,
  `crew_empty_reference` varchar(255) NOT NULL,
  `crew_group_photo_reference` varchar(255) NOT NULL,
  `crew_member` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for dialogue
-- ----------------------------
DROP TABLE IF EXISTS `dialogue`;
CREATE TABLE `dialogue` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `dialogue_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `number` int(10) unsigned NOT NULL,
  `spotlight` int(10) unsigned NOT NULL,
  `character` text NOT NULL,
  `animation` text NOT NULL,
  `highlighter` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for donation_contents
-- ----------------------------
DROP TABLE IF EXISTS `donation_contents`;
CREATE TABLE `donation_contents` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `image_reference` varchar(255) NOT NULL,
  `sponser_nickname` varchar(255) NOT NULL,
  `ce_common_type` varchar(255) NOT NULL,
  `item` varchar(255) NOT NULL,
  `item_value` int(10) unsigned NOT NULL,
  `system_message` varchar(255) NOT NULL,
  `sound` varchar(255) NOT NULL,
  `voice` varchar(255) NOT NULL,
  `animation` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for donation_message
-- ----------------------------
DROP TABLE IF EXISTS `donation_message`;
CREATE TABLE `donation_message` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `sponser_enum` varchar(255) NOT NULL,
  `ce_donation_trigger` varchar(255) NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for donation_sponser
-- ----------------------------
DROP TABLE IF EXISTS `donation_sponser`;
CREATE TABLE `donation_sponser` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `character_enum` varchar(255) NOT NULL,
  `sponser` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for donation_trigger
-- ----------------------------
DROP TABLE IF EXISTS `donation_trigger`;
CREATE TABLE `donation_trigger` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_donation_trigger` varchar(255) NOT NULL,
  `rate` double NOT NULL,
  `delay` double NOT NULL,
  `ce_streamer` varchar(255) NOT NULL,
  `bonus_trigger` text NOT NULL,
  `bonus_rate_plus` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for drop
-- ----------------------------
DROP TABLE IF EXISTS `drop`;
CREATE TABLE `drop` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `content_type` varchar(255) NOT NULL,
  `ce_common_type_drop` text NOT NULL,
  `icon` text NOT NULL,
  `drop` text NOT NULL,
  `rate` text NOT NULL,
  `value` text NOT NULL,
  `drop_next` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for emoticon
-- ----------------------------
DROP TABLE IF EXISTS `emoticon`;
CREATE TABLE `emoticon` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `emoticon_clear_reference` varchar(255) NOT NULL,
  `emoticon_white_reference` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for fate_card
-- ----------------------------
DROP TABLE IF EXISTS `fate_card`;
CREATE TABLE `fate_card` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_fate_card_grade` varchar(255) NOT NULL,
  `card_trigger` varchar(255) NOT NULL,
  `card_effect` text NOT NULL,
  `card_effect_value` text NOT NULL,
  `card_effect_duration` text NOT NULL,
  `main_grade_reference` varchar(255) NOT NULL,
  `slot_grade_reference` varchar(255) NOT NULL,
  `icon_reference` varchar(255) NOT NULL,
  `reward_grade_reference` varchar(255) NOT NULL,
  `image_reference` varchar(255) NOT NULL,
  `recipe` varchar(255) NOT NULL,
  `effect_reference` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha
-- ----------------------------
DROP TABLE IF EXISTS `gacha`;
CREATE TABLE `gacha` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `gacha_class` varchar(255) NOT NULL,
  `gacha_resource` varchar(255) NOT NULL,
  `slot_image` varchar(255) NOT NULL,
  `character_show` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_class
-- ----------------------------
DROP TABLE IF EXISTS `gacha_class`;
CREATE TABLE `gacha_class` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `prob` text NOT NULL,
  `pool_group` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_drop_a
-- ----------------------------
DROP TABLE IF EXISTS `gacha_drop_a`;
CREATE TABLE `gacha_drop_a` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `drop_character` varchar(255) NOT NULL,
  `pool_id` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_drop_b
-- ----------------------------
DROP TABLE IF EXISTS `gacha_drop_b`;
CREATE TABLE `gacha_drop_b` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `drop_character` varchar(255) NOT NULL,
  `pool_id` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_drop_c
-- ----------------------------
DROP TABLE IF EXISTS `gacha_drop_c`;
CREATE TABLE `gacha_drop_c` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `drop_character` varchar(255) NOT NULL,
  `pool_id` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_drop_d
-- ----------------------------
DROP TABLE IF EXISTS `gacha_drop_d`;
CREATE TABLE `gacha_drop_d` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `drop_character` varchar(255) NOT NULL,
  `pool_id` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_local
-- ----------------------------
DROP TABLE IF EXISTS `gacha_local`;
CREATE TABLE `gacha_local` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `kr` varchar(2047) NOT NULL,
  `en` varchar(2047) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_pool_group
-- ----------------------------
DROP TABLE IF EXISTS `gacha_pool_group`;
CREATE TABLE `gacha_pool_group` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `pool_condition` varchar(255) NOT NULL,
  `pool_id` text NOT NULL,
  `pool_id_rate` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_resource
-- ----------------------------
DROP TABLE IF EXISTS `gacha_resource`;
CREATE TABLE `gacha_resource` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `name_color` varchar(255) NOT NULL,
  `box_reference` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for gacha_stack_model
-- ----------------------------
DROP TABLE IF EXISTS `gacha_stack_model`;
CREATE TABLE `gacha_stack_model` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `a_prob_correction` double NOT NULL,
  `stack_correction_start` int(10) unsigned NOT NULL,
  `stack_correction_end` int(10) unsigned NOT NULL,
  `stack_100` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for guide
-- ----------------------------
DROP TABLE IF EXISTS `guide`;
CREATE TABLE `guide` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_guide_type` varchar(255) NOT NULL,
  `page` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for guide_page
-- ----------------------------
DROP TABLE IF EXISTS `guide_page`;
CREATE TABLE `guide_page` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `guide_resource_kr` varchar(255) NOT NULL,
  `guide_resource_en` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for item
-- ----------------------------
DROP TABLE IF EXISTS `item`;
CREATE TABLE `item` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_item` varchar(255) NOT NULL,
  `ce_item_sub` varchar(255) NOT NULL,
  `ce_item_grade` varchar(255) NOT NULL,
  `bundle_max` int(10) unsigned NOT NULL,
  `cost_value` int(10) unsigned NOT NULL,
  `grade_resource_thumbnail` varchar(255) NOT NULL,
  `grade_image_reference` varchar(255) NOT NULL,
  `grade_resource` varchar(255) NOT NULL,
  `image_reference` varchar(255) NOT NULL,
  `value` int(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for library_local
-- ----------------------------
DROP TABLE IF EXISTS `library_local`;
CREATE TABLE `library_local` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `kr` varchar(2047) NOT NULL,
  `en` varchar(2047) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for localization
-- ----------------------------
DROP TABLE IF EXISTS `localization`;
CREATE TABLE `localization` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `kr` varchar(2047) NOT NULL,
  `en` varchar(2047) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for map_abyss
-- ----------------------------
DROP TABLE IF EXISTS `map_abyss`;
CREATE TABLE `map_abyss` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `data` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`data`)),
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uniq_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for materials
-- ----------------------------
DROP TABLE IF EXISTS `materials`;
CREATE TABLE `materials` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_common_type_material` text NOT NULL,
  `material` text NOT NULL,
  `material_value` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for mention_tier_1
-- ----------------------------
DROP TABLE IF EXISTS `mention_tier_1`;
CREATE TABLE `mention_tier_1` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `character_appear` varchar(255) NOT NULL,
  `character_deadly` varchar(255) NOT NULL,
  `ally_kill` varchar(255) NOT NULL,
  `enemy_kill` varchar(255) NOT NULL,
  `exclusion` varchar(255) NOT NULL,
  `firstblood` varchar(255) NOT NULL,
  `resource_full` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for mention_tier_2
-- ----------------------------
DROP TABLE IF EXISTS `mention_tier_2`;
CREATE TABLE `mention_tier_2` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_mention_trigger` varchar(255) NOT NULL,
  `character_appear` varchar(255) NOT NULL,
  `voice_resource` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for mention_trigger
-- ----------------------------
DROP TABLE IF EXISTS `mention_trigger`;
CREATE TABLE `mention_trigger` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_mention_trigger` varchar(255) NOT NULL,
  `mention_rate` double NOT NULL,
  `mention_tier` int(10) unsigned NOT NULL,
  `mention_majority` int(10) unsigned NOT NULL,
  `viewable_both` tinyint(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for monster
-- ----------------------------
DROP TABLE IF EXISTS `monster`;
CREATE TABLE `monster` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `class` varchar(255) NOT NULL,
  `root` varchar(255) NOT NULL,
  `ce_character_property` varchar(255) NOT NULL,
  `subclass` varchar(255) NOT NULL,
  `tier` int(10) unsigned NOT NULL,
  `basestat` varchar(255) NOT NULL,
  `levelstat` varchar(255) NOT NULL,
  `active_radius` int(10) unsigned NOT NULL,
  `attack_radius` int(10) unsigned NOT NULL,
  `skill_basic` varchar(255) NOT NULL,
  `skill_active` varchar(255) NOT NULL,
  `resource` varchar(255) NOT NULL,
  `unit_size` double NOT NULL,
  `idle_ani` varchar(255) NOT NULL,
  `run_ani` varchar(255) NOT NULL,
  `hurt_ani` varchar(255) NOT NULL,
  `die_ani` varchar(255) NOT NULL,
  `revive_ani` varchar(255) NOT NULL,
  `destruction_ani` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for monster_pool
-- ----------------------------
DROP TABLE IF EXISTS `monster_pool`;
CREATE TABLE `monster_pool` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `monster` text NOT NULL,
  `weight` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `profile_resource` varchar(255) NOT NULL,
  `character_introduction` varchar(255) NOT NULL,
  `channel` varchar(255) NOT NULL,
  `habit` varchar(255) NOT NULL,
  `career` varchar(255) NOT NULL,
  `weakness` varchar(255) NOT NULL,
  `fandom` varchar(255) NOT NULL,
  `donation_reaction` varchar(255) NOT NULL,
  `crew` varchar(255) NOT NULL,
  `daily_unlock_condition` varchar(255) NOT NULL,
  `daily_life` varchar(255) NOT NULL,
  `real_name` varchar(255) NOT NULL,
  `gender` varchar(255) NOT NULL,
  `real_job` varchar(255) NOT NULL,
  `height` varchar(255) NOT NULL,
  `weight` varchar(255) NOT NULL,
  `favorite_thing` varchar(255) NOT NULL,
  `dislike_thing` varchar(255) NOT NULL,
  `personality` varchar(255) NOT NULL,
  `relationships` varchar(255) NOT NULL,
  `comment` varchar(255) NOT NULL,
  `viewers` int(10) unsigned NOT NULL,
  `illustrator` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for profile_resource
-- ----------------------------
DROP TABLE IF EXISTS `profile_resource`;
CREATE TABLE `profile_resource` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `library_thumbnail_reference` varchar(255) NOT NULL,
  `library_bg_reference` varchar(255) NOT NULL,
  `face_image_reference` text NOT NULL,
  `banner_image_reference` varchar(255) NOT NULL,
  `on_air_illust_reference` varchar(255) NOT NULL,
  `offline_illust_reference` varchar(255) NOT NULL,
  `color_code` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for recipe
-- ----------------------------
DROP TABLE IF EXISTS `recipe`;
CREATE TABLE `recipe` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_common_type_result` varchar(255) NOT NULL,
  `ce_craft_category` varchar(255) NOT NULL,
  `result` varchar(255) NOT NULL,
  `recipe_material` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for relationships
-- ----------------------------
DROP TABLE IF EXISTS `relationships`;
CREATE TABLE `relationships` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `friend` text NOT NULL,
  `friend_nickname` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for rule
-- ----------------------------
DROP TABLE IF EXISTS `rule`;
CREATE TABLE `rule` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `win_condition` varchar(255) NOT NULL,
  `elite_spawn_period` double NOT NULL,
  `elite_spawn_limit` int(10) unsigned NOT NULL,
  `minor_tile_generate_start` int(10) unsigned NOT NULL,
  `minor_tile_generate_max` int(10) unsigned NOT NULL,
  `minor_tile_addgen_time` double NOT NULL,
  `minor_spawn_limit_per_tile` int(10) unsigned NOT NULL,
  `minor_step_division` int(10) unsigned NOT NULL,
  `minor_spawn_period` double NOT NULL,
  `enhance_spawn_period` double NOT NULL,
  `enhance_rate` double NOT NULL,
  `enhance_period` double NOT NULL,
  `enhance_multiplier` double NOT NULL,
  `players_enhance_tiles` int(10) unsigned NOT NULL,
  `player_enhance_tiles_initial` int(10) unsigned NOT NULL,
  `neutral_enhance_tiles` int(10) unsigned NOT NULL,
  `neutral_enhance_tiles_initial` int(10) unsigned NOT NULL,
  `boss_spawn_period` double NOT NULL,
  `initial_boss_spawn` double NOT NULL,
  `chara_radius` int(10) unsigned NOT NULL,
  `chara_respawn_limit` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for signature_weapon
-- ----------------------------
DROP TABLE IF EXISTS `signature_weapon`;
CREATE TABLE `signature_weapon` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `base_atk` int(10) unsigned NOT NULL,
  `base_def` int(10) unsigned NOT NULL,
  `base_hp` int(10) unsigned NOT NULL,
  `ce_weapon_property` varchar(255) NOT NULL,
  `base_property` int(10) unsigned NOT NULL,
  `weapon_skill` varchar(255) NOT NULL,
  `weapon_growth_materials` varchar(255) NOT NULL,
  `weapon_level_stat` varchar(255) NOT NULL,
  `weapon_reference` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skill_data
-- ----------------------------
DROP TABLE IF EXISTS `skill_data`;
CREATE TABLE `skill_data` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `skill_type` varchar(255) NOT NULL,
  `skill_subtype` varchar(255) NOT NULL,
  `target_type` varchar(255) NOT NULL,
  `object_type` varchar(255) NOT NULL,
  `hit_radius` int(10) unsigned NOT NULL,
  `hit_center` varchar(255) NOT NULL,
  `animation` varchar(255) NOT NULL,
  `fragment_cost` int(10) unsigned NOT NULL,
  `skill_value` double NOT NULL,
  `range` int(10) unsigned NOT NULL,
  `duration` double NOT NULL,
  `cooldown` double NOT NULL,
  `resource_delay` double NOT NULL,
  `skill_resource` varchar(255) NOT NULL,
  `projectile_resource` varchar(255) NOT NULL,
  `hit_resource` varchar(255) NOT NULL,
  `hit_delay` double NOT NULL,
  `child` varchar(255) NOT NULL,
  `lag` double NOT NULL,
  `caused_effect` varchar(255) NOT NULL,
  `projectile_speed` double NOT NULL,
  `icon` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skilleffectproperties
-- ----------------------------
DROP TABLE IF EXISTS `skilleffectproperties`;
CREATE TABLE `skilleffectproperties` (
  `id` varchar(255) NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `coefficient` double NOT NULL,
  `neu_coefficient` double NOT NULL,
  `stagger_value` double NOT NULL,
  `unavoidable` int(10) unsigned NOT NULL,
  `unblockable` int(10) unsigned NOT NULL,
  `perstack` int(10) unsigned NOT NULL,
  `grade` double NOT NULL,
  `constvalue` double NOT NULL,
  `percentagevalue` double NOT NULL,
  `duration` double NOT NULL,
  `dur_independence` int(10) unsigned NOT NULL,
  `dur_perstack` double NOT NULL,
  `ce_enable_cond` varchar(255) NOT NULL,
  `cond_var` double NOT NULL,
  `apply_trigger` varchar(255) NOT NULL,
  `frequency` double NOT NULL,
  `maxstack` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skilleffects
-- ----------------------------
DROP TABLE IF EXISTS `skilleffects`;
CREATE TABLE `skilleffects` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_effect_category` varchar(255) NOT NULL,
  `icon_reference` varchar(255) NOT NULL,
  `leveladjust` varchar(255) NOT NULL,
  `resource` varchar(255) NOT NULL,
  `ce_remove_cond` varchar(255) NOT NULL,
  `lasts` int(10) unsigned NOT NULL,
  `ce_enable_cond` varchar(255) NOT NULL,
  `bind` int(10) unsigned NOT NULL,
  `proh_attack` int(10) unsigned NOT NULL,
  `proh_skills` int(10) unsigned NOT NULL,
  `invincibility` int(10) unsigned NOT NULL,
  `ce_stat` varchar(255) NOT NULL,
  `ce_stat_affected` varchar(255) NOT NULL,
  `ce_stat_subtype` varchar(255) NOT NULL,
  `ce_affect_type` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skillelements
-- ----------------------------
DROP TABLE IF EXISTS `skillelements`;
CREATE TABLE `skillelements` (
  `id` varchar(255) NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `radius` varchar(255) NOT NULL,
  `radius_center` varchar(255) NOT NULL,
  `ce_target` varchar(255) NOT NULL,
  `effect` text NOT NULL,
  `resource` text NOT NULL,
  `property` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skilllocalization
-- ----------------------------
DROP TABLE IF EXISTS `skilllocalization`;
CREATE TABLE `skilllocalization` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `kr` varchar(255) NOT NULL,
  `en` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skillpreset
-- ----------------------------
DROP TABLE IF EXISTS `skillpreset`;
CREATE TABLE `skillpreset` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_skill_subtype` varchar(255) NOT NULL,
  `icon_reference` varchar(255) NOT NULL,
  `animation` varchar(255) NOT NULL,
  `child` varchar(255) NOT NULL,
  `target` varchar(255) NOT NULL,
  `ce_target_condition` varchar(255) NOT NULL,
  `ce_order_type` varchar(255) NOT NULL,
  `add_target` int(10) unsigned NOT NULL,
  `trigger` varchar(255) NOT NULL,
  `range` int(10) unsigned NOT NULL,
  `cd` double NOT NULL,
  `startlag` double NOT NULL,
  `endlag` double NOT NULL,
  `ce_stat` varchar(255) NOT NULL,
  `skill_value` double NOT NULL,
  `ce_cost_type` varchar(255) NOT NULL,
  `cost_required` double NOT NULL,
  `cost_consumed` double NOT NULL,
  `ce_enable_cond` text NOT NULL,
  `cond_value` text NOT NULL,
  `element` text NOT NULL,
  `skill_resource` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skillradius
-- ----------------------------
DROP TABLE IF EXISTS `skillradius`;
CREATE TABLE `skillradius` (
  `id` varchar(255) NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `range` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skilltargettype
-- ----------------------------
DROP TABLE IF EXISTS `skilltargettype`;
CREATE TABLE `skilltargettype` (
  `id` varchar(255) NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_target_start` varchar(255) NOT NULL,
  `ce_target_end` varchar(255) NOT NULL,
  `ce_result_time` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for skilltriggers
-- ----------------------------
DROP TABLE IF EXISTS `skilltriggers`;
CREATE TABLE `skilltriggers` (
  `id` varchar(255) NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `value` double NOT NULL,
  `ce_skill_trigger` varchar(255) NOT NULL,
  `activate_chance` double NOT NULL,
  `priority_duration` double NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for sprite_management
-- ----------------------------
DROP TABLE IF EXISTS `sprite_management`;
CREATE TABLE `sprite_management` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `sprite_name` varchar(255) NOT NULL,
  `atlas_address` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for stargem_growth
-- ----------------------------
DROP TABLE IF EXISTS `stargem_growth`;
CREATE TABLE `stargem_growth` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `level` int(10) unsigned NOT NULL,
  `exp` int(10) unsigned NOT NULL,
  `exp_plus` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for stargem_resource
-- ----------------------------
DROP TABLE IF EXISTS `stargem_resource`;
CREATE TABLE `stargem_resource` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `image` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for start
-- ----------------------------
DROP TABLE IF EXISTS `start`;
CREATE TABLE `start` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_common_type` varchar(255) NOT NULL,
  `reward_id` varchar(255) NOT NULL,
  `value` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for story_stage
-- ----------------------------
DROP TABLE IF EXISTS `story_stage`;
CREATE TABLE `story_stage` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `stage_index` int(10) unsigned NOT NULL,
  `deck` varchar(255) NOT NULL,
  `abyss` varchar(255) NOT NULL,
  `reward` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for string_local
-- ----------------------------
DROP TABLE IF EXISTS `string_local`;
CREATE TABLE `string_local` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `kr` varchar(255) NOT NULL,
  `en` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tier
-- ----------------------------
DROP TABLE IF EXISTS `tier`;
CREATE TABLE `tier` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `icon_reference` varchar(255) NOT NULL,
  `rank_up_necessary` int(10) unsigned NOT NULL,
  `ai_pool` varchar(255) NOT NULL,
  `map_pool` varchar(255) NOT NULL,
  `viewers_final_min` int(10) unsigned NOT NULL,
  `viewers_final_max` int(10) unsigned NOT NULL,
  `win_point` int(10) unsigned NOT NULL,
  `win_reward` varchar(255) NOT NULL,
  `lose_point` int(10) NOT NULL,
  `lose_reward` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tutorial_event
-- ----------------------------
DROP TABLE IF EXISTS `tutorial_event`;
CREATE TABLE `tutorial_event` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_tutorial_event` varchar(255) NOT NULL,
  `group` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tutorial_event_step
-- ----------------------------
DROP TABLE IF EXISTS `tutorial_event_step`;
CREATE TABLE `tutorial_event_step` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_root_event` varchar(255) NOT NULL,
  `event_step_enum` varchar(255) NOT NULL,
  `advance_trigger` varchar(255) NOT NULL,
  `tutorial_group` varchar(255) NOT NULL,
  `ce_tutorial_resource_type` varchar(255) NOT NULL,
  `resource` varchar(255) NOT NULL,
  `resource_optional_string` varchar(255) NOT NULL,
  `resource_index` int(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tutorial_group
-- ----------------------------
DROP TABLE IF EXISTS `tutorial_group`;
CREATE TABLE `tutorial_group` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `next_group` varchar(255) NOT NULL,
  `group_reward` varchar(255) NOT NULL,
  `group_contents` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tutorial_trigger
-- ----------------------------
DROP TABLE IF EXISTS `tutorial_trigger`;
CREATE TABLE `tutorial_trigger` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `ce_tutorial_trigger` varchar(255) NOT NULL,
  `dialogue_id` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_common_production
-- ----------------------------
DROP TABLE IF EXISTS `vstory_common_production`;
CREATE TABLE `vstory_common_production` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `place` varchar(255) NOT NULL,
  `place_sprite` varchar(255) NOT NULL,
  `donation` varchar(255) NOT NULL,
  `donation_target` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_dialogue
-- ----------------------------
DROP TABLE IF EXISTS `vstory_dialogue`;
CREATE TABLE `vstory_dialogue` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `crew` varchar(255) NOT NULL,
  `dialogue` varchar(255) NOT NULL,
  `ce_dialogue_format` varchar(255) NOT NULL,
  `thumbnail` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_oracle_broadcast_pool
-- ----------------------------
DROP TABLE IF EXISTS `vstory_oracle_broadcast_pool`;
CREATE TABLE `vstory_oracle_broadcast_pool` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `character_enum_id` varchar(255) NOT NULL,
  `event` text NOT NULL,
  `regular` text NOT NULL,
  `irregular` text NOT NULL,
  `no_get` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_oracle_chat_pool
-- ----------------------------
DROP TABLE IF EXISTS `vstory_oracle_chat_pool`;
CREATE TABLE `vstory_oracle_chat_pool` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `chat` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_oracle_nickname_pool
-- ----------------------------
DROP TABLE IF EXISTS `vstory_oracle_nickname_pool`;
CREATE TABLE `vstory_oracle_nickname_pool` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `nickname` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_oracle_production
-- ----------------------------
DROP TABLE IF EXISTS `vstory_oracle_production`;
CREATE TABLE `vstory_oracle_production` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `sponsorable` tinyint(1) NOT NULL,
  `place` varchar(255) NOT NULL,
  `place_sprite` varchar(255) NOT NULL,
  `donation` varchar(255) NOT NULL,
  `on_air` tinyint(1) NOT NULL,
  `chat_cycle_min` double NOT NULL,
  `chat_cycle_max` double NOT NULL,
  `oracle_common_chat_pool` text NOT NULL,
  `oracle_nickname_pool` text NOT NULL,
  `special_chat` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_oracle_special_chat
-- ----------------------------
DROP TABLE IF EXISTS `vstory_oracle_special_chat`;
CREATE TABLE `vstory_oracle_special_chat` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `nickname` text NOT NULL,
  `special_chat` text NOT NULL,
  `special_chat_format` text NOT NULL,
  `optional_string` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for vstory_timeline
-- ----------------------------
DROP TABLE IF EXISTS `vstory_timeline`;
CREATE TABLE `vstory_timeline` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `timeline_track` varchar(255) NOT NULL,
  `ce_story_type` varchar(255) NOT NULL,
  `optional_int` int(10) NOT NULL,
  `story_reward` varchar(255) NOT NULL,
  `banner_background` varchar(255) NOT NULL,
  `dialogue` text NOT NULL,
  `production` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for weapon_level_stat
-- ----------------------------
DROP TABLE IF EXISTS `weapon_level_stat`;
CREATE TABLE `weapon_level_stat` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `level_atk` int(10) unsigned NOT NULL,
  `level_def` int(10) unsigned NOT NULL,
  `level_hp` int(10) unsigned NOT NULL,
  `level_property` int(10) unsigned NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for world_map_local
-- ----------------------------
DROP TABLE IF EXISTS `world_map_local`;
CREATE TABLE `world_map_local` (
  `id` int(10) unsigned NOT NULL,
  `enum_id` varchar(255) NOT NULL,
  `kr` varchar(255) NOT NULL,
  `en` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_enum_id` (`enum_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
