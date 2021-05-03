CREATE TABLE `users` (
  `id` char(26) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーULID',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザー名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='ユーザー';
