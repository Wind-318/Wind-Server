CREATE TABLE `storagefile` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `account` varchar(255) NOT NULL COMMENT '账号',
  `filename` varchar(255) NOT NULL COMMENT '文件夹',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;