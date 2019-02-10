-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: Feb 10, 2019 at 06:30 PM
-- Server version: 5.6.41
-- PHP Version: 7.2.10

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `chat`
--

-- --------------------------------------------------------

--
-- Table structure for table `chats`
--

CREATE TABLE `chats` (
  `id` int(11) NOT NULL,
  `name` char(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `chats`
--

INSERT INTO `chats` (`id`, `name`) VALUES
(1, 'First chat'),
(2, 'Second chat'),
(3, 'Third chat');

-- --------------------------------------------------------

--
-- Table structure for table `messages`
--

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
  `chat_id` int(11) NOT NULL,
  `from_user` int(11) NOT NULL,
  `message` varchar(535) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `messages`
--

INSERT INTO `messages` (`id`, `chat_id`, `from_user`, `message`, `created_at`) VALUES
(1, 1, 1, 'Сообщение в первом чате', '2019-02-03 20:00:46'),
(2, 1, 2, 'Вот и сообщение петраа в первом чате', '2019-02-03 20:00:46'),
(3, 1, 1, 'Заглушка для сообщения', '2019-02-04 14:07:35'),
(4, 2, 2, 'Заглушка для сообщения', '2019-02-04 14:39:56'),
(5, 3, 1, 'Заглушка для сообщения', '2019-02-04 14:59:08'),
(6, 3, 3, 'message to chat 1 from user 3', '2019-02-04 14:59:45'),
(7, 3, 3, 'message to chat 1 from user 3', '2019-02-05 11:57:07'),
(8, 1, 1, 'message to chat 1 from user 3', '2019-02-05 11:58:33'),
(9, 1, 3, 'message to chat 1 from user 3', '2019-02-05 11:58:56'),
(10, 1, 1, 'message to chat 1 from user 3', '2019-02-05 12:01:12'),
(11, 1, 3, 'message to chat 1 from user 3', '2019-02-05 12:04:21'),
(12, 1, 3, 'message to chat 1 from user 3', '2019-02-05 12:05:11'),
(13, 1, 3, 'message to chat 1 from user 3', '2019-02-05 12:09:50'),
(14, 3, 3, 'message to chat 1 from user 3', '2019-02-05 12:10:40'),
(15, 1, 3, 'message to chat 1 from user 3', '2019-02-05 14:56:55'),
(25, 2, 3, 'thisistestMessga', '2019-02-06 12:24:42'),
(35, 1, 0, '1', '2019-02-06 13:45:53'),
(36, 2, 0, '2', '2019-02-06 13:47:06'),
(37, 2, 0, '4', '2019-02-06 13:47:10'),
(38, 1, 0, '1', '2019-02-06 13:47:24'),
(39, 1, 0, '2', '2019-02-06 13:47:28'),
(40, 3, 0, '424', '2019-02-06 13:48:01'),
(41, 2, 0, '123', '2019-02-06 13:49:40'),
(42, 3, 0, '231', '2019-02-06 13:49:42'),
(43, 2, 0, '2\ndffdsf\n', '2019-02-06 14:46:12'),
(44, 2, 0, 'dsfdsf\n', '2019-02-06 14:46:29'),
(45, 1, 0, '1', '2019-02-06 14:49:03'),
(46, 2, 0, '111', '2019-02-06 14:51:48'),
(47, 2, 0, '1', '2019-02-06 15:24:54'),
(48, 2, 0, 'ww', '2019-02-06 15:26:04'),
(49, 1, 0, 'dsf', '2019-02-06 15:26:24'),
(50, 1, 0, '1', '2019-02-06 15:27:10'),
(51, 1, 0, '1', '2019-02-06 15:28:01'),
(52, 3, 0, '4', '2019-02-06 15:28:13'),
(53, 3, 0, '<scrit>alert(1)</script>\n', '2019-02-06 15:35:49'),
(54, 3, 0, '<script>alert(1)</script>', '2019-02-06 15:36:08'),
(55, 3, 0, '<scrit>alert(1);</script>', '2019-02-06 15:36:30'),
(56, 1, 1, '1', '2019-02-10 14:02:06');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` char(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`) VALUES
(1, 'Ivan Ivanov'),
(2, 'Petr Petrov'),
(3, 'Oleg Gerasimovich');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `chats`
--
ALTER TABLE `chats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=57;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
