CREATE TABLE IF NOT EXISTS `todos` (
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `account_id` INTEGER NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `is_completed` INTEGER NOT NULL DEFAULT 0,
  `is_deleted` INTEGER NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NULL
);

CREATE TABLE IF NOT EXISTS `accounts` (
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `email` VARCHAR(255) NOT NULL,
  `is_deleted` INTEGER NOT NULL DEFAULT 0,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NULL
);