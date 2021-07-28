## database
- ### MySQL
  ```Go
  package main

  import (
      _ "github.com/go-sql-driver/mysql"
      "github.com/jmoiron/sqlx"
  )

  func main() {
      db, e := sqlx.Connect("mysql", "root:密码@tcp(localhost:3306)/数据库名称")
  }
  ```
- ### Redis
  ```Go
  package main

  import (
      "fmt"
      "github.com/gomodule/redigo/redis"
  )

  func main() {
      conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
      defer conn.Close()
      reply, _ := conn.Do("mget", "name1", "name2")
      fmt.Println(redis.Strings(reply, nil))
  }
  ```