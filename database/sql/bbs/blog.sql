CREATE TABLE `blog` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `author` varchar(255) NOT NULL COMMENT '作者',
  `authoremail` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL COMMENT '标题',
  `description` varchar(255) NOT NULL COMMENT '简介',
  `content` text NOT NULL COMMENT '内容',
  `types` varchar(255) NOT NULL,
  `clicknum` int NOT NULL COMMENT '点击数',
  `great` int NOT NULL COMMENT '点赞数',
  `authority` int NOT NULL,
  `create_time` datetime DEFAULT NULL COMMENT 'create time',
  `update_time` datetime DEFAULT NULL COMMENT 'update time',
  `authorid` int NOT NULL,
  `url` varchar(255) NOT NULL,
  `isdelete` int NOT NULL,
  `picurl` varchar(255) NOT NULL,
  `smallpic` varchar(255) NOT NULL COMMENT '缩略图',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
