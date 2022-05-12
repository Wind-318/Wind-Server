CREATE TABLE `collections` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `url` varchar(255) NOT NULL COMMENT '链接',
  `description` varchar(255) NOT NULL COMMENT '描述',
  `picurl` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;