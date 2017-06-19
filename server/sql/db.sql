SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

CREATE DATABASE IF NOT EXISTS `netburn` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `netburn`;

CREATE TABLE IF NOT EXISTS `user` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(128) NOT NULL DEFAULT '',
	`email` varchar(128) NOT NULL DEFAULT '',
    `age` int(3) NOT NULL DEFAULT 0,
    `memo` varchar(254) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;