# Wind's Server 个人网站服务端架设流程
- 需要自行安装并启动 MySQL 和 Redis 服务端，需要配置 golang 环境
- 首先新建数据库，然后新建如下几张表，指令如下：
  - blog 表（记录文章）
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
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
  - comments 表（记录评论）
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
  - collections 表（记录收藏网站）
    ```sql
    CREATE TABLE `collections` (
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
    `url` varchar(255) NOT NULL COMMENT '链接',
    `description` varchar(255) NOT NULL COMMENT '描述',
    `picurl` varchar(255) NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
  - urlinfo 表（爬虫信息库）
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
  - user 表（用户表）
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
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
  - subscribe 表（订阅列表）
    ```sql
    CREATE TABLE `subscribe` (
      `id` int NOT NULL AUTO_INCREMENT COMMENT 'primary key',
      `account` varchar(255) NOT NULL COMMENT '用户邮箱',
      `stock` int NOT NULL COMMENT '股市相关新闻',
      PRIMARY KEY (`id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    ```
- 配置信息在 gofiles/config/config.go 中，需要根据情况自行填写；
- 需要自行添加 https 证书（或关闭 https）；
- 运行：go run main.go
## 访问地址：https://windserver.top/
## TODO
- [ ] **每季度新番列表导航；**
- [ ] **增加资源板块；**
- [ ] 用户注册机制可自定义为开放注册或邀请注册；
- [ ] 完成用户等级和权限系统；
- [ ] 完成文章权限系统；
- [ ] 增加积分制度；
- [ ] 消息推送机制优化（长期）；
- [ ] 博客系统优化及修复（长期）；
- [ ] 前端页面美化（长期）；
- [ ] 其它横向和纵向内容扩展（未定）；
# 版本更新说明
## 0.3.32
- 修复部分特性
## 0.3.31
- 修复更改标题后标题不改变的特性；
- 将新建文章时原地跳转到页面的行为修改为打开新窗口并跳转；
## 0.3.3 
- 支持头像修改；
- 更多注释；
## 0.3.2 
- 导航栏优化，目前导航栏默认为展开状态；
## 0.3.1 
- 增加查找文章功能；
- 修复了部分 ~~问题~~ 特性；
## 0.3.0 
- ~~更新~~ 发展为个人（多用户）网站（已暂时关闭注册功能）；
- 支持增加、删除、修改文章；
- 文章及评论都支持 markdown 格式且支持实时预览；
- 可以添加收藏的网站；
- 增加订阅推送数据表，当前可以选择推送股市相关消息，每天 6 点和 18 点会推送 10 条新闻链接到邮箱中；
- 还未开放订阅方式，当前只能在数据库中手动添加；
## 0.2 
- 订阅后自动推送（每天 6 点和 18 点，一次 10 条）；
- ~~需要在 infomation.go 中更改配置；~~
## 0.1 
- ~~需要自行配置数据库表，账号密码存储在 MySQL 中；~~
- 默认一次推送 ~~20~~ （10）条消息， 可以自行修改；
- 验证码由 Redis 缓存 ~~5 分钟~~（2 分钟）后过期，需要重新发送；
- ~~使用 go 标准日志库记录错误；~~
- ~~配置信息在 infomation.go 文件中~~