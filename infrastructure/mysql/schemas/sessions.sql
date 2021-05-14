CREATE TABLE `sessions` (
  `user_id` char(26) COLLATE utf8mb4_bin NOT NULL COMMENT 'ユーザーID',
  `session_token` char(36) COLLATE utf8mb4_bin NOT NULL COMMENT 'セッショントークン',
  `expired_at` datetime NOT NULL COMMENT 'セッションの有効期限',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `uk_sessions_session_token` (`session_token`),
  CONSTRAINT `fk_sessions_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='セッション';
