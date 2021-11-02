CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `password` varchar(256) DEFAULT NULL,
  `username` varchar(40) NOT NULL,
  `level` int NOT NULL DEFAULT '0',
  `authority` int NOT NULL DEFAULT '0',
  `pic` varchar(255) NOT NULL COMMENT '头像',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;