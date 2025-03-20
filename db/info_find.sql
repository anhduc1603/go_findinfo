-- phpMyAdmin SQL Dump
-- version 5.2.2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 20, 2025 at 10:02 AM
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
-- Table structure for table `google_tokens`
--

CREATE TABLE `google_tokens` (
  `id` int NOT NULL,
  `userid` bigint UNSIGNED NOT NULL,
  `google_access_token` text NOT NULL,
  `google_refresh_token` text,
  `token_expiry` datetime DEFAULT NULL,
  `user_cookie` varchar(2000) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `google_tokens`
--

INSERT INTO `google_tokens` (`id`, `userid`, `google_access_token`, `google_refresh_token`, `token_expiry`, `user_cookie`, `created_at`, `updated_at`) VALUES
(1, 19, 'ya29.a0AeXRPp4eCRwIXXB9MDX6exNWJdrdSxNoGh0nEwa1l_CYxKhONtP5xs0x_uOKqxH4eUttX4Gs3o6LPi6UTye20RljzCT1fXm2-M4hi4kJ70HYOOo2zlDD8sYKT1hUCOPZd0qM-qAiptQ5G6jzi8p_s7M37rHWqmad5uLyIbLQDAaCgYKAaUSARISFQHGX2MiAa0UyNTMEQxhNC9ZR3aYCg0177', '', '2025-03-20 15:33:58', '', '2025-03-20 07:00:34', '2025-03-20 07:34:01');

-- --------------------------------------------------------

--
-- Table structure for table `response_history_info`
--

CREATE TABLE `response_history_info` (
  `id` bigint NOT NULL,
  `info` varchar(255) DEFAULT NULL,
  `content` varchar(2000) DEFAULT NULL,
  `status` int DEFAULT NULL,
  `userid` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `response_history_info`
--

INSERT INTO `response_history_info` (`id`, `info`, `content`, `status`, `userid`, `created_at`, `updated_at`) VALUES
(37, 'Info 1', 'Content 1', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(38, 'Info 2', 'Sample content 2', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(39, 'Info 3', 'Example content 3', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(40, 'Info 4', 'Content data 4', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(41, 'Info 5', 'Another content 5', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(42, 'Info 6', 'Random content 6', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(43, 'Info 7', 'Test content 7', 0, 4, '2025-03-18 02:53:07', '2025-03-19 07:55:41'),
(44, 'Info 8', 'Demo content 8', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(45, 'Info 9', 'Content sample 9', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(46, 'Info 10', 'Final content 10', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(47, 'Info 1', 'Content 1', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(48, 'Info 2', 'Sample content 2', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(49, 'Info 3', 'Example content 3', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(50, 'Info 4', 'Content data 4', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(51, 'Info 5', 'Another content 5', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(52, 'Info 6', 'Random content 6', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(53, 'Info 7', 'Test content 7', 2, 4, '2025-03-18 02:53:07', '2025-03-18 02:53:07'),
(54, 'Info 8', 'Demo content 8', 4, 4, '2025-03-18 02:53:07', '2025-03-18 08:33:59'),
(55, 'Info 9', 'Content sample 9', 4, 4, '2025-03-18 02:53:07', '2025-03-18 08:31:49'),
(56, 'Info 10', '1. Official thread cho phép đăng tải thông tin và thảo luận về tình hình xung đột vũ trang ở các điểm nóng trên thế giới như Israel - Hamas, Nga - Ukraine, tình hình Trung Đông, bán đảo Triều Tiên và các nơi khác... Đây là thread duy nhất được phép đăng tải những nội dung trên, lập thread điểm báo ngoài liên quan đến chúng sẽ bị xử lý.', 1, 4, '2025-03-18 02:53:07', '2025-03-18 09:34:40');

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
  `status` int DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL
) ;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `role`, `created_at`, `updated_at`, `status`, `email`) VALUES
(1, 'admin', '$2a$10$D.xurCPhfb4wzS.1XtOj6OSMs0yQg/O9FQwArjJkeTrPyvyZ4m50G', 'admin', '2025-03-12 04:18:05', '2025-03-13 02:15:09', 1, NULL),
(4, 'ducha', '$2a$10$UKUkk03Ui/TIor/SIYwtY.YyOzi2BxBl42cUpeve9zd/qiBX7U6Je', 'admin', '2025-03-12 08:14:57', '2025-03-13 02:15:09', 1, NULL),
(5, 'test', '$2a$10$RMo2XthFy2xqukjS9/CNqOXFgwMa7H3U4Va6geAxT1Yb.Bioz.eeq', 'user', '2025-03-13 09:20:45', '2025-03-13 09:20:45', 1, 'ducha.vnpay@gmail.com'),
(6, 'user', '$2a$10$jtBGldSmuKuK4lgFHjYymuTzpDkk11NrMW0HzrZwc/8YO5.iHlgiy', 'user', '2025-03-19 07:25:41', '2025-03-19 07:25:41', 1, 'ducha.vnpay@gmail.com'),
(19, 'ducha1@vnpay.vn', '', 'user', '2025-03-19 09:46:11', '2025-03-19 09:46:11', 1, 'ducha1@vnpay.vn');

-- --------------------------------------------------------

--
-- Table structure for table `user_logs`
--

CREATE TABLE `user_logs` (
  `id` int NOT NULL,
  `userid` int NOT NULL,
  `ip_public` varchar(50) DEFAULT NULL,
  `user_agent` varchar(255) DEFAULT NULL,
  `status_download_file` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `action` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `user_logs`
--

INSERT INTO `user_logs` (`id`, `userid`, `ip_public`, `user_agent`, `status_download_file`, `created_at`, `action`) VALUES
(1, 19, '192.168.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36', 1, '2025-03-20 08:27:32', NULL),
(2, 19, '192.168.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36', 0, '2025-03-20 08:28:00', NULL),
(3, 4, '::1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36', 0, '2025-03-20 09:20:15', NULL),
(4, 4, '210.245.49.42', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36', 0, '2025-03-20 09:21:34', NULL),
(5, 4, '210.245.49.42', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36', 0, '2025-03-20 09:30:31', NULL),
(6, 4, '210.245.49.42', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36', 0, '2025-03-20 09:31:23', 'login');

-- --------------------------------------------------------

--
-- Table structure for table `user_providers`
--

CREATE TABLE `user_providers` (
  `id` bigint UNSIGNED NOT NULL COMMENT 'Primary key, auto increment ID',
  `userid` bigint UNSIGNED NOT NULL COMMENT 'Foreign key liên kết với bảng users',
  `provider` varchar(50) NOT NULL COMMENT 'Tên provider đăng nhập, ví dụ: google, facebook, github',
  `provider_id` varchar(100) NOT NULL COMMENT 'ID duy nhất của người dùng bên provider (Google user id, Facebook user id,...)',
  `email` varchar(100) DEFAULT NULL COMMENT 'Email từ provider, dùng để đồng bộ thông tin',
  `name` varchar(100) DEFAULT NULL COMMENT 'Tên người dùng từ provider',
  `avatar` text COMMENT 'Link avatar/profile picture từ provider nếu có',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Thời gian tạo record',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Thời gian cập nhật record gần nhất'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='Bảng lưu thông tin đăng nhập từ các OAuth providers như Google, Facebook';

--
-- Dumping data for table `user_providers`
--

INSERT INTO `user_providers` (`id`, `userid`, `provider`, `provider_id`, `email`, `name`, `avatar`, `created_at`, `updated_at`) VALUES
(3, 19, 'google', '101275461407881757719', NULL, '', 'https://lh3.googleusercontent.com/a/ACg8ocKTc-vM7UnUlTESk36W0QRUWG8X9A8Z7QGgOZei83G5IJ6L89I=s96-c', '2025-03-19 09:46:17', '2025-03-19 09:46:17');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `google_tokens`
--
ALTER TABLE `google_tokens`
  ADD PRIMARY KEY (`id`),
  ADD KEY `userid` (`userid`);

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
-- Indexes for table `user_logs`
--
ALTER TABLE `user_logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_providers`
--
ALTER TABLE `user_providers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `provider` (`provider`,`provider_id`),
  ADD KEY `fk_user_provider` (`userid`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `google_tokens`
--
ALTER TABLE `google_tokens`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `response_history_info`
--
ALTER TABLE `response_history_info`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=57;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `user_logs`
--
ALTER TABLE `user_logs`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `user_providers`
--
ALTER TABLE `user_providers`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Primary key, auto increment ID', AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `google_tokens`
--
ALTER TABLE `google_tokens`
  ADD CONSTRAINT `google_tokens_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `user_providers`
--
ALTER TABLE `user_providers`
  ADD CONSTRAINT `fk_user_provider` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
