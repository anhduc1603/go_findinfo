-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 13, 2025 at 08:21 AM
-- Server version: 8.0.30
-- PHP Version: 8.1.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `info_find`
--

-- --------------------------------------------------------

--
-- Table structure for table `response_history_info`
--

CREATE TABLE `response_history_info` (
  `id` bigint NOT NULL,
  `info` varchar(255) DEFAULT NULL,
  `content` varchar(255) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `userid` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `response_history_info`
--

INSERT INTO `response_history_info` (`id`, `info`, `content`, `status`, `userid`, `created_at`, `updated_at`) VALUES
(1, 'Info 1', 'Content 1', 1, 101, '2025-03-10 07:16:56', '2025-03-10 07:16:56'),
(2, 'Info 2', 'Content 2', 0, 102, '2025-03-10 07:16:56', '2025-03-10 07:16:56'),
(3, 'Info 3', 'Content 3', 1, 103, '2025-03-10 07:16:56', '2025-03-10 07:16:56'),
(4, 'Info 4', 'Content 4', 1, 104, '2025-03-10 07:16:56', '2025-03-10 07:16:56'),
(5, 'Info 5', 'Content 5', 0, 105, '2025-03-10 07:16:56', '2025-03-10 07:16:56');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint UNSIGNED NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` text NOT NULL,
  `role` varchar(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` int DEFAULT NULL
) ;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `role`, `created_at`, `updated_at`, `status`) VALUES
(1, 'admin', '$2a$10$D.xurCPhfb4wzS.1XtOj6OSMs0yQg/O9FQwArjJkeTrPyvyZ4m50G', 'admin', '2025-03-12 04:18:05', '2025-03-13 02:15:09', 1),
(2, 'user', '$2a$10$D.xurCPhfb4wzS.1XtOj6OSMs0yQg/O9FQwArjJkeTrPyvyZ4m50G', 'user', '2025-03-12 04:18:05', '2025-03-13 02:15:09', 1),
(4, 'ducha', '$2a$10$UKUkk03Ui/TIor/SIYwtY.YyOzi2BxBl42cUpeve9zd/qiBX7U6Je', 'admin', '2025-03-12 08:14:57', '2025-03-13 02:15:09', 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `response_history_info`
--
ALTER TABLE `response_history_info`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `id` (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `response_history_info`
--
ALTER TABLE `response_history_info`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
