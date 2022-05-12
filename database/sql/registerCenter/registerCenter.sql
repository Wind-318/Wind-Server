CREATE TABLE `register` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'Primary Key',
  `ip` varchar(255) NOT NULL COMMENT 'IP 地址',
  `port` varchar(255) NOT NULL COMMENT '端口地址',
  `serviceName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='newTable';