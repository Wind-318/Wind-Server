CREATE TABLE `urlinfo` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `origin` varchar(255) NOT NULL COMMENT '来源',
  `url` varchar(255) NOT NULL COMMENT '链接',
  `isDownload` int NOT NULL DEFAULT '0' COMMENT '资源是否已下载',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  PRIMARY KEY (`id`),
  KEY `test` (`origin`,`url`,`title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;