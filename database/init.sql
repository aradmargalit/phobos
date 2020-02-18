CREATE TABLE `users` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `name` text,
  `given_name` text,
  `email` varchar(200) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY `email` (`email`)
)