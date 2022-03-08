CREATE TABLE `storage` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `account` varchar(255) NOT NULL COMMENT '账号',
  `size` bigint NOT NULL DEFAULT '0' COMMENT '文件大小',
  `filePath` varchar(255) NOT NULL COMMENT '归属文件夹',
  `name` varchar(255) NOT NULL COMMENT '文件名',
  `type` varchar(255) NOT NULL COMMENT '文件类型',
  `path` varchar(255) NOT NULL COMMENT '文件路径',
  `smallpath` varchar(255) DEFAULT NULL COMMENT '缩略图路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci