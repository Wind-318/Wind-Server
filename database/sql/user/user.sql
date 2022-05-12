CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL COMMENT '账号',
  `password` varchar(256) DEFAULT NULL COMMENT '密码',
  `username` varchar(40) NOT NULL COMMENT '用户名',
  `level` int NOT NULL DEFAULT '0' COMMENT '等级',
  `authority` int NOT NULL DEFAULT '0' COMMENT '权限',
  `pic` varchar(255) NOT NULL COMMENT '头像',
  `score` int DEFAULT '0' COMMENT '积分',
  `capacity` bigint DEFAULT '0' COMMENT '存储容量',
  `unusedCapacity` bigint NOT NULL DEFAULT '0' COMMENT '未用存储容量',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;