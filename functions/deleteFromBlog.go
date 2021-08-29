package functions

import (
	"Project/infomation"
	"encoding/json"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 删除一篇文章
func DeleteFromBlog(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	var ids []string
	jsonArr := ctx.PostForm("checked")
	json.Unmarshal([]byte(jsonArr), &ids)
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

	for _, id := range ids {
		conn.Exec("DELETE FROM blog WHERE id = ?", id)
	}
}

// 删除评论
func DeleteComment(ctx *gin.Context) {
	id := ctx.PostForm("id")

	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

	conn.Exec("DELETE FROM comments WHERE id = ?", id)
}

func ModifyBlog(ctx *gin.Context) {
	_, err := ctx.Cookie("cookie")
	if err != nil {
		return
	}
	var ids []string
	jsonArr := ctx.PostForm("checked")
	json.Unmarshal([]byte(jsonArr), &ids)
	conn := sqlx.MustConnect("mysql", infomation.MySQLInfo)
	defer conn.Close()

}
