package storage

import (
	"Project/gofiles/config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DeleteFiles(ctx *gin.Context) {
	conn, err := sqlx.Connect("mysql", config.MySQLInfo)
	if err != nil {
		return
	}
	defer conn.Close()
}
