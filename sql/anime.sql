CREATE TABLE `animesource` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `anime` varchar(1024) NOT NULL COMMENT '动漫名称',
  `source` varchar(255) NOT NULL COMMENT '播放来源',
  `urls` varchar(1024) NOT NULL COMMENT '播放地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8