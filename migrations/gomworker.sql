-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: database:3306
-- Generation Time: Jul 23, 2018 at 05:55 AM
-- Server version: 5.7.19
-- PHP Version: 7.2.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gomworker`
--

-- --------------------------------------------------------

--
-- Table structure for table `workers`
--

CREATE TABLE `workers` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` enum('started','running','stopped','asked_to_stop') COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'started',
  `usage_cpu` double(8,2) UNSIGNED NOT NULL DEFAULT '0.00',
  `usage_memory` double(8,2) UNSIGNED NOT NULL DEFAULT '0.00',
  `queue` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default',
  `once` tinyint(1) NOT NULL DEFAULT '0',
  `delay` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `force` tinyint(1) NOT NULL DEFAULT '0',
  `memory` int(10) UNSIGNED NOT NULL DEFAULT '128',
  `sleep` int(10) UNSIGNED NOT NULL DEFAULT '3',
  `timeout` int(10) UNSIGNED NOT NULL DEFAULT '60',
  `tries` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `started_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `workers`
--
ALTER TABLE `workers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `workers_name_unique` (`name`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `workers`
--
ALTER TABLE `workers`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
