CREATE Database chat_room;

DROP TABLE IF EXISTS `users`;
CREATE TABLE IF NOT EXISTS `users` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uuid` varchar(150) NOT NULL COMMENT 'uuid',
    `username` varchar(191) NOT NULL COMMENT '''用户名''',
    `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
    `email` varchar(80) DEFAULT NULL COMMENT '邮箱',
    `password` varchar(150) NOT NULL COMMENT '密码',
    `avatar` varchar(250) NOT NULL COMMENT '头像',
    `create_at` datetime(3) DEFAULT NULL,
    `update_at` datetime(3) DEFAULT NULL,
    `delete_at` bigint DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `username` (`username`),
    UNIQUE KEY `idx_uuid` (`uuid`),
    UNIQUE KEY `username_2` (`username`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT '用户表';