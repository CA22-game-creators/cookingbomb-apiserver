CREATE TABLE `users` (
  `id` char(26) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザULID',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザ名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='ユーザ';
