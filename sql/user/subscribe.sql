CREATE TABLE `subscribe` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `account` varchar(255) NOT NULL COMMENT '账号',
  `stock` int NOT NULL COMMENT '是否订阅股票信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;