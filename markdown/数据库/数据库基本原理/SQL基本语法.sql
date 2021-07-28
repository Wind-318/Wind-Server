-- 创建一个teach数据库
CREATE DATABASE teach;
-- 使用teach数据库
USE teach;
--创建一个test表
CREATE TABLE test(
	-- ID作为主键，自动增长
    ID INT UNSIGNED PRIMARY KEY AUTO_INCREMENT ,
	-- varchar类型的name，长度最大为50
    NAME VARCHAR(50) NOT NULL
);
-- 插入一个路人甲和路人乙
INSERT INTO test VALUES(1, "路人甲");
INSERT INTO test VALUES(2, "路人乙");
-- 删除路人乙
DELETE FROM test WHERE NAME = "路人乙";
-- 删除table表
DROP TABLE test;
-- 将路人甲改名为龙套丙
UPDATE test SET NAME = "龙套丙" WHERE ID = 1;
-- 给表中添加一行年龄字段
ALTER TABLE test ADD COLUMN AGE INT UNSIGNED;
-- 查询表中所有信息
SELECT * FROM test;