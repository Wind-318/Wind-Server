package anime

import (
	"Project/gofiles/algorithm"
	"Project/gofiles/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 动漫属性
type AnimeInfo struct {
	// 名称
	Name string `db:"name"`
	// bangumi 链接
	Url string `db:"url"`
	// 年份
	Year int `db:"year"`
	// 简介
	Description string `db:"description"`
	// 播放来源
	Source []string `db:"source"`
	// 播放地址
	Urls []string `db:"urls"`
	// 封面地址
	Picurl string `db:"picurl"`
}

// 搜索动漫
func Search(ctx *gin.Context) {
	text := ctx.PostForm("text")
	result := map[string]interface{}{}

	// 若搜索内容全为空格，直接返回空
	isEmpty := true
	for i := range text {
		if text[i] != ' ' {
			isEmpty = false
			break
		}
	}
	if isEmpty {
		result["count"] = 0
		ctx.JSON(http.StatusOK, result)
		return
	}

	conn := sqlx.MustConnect("mysql", config.MySQLInfo)
	defer conn.Close()

	names := make([]string, 0)
	idNums := []string{}
	conn.Select(&names, "SELECT name FROM bangumi")

	for index, name := range names {
		if algorithm.Match(name, text) == 1 {
			tempInfo := AnimeInfo{}
			conn.Get(&tempInfo, "SELECT name, url, year, description, picurl FROM bangumi WHERE name = ?", name)
			conn.Select(&tempInfo.Source, "SELECT source FROM animesource WHERE anime = ?", name)
			conn.Select(&tempInfo.Urls, "SELECT urls FROM animesource WHERE anime = ?", name)
			result[strconv.Itoa(index)] = tempInfo
			idNums = append(idNums, strconv.Itoa(index))
			// 超过 1000 时自动停止搜索
			if len(idNums) >= 1000 {
				break
			}
		}
	}
	result["count"] = idNums

	ctx.JSON(http.StatusOK, result)
}
