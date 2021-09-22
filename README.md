# Wind's Server 个人网站服务端代码
- 需要自行启动 MySQL 和 Redis 服务端
- MySQL 建表指令如下：
  - blog 表
    ```sql
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
    ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
  - comments 表
    ```sql
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
    ```
  - collections 表
    ```sql
    CREATE TABLE `collections` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `url` varchar(255) NOT NULL COMMENT '链接',
    `description` varchar(255) NOT NULL COMMENT '描述',
    `picurl` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
  - urlinfo 表
    ```sql
    CREATE TABLE `urlinfo` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `origin` varchar(255) NOT NULL COMMENT '来源',
    `url` varchar(255) NOT NULL COMMENT '链接',
    `isDownload` int NOT NULL DEFAULT '0' COMMENT '资源是否已下载',
    `title` varchar(255) DEFAULT NULL COMMENT '标题',
    PRIMARY KEY (`id`),
    KEY `test` (`origin`,`url`,`title`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    ```
  - user 表
    ```sql
    CREATE TABLE `user` (
      `id` int NOT NULL AUTO_INCREMENT,
      `account` varchar(255) NOT NULL,
      `password` varchar(256) DEFAULT NULL,
      `username` varchar(40) NOT NULL,
      `level` int NOT NULL DEFAULT '0',
      `authority` int NOT NULL DEFAULT '0',
      `pic` varchar(255) NOT NULL COMMENT '头像',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
## TODO
- 博客系统完善；
- 新闻推送机制完善；
- 完成搜索功能；
- 横向和纵向内容扩展（未定）；
- 前端页面美化；
- 用户注册机制可更改为开放注册或邀请注册；
## 0.3 版
- 配置信息在 infomation/infomation.go 中，需要自行填写；
- 目前只能推送财经相关新闻；
- 更新为个人网站（已暂时关闭注册功能），访问地址：http://windserver.top/
## 0.2 版
- 订阅后自动推送（每天 6 点和 18 点，一次 10 条）；
- 需要在 infomation.go 中更改配置；
## 0.1 版本
- 需要自行配置数据库表，账号密码存储在 MySQL 中；
- 默认一次推送 20 条消息，可以自行修改；
- 验证码由 Redis 缓存 5 分钟后过期，需要重新发送；
- 使用 go 标准日志库记录错误；
- 配置信息在 infomation.go 文件中