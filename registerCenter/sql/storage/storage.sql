CREATE TABLE `storage` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `account` varchar(255) NOT NULL COMMENT '账号',
  `size` bigint NOT NULL DEFAULT '0' COMMENT '文件大小',
  `filepath` varchar(255) NOT NULL COMMENT '归属文件夹',
  `name` varchar(255) NOT NULL COMMENT '文件名',
  `type` varchar(255) NOT NULL COMMENT '文件类型',
  `path` varchar(255) NOT NULL COMMENT '文件路径',
  `smallpath` varchar(255) DEFAULT NULL COMMENT '缩略图路径',
  PRIMARY KEY (`id`),
  KEY `index1` (`path`),
  KEY `indexName1` (`smallpath`),
  KEY `indexName2` (`path`),
  KEY `indexName3` (`type`),
  KEY `indexName4` (`name`),
  KEY `indexName5` (`filepath`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;