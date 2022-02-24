CREATE TABLE `bangumi` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `name` varchar(255) NOT NULL COMMENT '番剧名称',
  `url` varchar(255) NOT NULL COMMENT '指向链接',
  `year` int NOT NULL COMMENT '年份',
  `description` varchar(255) NOT NULL COMMENT '简介',
  `picurl` varchar(255) NOT NULL COMMENT '动漫封面',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8