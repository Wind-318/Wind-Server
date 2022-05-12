CREATE TABLE `bangumi` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `name` varchar(1024) NOT NULL COMMENT '番剧名称',
  `url` varchar(1024) NOT NULL COMMENT '指向链接',
  `year` varchar(255) NOT NULL COMMENT '年份',
  `description` varchar(4096) NOT NULL COMMENT '简介',
  `picurl` varchar(255) NOT NULL COMMENT '动漫封面',
  `isNew` int DEFAULT '0' COMMENT '新番判定',
  PRIMARY KEY (`id`),
  KEY `index_1` (`name`),
  KEY `index_2` (`url`),
  KEY `index_3` (`year`),
  KEY `index_11` (`url`),
  KEY `index_12` (`year`),
  KEY `index_14` (`picurl`),
  KEY `index_15` (`isNew`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;
