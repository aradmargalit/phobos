# ======== USERS ========
CREATE TABLE `users` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `name` text,
  `given_name` text,
  `email` varchar(200) DEFAULT NULL,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY `email` (`email`)
);

# ======== ACTIVITY TYPES ========
CREATE TABLE `activity_types` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `name` text,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

# ======== ACTIVITIES ========
CREATE TABLE `activities` (
  `id` int(11) AUTO_INCREMENT NOT NULL,
  `name` text,
  `activity_date` DATETIME,
  `activity_type_id` int(11),
  `duration` DOUBLE,
  `distance` DOUBLE,
  `unit` text,
  `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);