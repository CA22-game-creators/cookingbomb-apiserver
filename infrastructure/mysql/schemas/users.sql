CREATE TABLE `users` (
  `id` char(26) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーULID',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザー名',
  `auth_token` char(60) COLLATE utf8mb4_bin NOT NULL COMMENT '認証トークンのハッシュ値',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='ユーザー';
