CREATE TABLE `offer_company` (
    `offer_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `client_id` int(10) unsigned NOT NULL,
    `country` char(4) NOT NULL,
    `image` varchar(100) NOT NULL,
    `image_width` int(11) NOT NULL DEFAULT 0,
    `image_height` int(11) NOT NULL DEFAULT 0,
    `text_locale` tinytext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
    `validity_text_locale` tinytext NOT NULL,
    `position` int(10) NOT NULL,
    `valid_from` datetime NOT NULL DEFAULT '1970-01-01 00:00:00',
    `show_from` datetime DEFAULT '1970-01-01 00:00:00',
    `valid_to` datetime NOT NULL DEFAULT '2100-01-01 00:00:00',
    `flag` int(10) unsigned NOT NULL DEFAULT 1 COMMENT '0 - Hide\n1 - Show',
    `page_count` int(10) unsigned NOT NULL DEFAULT 0,
    `store_url` varchar(100) DEFAULT '',
    `store_url_title` varchar(100) NOT NULL DEFAULT 0,
    `offer_home` int(11) NOT NULL DEFAULT 1,
    PRIMARY KEY (`offer_id`),
    UNIQUE KEY `offer_id_UNIQUE` (`offer_id`),
    KEY `valid_from_INDEX` (`valid_from`),
    KEY `valid_to_INDEX` (`valid_to`)
    ) ENGINE=InnoDB AUTO_INCREMENT=307469 DEFAULT CHARSET=utf8mb4
    COLLATE=utf8mb4_unicode_ci