-- phpMyAdmin SQL Dump
-- version 4.8.4
-- https://www.phpmyadmin.net/
--
-- Host: db
-- Generation Time: Jan 30, 2019 at 08:57 AM
-- Server version: 5.7.24
-- PHP Version: 7.2.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `goheroes`
--

-- --------------------------------------------------------

--
-- Table structure for table `authentication_tokens`
--

CREATE TABLE `authentication_tokens` (
  `id` int(10) UNSIGNED NOT NULL,
  `token` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `additional` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `expire_at` datetime NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `topic_id` int(10) UNSIGNED NOT NULL,
  `comment` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `downloads`
--

CREATE TABLE `downloads` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `forums`
--

CREATE TABLE `forums` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `friend_requests`
--

CREATE TABLE `friend_requests` (
  `id` int(10) UNSIGNED NOT NULL,
  `sender` int(10) UNSIGNED NOT NULL,
  `receiver` int(10) UNSIGNED NOT NULL,
  `status` enum('pending','accepted','declined') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'pending',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `games`
--

CREATE TABLE `games` (
  `gid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `game_ip` varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
  `game_port` int(11) NOT NULL,
  `game_version` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status_join` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status_mapname` varchar(16) COLLATE utf8mb4_unicode_ci NOT NULL,
  `players_connected` int(11) NOT NULL,
  `players_joining` int(11) NOT NULL,
  `players_max` int(11) NOT NULL DEFAULT '32',
  `team_1` int(11) NOT NULL DEFAULT '0',
  `team_2` int(11) NOT NULL DEFAULT '0',
  `team_distribution` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `game_heroes`
--

CREATE TABLE `game_heroes` (
  `heroID` int(11) NOT NULL,
  `heroName` varchar(50) NOT NULL,
  `online` tinyint(1) NOT NULL DEFAULT '0',
  `ip_address` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `user_id` int(11) NOT NULL,
  `hero_stats` text NOT NULL,
  `game_token` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `game_heroes`
--

INSERT INTO `game_heroes` (`heroID`, `heroName`, `online`, `ip_address`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `hero_stats`, `game_token`) VALUES
(2, '1234', 0, NULL, NULL, NULL, NULL, 2, '{\"c_ft\":\"\",\"c_team\":\"\",\"c_hrc\":\"\",\"c_hrs\":\"\",\"c_skc\":\"\",\"c_ltp\":\"9787.0000\",\"c_ltm\":\"9787.0000\",\"c_fhrs\":\"\",\"c_slm\":\"0.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"\",\"c_wallet_hero\":\"14.0000\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"\",\"level\":\"\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{\"0\":\"6000.0000\",\"1\":\"6000.0000\",\"2\":\"6000.0000\"},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{\"0\":\"0.0000\",\"1\":\"0.0000\",\"2\":\"0.0000\"},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[\"10\",\"1330\"],\"c_emo\":[],\"c_eqp\":[\"3167\",\"0\",\"0\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[\"2003\",\"2026\",\"2027\",\"2028\",\"2031\",\"2032\",\"2033\",\"2046\",\"2047\",\"2048\",\"2055\",\"2056\",\"2057\",\"2091\",\"2141\",\"2142\",\"2149\",\"2153\",\"2157\",\"2158\",\"2160\",\"2162\",\"2164\",\"2166\",\"2168\",\"2170\",\"2171\",\"2174\"],\"ige\":{}}', '1234'),
(3, 'bianca', 0, '172.19.0.1', '2019-01-28 02:22:49', '2019-01-28 02:22:49', NULL, 12, '{\"c_ft\":\"0\",\"c_team\":\"1\",\"c_hrc\":\"1\",\"c_hrs\":\"122\",\"c_skc\":\"3\",\"c_ltp\":\"9787.0000\",\"c_ltm\":\"9337.0000\",\"c_fhrs\":\"\",\"c_slm\":\"2.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[],\"c_emo\":[],\"c_eqp\":[\"3184\",\"0\",\"0\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', '12'),
(4, 'jovempan', 0, '172.19.0.1', '2019-01-28 05:05:48', '2019-01-28 05:05:48', NULL, 12, '{\"c_ft\":\"130\",\"c_team\":\"1\",\"c_hrc\":\"2\",\"c_hrs\":\"120\",\"c_skc\":\"8\",\"c_ltp\":\"9789.0000\",\"c_ltm\":\"9337.0000\",\"c_fhrs\":\"\",\"c_slm\":\"2.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[],\"c_emo\":[],\"c_eqp\":[\"0\",\"3166\",\"0\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', NULL),
(5, 'celular', 0, '172.19.0.1', '2019-01-28 05:23:24', '2019-01-28 05:23:24', NULL, 12, '{\"c_ft\":\"102\",\"c_team\":\"2\",\"c_hrc\":\"2\",\"c_hrs\":\"0\",\"c_skc\":\"7\",\"c_ltp\":\"9788.0000\",\"c_ltm\":\"9337.0000\",\"c_fhrs\":\"\",\"c_slm\":\"2.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"0\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[\"1355\"],\"c_emo\":[],\"c_eqp\":[\"3172\",\"3196\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', NULL),
(6, 'coroahehe', 0, '172.19.0.1', '2019-01-30 00:25:41', '2019-01-30 00:25:41', NULL, 13, '{\"c_ft\":\"0\",\"c_team\":\"2\",\"c_hrc\":\"5\",\"c_hrs\":\"82\",\"c_skc\":\"2\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9790.0000\",\"c_fhrs\":\"\",\"c_slm\":\"12.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"1\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"0.0000\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[\"1264\",\"1355\"],\"c_emo\":[],\"c_eqp\":[\"3179\",\"3164\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', 'lgopt'),
(7, '5656565', 0, '172.19.0.1', '2019-01-30 00:55:18', '2019-01-30 00:55:18', NULL, 13, '{\"c_ft\":\"132\",\"c_team\":\"1\",\"c_hrc\":\"1\",\"c_hrs\":\"0\",\"c_skc\":\"7\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9790.0000\",\"c_fhrs\":\"\",\"c_slm\":\"0.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"1.0000\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{\"0\":\"6000.0000\"},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[],\"c_emo\":[\"5000\",\"5007\",\"5016\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_eqp\":[\"3004\",\"3010\",\"2028\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', 'lgopt'),
(8, 'niuijh', 0, '172.19.0.1', '2019-01-30 00:55:44', '2019-01-30 00:55:44', NULL, 13, '{\"c_ft\":\"130\",\"c_team\":\"1\",\"c_hrc\":\"4\",\"c_hrs\":\"122\",\"c_skc\":\"8\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9337.0000\",\"c_fhrs\":\"\",\"c_slm\":\"2.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"0\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{\"0\":\"6000.0000\"},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[],\"c_emo\":[\"5000\",\"5007\",\"5016\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_eqp\":[\"3002\",\"3014\",\"2141\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', 'lgopt'),
(9, 'consegui', 0, '172.19.0.1', '2019-01-30 05:05:51', '2019-01-30 05:05:51', NULL, 13, '{\"c_ft\":\"106\",\"c_team\":\"2\",\"c_hrc\":\"3\",\"c_hrs\":\"83\",\"c_skc\":\"7\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9790.0000\",\"c_fhrs\":\"\",\"c_slm\":\"1.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"1.0000\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[\"1341\"],\"c_emo\":[],\"c_eqp\":[\"3197\",\"3161\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', NULL),
(10, 'conqueer', 0, '172.19.0.1', '2019-01-30 05:07:42', '2019-01-30 05:07:42', NULL, 10, '{\"c_ft\":\"130\",\"c_team\":\"1\",\"c_hrc\":\"4\",\"c_hrs\":\"124\",\"c_skc\":\"8\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9790.0000\",\"c_fhrs\":\"\",\"c_slm\":\"0.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[\"10\",\"1347\"],\"c_emo\":[],\"c_eqp\":[\"3200\",\"3191\",\"0\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', '4321'),
(11, 'chora', 0, NULL, NULL, NULL, NULL, 10, '{\"c_ft\":\"130\",\"c_team\":\"1\",\"c_hrc\":\"4\",\"c_hrs\":\"124\",\"c_skc\":\"8\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9790.0000\",\"c_fhrs\":\"\",\"c_slm\":\"0.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[\"10\",\"1347\"],\"c_emo\":[],\"c_eqp\":[\"3200\",\"3191\",\"0\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', '4321'),
(12, 'TERRIBLE', 0, '172.19.0.1', '2019-01-30 08:53:29', '2019-01-30 08:53:29', NULL, 16, '{\"c_ft\":\"109\",\"c_team\":\"2\",\"c_hrc\":\"5\",\"c_hrs\":\"85\",\"c_skc\":\"8\",\"c_ltp\":\"9790.0000\",\"c_ltm\":\"9337.0000\",\"c_fhrs\":\"\",\"c_slm\":\"2.0000\",\"cdm\":\"\",\"edm\":\"\",\"c_kit\":\"2\",\"c_wallet_hero\":\"\",\"c_wallet_valor\":\"\",\"games\":\"\",\"elo\":\"1000\",\"level\":\"1\",\"xp\":\"\",\"ct\":\"\",\"ki\":\"\",\"dt\":\"\",\"su\":\"\",\"win\":\"\",\"los\":\"\",\"fi\":\"\",\"hi\":\"\",\"rs\":\"\",\"ts\":\"\",\"ss\":\"\",\"cs\":\"\",\"prs\":\"\",\"ppt\":\"\",\"c_tut\":\"\",\"awybt\":\"\",\"dmc\":\"\",\"gsco\":\"\",\"expts\":\"\",\"bnspt\":\"\",\"aw\":{},\"mid\":{},\"c_mid\":{},\"c_cmid\":{},\"m0c\":{},\"m1c\":{},\"m2c\":{},\"c_wmid\":{},\"startLVL\":\"\",\"roundXP\":\"\",\"roundBXP\":\"\",\"roundVP\":\"\",\"roundBVP\":\"\",\"roundHP\":\"\",\"roundPP\":\"\",\"roundBPP\":\"\",\"totalPP\":\"\",\"cpc\":\"\",\"cpa\":\"\",\"cpd\":\"\",\"rc\":\"\",\"ks\":\"\",\"ds\":\"\",\"ft_rs\":{},\"ft_ki\":{},\"ft_dt\":{},\"ft_win\":{},\"ft_los\":{},\"fc_rs\":{},\"fc_ki\":{},\"fc_dt\":{},\"fc_win\":{},\"fc_los\":{},\"m_ct\":{},\"m_win\":{},\"m_los\":{},\"tv\":{},\"kv\":{},\"dfv\":{},\"kvr\":{},\"dstrv\":{},\"div\":{},\"tw\":{},\"twk\":{},\"kw\":{},\"dfw\":{},\"sw\":{},\"hw\":{},\"dww\":{},\"kk\":{},\"kkb\":{},\"ka\":\"\",\"he\":\"\",\"drka\":\"\",\"c_apr\":[],\"c_emo\":[],\"c_eqp\":[\"3184\",\"0\",\"0\",\"3155\",\"0\",\"0\",\"0\",\"0\",\"0\",\"0\"],\"c_items\":[],\"ige\":{}}', '');

-- --------------------------------------------------------

--
-- Table structure for table `game_player_regions`
--

CREATE TABLE `game_player_regions` (
  `userid` int(11) NOT NULL,
  `region` varchar(2) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `game_player_server_preferences`
--

CREATE TABLE `game_player_server_preferences` (
  `userid` int(11) NOT NULL,
  `gid` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `game_stats`
--

CREATE TABLE `game_stats` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `heroID` int(10) UNSIGNED NOT NULL,
  `statsKey` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `statsValue` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `password_resets`
--

CREATE TABLE `password_resets` (
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `token` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `password_resets`
--

INSERT INTO `password_resets` (`email`, `token`, `created_at`) VALUES
('plsdonthackme@gmail.com', '$2y$10$STTxd6g2mVuqVxLXOyHGQOj97pF7w7LT.9/F.lP3qgd37AKL6PnJy', '2017-12-10 02:56:49');

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` int(10) UNSIGNED NOT NULL,
  `slug` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `slug`, `description`) VALUES
(1, 'game.createhero', NULL),
(2, 'game.unlimitedheroes', NULL),
(3, 'game.login', NULL),
(4, 'game.matchmake', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `permission_role`
--

CREATE TABLE `permission_role` (
  `permission_id` int(10) UNSIGNED NOT NULL,
  `role_id` int(10) UNSIGNED NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `permission_role`
--

INSERT INTO `permission_role` (`permission_id`, `role_id`) VALUES
(1, 1),
(2, 1),
(3, 1),
(4, 1);

-- --------------------------------------------------------

--
-- Table structure for table `players`
--

CREATE TABLE `players` (
  `user_id` int(11) NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(50) NOT NULL,
  `game_token` varchar(50) DEFAULT NULL,
  `prefer_server` int(11) DEFAULT NULL,
  `selected_heroID` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `players`
--

INSERT INTO `players` (`user_id`, `username`, `password`, `game_token`, `prefer_server`, `selected_heroID`) VALUES
(4, 'diego', 'diego', 'diego', NULL, NULL),
(6, 'aldo', 'aldo', 'aldo', NULL, NULL),
(7, 'aldo2', 'banana', 'banana', NULL, NULL),
(8, 'bfhgamer', 'japa', 'japa', NULL, NULL),
(9, 'SomeUser', '1234', '1234', NULL, 9),
(11, 'OtherUser', '123', '123', NULL, 11);

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int(10) UNSIGNED NOT NULL,
  `title` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `slug` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `title`, `slug`) VALUES
(1, 'administrator', 'administrator\r\n');

-- --------------------------------------------------------

--
-- Table structure for table `role_user`
--

CREATE TABLE `role_user` (
  `user_id` int(10) UNSIGNED NOT NULL,
  `role_id` int(10) UNSIGNED NOT NULL,
  `expire_at` datetime DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `role_user`
--

INSERT INTO `role_user` (`user_id`, `role_id`, `expire_at`) VALUES
(2, 1, NULL),
(3, 1, NULL),
(5, 1, NULL),
(6, 1, NULL),
(7, 1, NULL),
(8, 1, NULL),
(9, 1, NULL),
(10, 1, NULL),
(11, 1, NULL),
(12, 1, NULL),
(13, 1, NULL),
(14, 1, NULL),
(15, 1, NULL),
(16, 1, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `servers`
--

CREATE TABLE `servers` (
  `server_id` int(11) NOT NULL,
  `soldier_name` varchar(50) DEFAULT NULL,
  `account_username` varchar(50) DEFAULT NULL,
  `account_password` varchar(50) DEFAULT NULL,
  `api_key` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `servers`
--

INSERT INTO `servers` (`server_id`, `soldier_name`, `account_username`, `account_password`, `api_key`) VALUES
(1, 'MargeSimpson', 'MargeSimpson', 'MargeSimpson', 'MargeSimpson');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `username` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `birthday` date NOT NULL,
  `language` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'enUS',
  `country` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `password` varchar(191) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remember_token` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `notifications` mediumtext COLLATE utf8mb4_unicode_ci,
  `ip_address` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `game_token` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `email`, `birthday`, `language`, `country`, `password`, `remember_token`, `created_at`, `updated_at`, `notifications`, `ip_address`, `game_token`) VALUES
(1, 'MargeSimpson', 'game-server@gmail.com', '2017-12-06', 'enUS', 'Brazil', 'MargeSimpson', NULL, NULL, NULL, NULL, NULL, 'MargeSimpson'),
(2, 'lok', 'eeee@gmail.com', '1996-07-31', 'enUS', '', '1234', NULL, '2017-12-10 02:37:43', '2017-12-10 02:37:43', NULL, NULL, '1234'),
(3, 'Kiop', '1234@gmail.com', '1997-07-31', 'enUS', '', '$2y$10$yqc94mokmEtovgK1oXjoOeAN2/14iacs9BKDu.qokd8NWN4GoeGM2', NULL, '2017-12-10 02:59:21', '2017-12-10 02:59:21', NULL, NULL, '123'),
(4, 'diego', 'syn@gmail.com', '2018-01-27', 'fr', 'FR', 'diego', NULL, '2018-01-27 07:19:31', '2018-01-27 07:19:31', NULL, NULL, '434334'),
(7, 'aldo2', 'aldo2@email.com', '1980-07-21', 'enUS', '', '$2y$10$bQN2f08GJCMjukkpUn.EneU3GJ6PG7Bl11pJ2hSIOxZFhhsW57ynK', NULL, '2019-01-06 04:48:35', '2019-01-06 04:52:30', NULL, '26.36.195.228', 'banana'),
(5, 'caio', '22232@gmail.com', '1997-07-31', 'enUS', '', '$2y$10$M9RiC3XlhebivlXRPXeccupEZ2aajmfS4UMwiPICLeNkhsCYdI5Qe', NULL, '2018-01-30 03:11:10', '2018-01-30 04:52:54', '{\"news\":false}', '127.0.0.1', '6CUMvXUu6HNBajzR3wEtLwzfF4EgIrCJ'),
(6, 'aldo', 'aldo@email.com', '1980-07-21', 'enUS', '', '$2y$10$oLU58YbnAU0Rz7I7niVFZu4.CjPLJDCsTcBJB//yOGMzQo7CheNSq', NULL, '2019-01-06 03:11:42', '2019-01-06 03:11:42', NULL, NULL, 'aldo'),
(8, 'bfhgamer', 'bfhgamer@email.com', '1998-10-12', 'enUS', '', '$2y$10$E5jEGPKiNSC9r8Ih.Z.Owe/GnwCBBTurmji6mbuJxc2YI.QvSTJm.', NULL, '2019-01-06 04:50:56', '2019-01-06 04:50:56', NULL, NULL, 'japa'),
(9, 'japa', 'japa@email.com', '1983-03-21', 'enUS', '', '$2y$10$mlHm0kMrQTqx0Q5M5HpsyOkk5km53ZXJXi0lLkAAovg9AMn0AYBOa', 'ctjfri8us6gzRifWzV7gPDJHdusZTeeTzmp04ylJp13ZaE7lyXljL6UmIpZP', '2019-01-20 06:34:08', '2019-01-20 06:34:08', NULL, NULL, NULL),
(10, 'dev', 'dev@email.com', '1992-08-21', 'enUS', '', '$2y$10$vW783q3GkTPvECEji546DuU5l.6uwPaq28TIxn..ai71qxLga7YAi', NULL, '2019-01-20 06:37:03', '2019-01-20 06:37:03', NULL, NULL, NULL),
(11, 'teste', 'teste@email.com', '1992-01-21', 'enUS', '', '$2y$10$/x58zz6nYY9ONyyDQVLYv.SiUJbRs6phBzgh3s9kXoTF.VMtHQ/Ce', 'UTku5Y6PZKhr6UIdjMmymNPyZqX5wSE3EBvjjFA21Xt2kuCaauq3kGHLBKyd', '2019-01-27 02:46:16', '2019-01-27 03:18:15', NULL, '::1', 'jaGCFcGOHh1yVl4rDZ3Zy11S8j0E6ZKw'),
(12, 'breno', 'breno@email.com', '1991-01-21', 'enUS', '', '$2y$10$5Z8df4jFbNwUxSmFLjDl6.Jg0FJgkc4FPjPswIjNHY1dWB01nfFgu', 'eCunpZ8a3robWUqWIVl1EPcQO9f8531xSZOfRCyxQrqqQXQeWjdbG7e5WFeF', '2019-01-28 02:08:27', '2019-01-28 02:08:27', NULL, NULL, NULL),
(13, 'lgopt', 'lgopt@email.com', '1991-01-21', 'enUS', '', '$2y$10$F6klXNC1NoKMJkFosuN4DejT7P45U3y/gYbf5LJ92kdpSs22DeNuW', 'HCP4tK9GECTUVi7yWu3JxZWjCaF4eBtAzb6lh8IzkO1xCGabCZ6FDvqTCkK0', '2019-01-30 00:25:15', '2019-01-30 00:25:15', NULL, NULL, NULL),
(14, 'sorriso', 'sorriso@email.com', '1991-01-21', 'enUS', '', '$2y$10$wxaPBHdZnHaJn63BqTQEYuv3v2DLFgYHG2ehb7lGUnP3BN3X3cnK6', NULL, '2019-01-30 05:07:27', '2019-01-30 05:07:27', NULL, NULL, NULL),
(15, 'lgopt1234', 'lgopt1234@email.com', '1991-01-21', 'enUS', '', '$2y$10$XcJroP8znw8UZ4rPiR3QCObT04qwQA7g2nXtRD9kME6DnpWfxnbI6', NULL, '2019-01-30 08:07:35', '2019-01-30 08:07:35', NULL, NULL, NULL),
(16, 'quit', 'quit@email.com', '1991-01-21', 'enUS', '', '$2y$10$7CcCBPy12VeyCc4NM8/V5.w67gbsi8ua2cG2IAC0nc6loPPRm44xu', 'c1otZ4iTtCzaJATlvchPl2i0iXyUIjRfslTg7KKf0WOJfDZuxKaZOTdOZZzK', '2019-01-30 08:52:52', '2019-01-30 09:34:13', NULL, '172.19.0.1', 'o2y2mNh9LEzEwn75CFI6qacDzZ84nKnQ');

-- --------------------------------------------------------

--
-- Table structure for table `user_friends`
--

CREATE TABLE `user_friends` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(10) UNSIGNED NOT NULL,
  `friend_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `authentication_tokens`
--
ALTER TABLE `authentication_tokens`
  ADD PRIMARY KEY (`id`),
  ADD KEY `authentication_tokens_user_id_foreign` (`user_id`);

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `comments_user_id_foreign` (`user_id`),
  ADD KEY `comments_topic_id_foreign` (`topic_id`);

--
-- Indexes for table `downloads`
--
ALTER TABLE `downloads`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `forums`
--
ALTER TABLE `forums`
  ADD PRIMARY KEY (`id`),
  ADD KEY `forums_user_id_foreign` (`user_id`);

--
-- Indexes for table `friend_requests`
--
ALTER TABLE `friend_requests`
  ADD PRIMARY KEY (`id`),
  ADD KEY `friend_requests_sender_foreign` (`sender`),
  ADD KEY `friend_requests_receiver_foreign` (`receiver`);

--
-- Indexes for table `games`
--
ALTER TABLE `games`
  ADD PRIMARY KEY (`gid`);

--
-- Indexes for table `game_heroes`
--
ALTER TABLE `game_heroes`
  ADD PRIMARY KEY (`heroID`);

--
-- Indexes for table `game_player_regions`
--
ALTER TABLE `game_player_regions`
  ADD PRIMARY KEY (`userid`);

--
-- Indexes for table `game_player_server_preferences`
--
ALTER TABLE `game_player_server_preferences`
  ADD PRIMARY KEY (`userid`,`gid`);

--
-- Indexes for table `game_stats`
--
ALTER TABLE `game_stats`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `game_stats_user_id_heroid_statskey_unique` (`user_id`,`heroID`,`statsKey`);

--
-- Indexes for table `password_resets`
--
ALTER TABLE `password_resets`
  ADD KEY `password_resets_email_index` (`email`);

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `permission_role`
--
ALTER TABLE `permission_role`
  ADD KEY `permission_role_permission_id_foreign` (`permission_id`),
  ADD KEY `permission_role_role_id_foreign` (`role_id`);

--
-- Indexes for table `players`
--
ALTER TABLE `players`
  ADD PRIMARY KEY (`user_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `role_user`
--
ALTER TABLE `role_user`
  ADD KEY `role_user_user_id_foreign` (`user_id`),
  ADD KEY `role_user_role_id_foreign` (`role_id`);

--
-- Indexes for table `servers`
--
ALTER TABLE `servers`
  ADD PRIMARY KEY (`server_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `users_username_unique` (`username`),
  ADD UNIQUE KEY `users_email_unique` (`email`),
  ADD UNIQUE KEY `users_game_token_unique` (`game_token`);

--
-- Indexes for table `user_friends`
--
ALTER TABLE `user_friends`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_friends_user_id_foreign` (`user_id`),
  ADD KEY `user_friends_friend_id_foreign` (`friend_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `authentication_tokens`
--
ALTER TABLE `authentication_tokens`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `downloads`
--
ALTER TABLE `downloads`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `forums`
--
ALTER TABLE `forums`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `friend_requests`
--
ALTER TABLE `friend_requests`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `game_heroes`
--
ALTER TABLE `game_heroes`
  MODIFY `heroID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT for table `game_stats`
--
ALTER TABLE `game_stats`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=119;

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `players`
--
ALTER TABLE `players`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `servers`
--
ALTER TABLE `servers`
  MODIFY `server_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `user_friends`
--
ALTER TABLE `user_friends`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
