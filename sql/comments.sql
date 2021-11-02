CREATE TABLE `comments` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `blog` int DEFAULT NULL,
  `content` varchar(3000) NOT NULL COMMENT '内容',
  `create_time` datetime DEFAULT NULL COMMENT 'create time',
  `update_time` datetime DEFAULT NULL COMMENT 'update time',
  `parent` int DEFAULT '-1' COMMENT '父级评论',
  `pic` varchar(255) NOT NULL COMMENT '头像',
  `author` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;