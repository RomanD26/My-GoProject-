-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Хост: 127.0.0.1:8889
-- Время создания: Июн 17 2021 г., 19:55
-- Версия сервера: 8.0.19
-- Версия PHP: 7.1.33
SET
  SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET
  time_zone = "+00:00";
  /*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
  /*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
  /*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
  /*!40101 SET NAMES utf8mb4 */;
--
  -- База данных: `gobd`
  --
  -- --------------------------------------------------------
  --
  -- Структура таблицы `users`
  --
  CREATE TABLE `users` (
    `id` int UNSIGNED NOT NULL,
    `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
    `age` int UNSIGNED NOT NULL
  ) ENGINE = MyISAM DEFAULT CHARSET = utf8;
--
  -- Дамп данных таблицы `users`
  --
INSERT INTO
  `users` (`id`, `name`, `age`)
VALUES
  (1, 'bob', 23),
  (2, 'sem', 23),
  (3, 'david', 23),
  (4, 'adolf', 12),
  (5, 'wer', 23),
  (6, 'koli', 23),
  (7, 'elf', 34);
--
  -- Индексы сохранённых таблиц
  --
  --
  -- Индексы таблицы `users`
  --
ALTER TABLE
  `users`
ADD
  PRIMARY KEY (`id`);
--
  -- AUTO_INCREMENT для сохранённых таблиц
  --
  --
  -- AUTO_INCREMENT для таблицы `users`
  --
ALTER TABLE
  `users`
MODIFY
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  AUTO_INCREMENT = 8;
COMMIT;
  /*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
  /*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
  /*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;