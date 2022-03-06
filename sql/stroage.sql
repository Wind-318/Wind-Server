CREATE TABLE `stroage` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `account` varchar(255) NOT NULL COMMENT '账号',
  `type` varchar(255) NOT NULL COMMENT '文件类型',
  `path` varchar(1024) NOT NULL COMMENT '文件路径',
  `size` int DEFAULT '0' COMMENT '文件大小',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci