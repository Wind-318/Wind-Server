package anime

import (
	"Project/config"
	"Project/gofiles/algorithm"
	"Project/gofiles/info"
	"Project/gofiles/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// 搜索功能
func Search(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 检查权限，非登录状态直接返回空
	if !user.IsLogin(ctx) {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 用户输入字段
	text := ctx.PostForm("text")

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

	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo+"spider")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()

	// 选出名称列表
	names := make([]string, 0)
	// 计数
	idNums := []string{}
	conn.Select(&names, "SELECT name FROM bangumi")

	// 轮询数据库查询
	for index := len(names) - 1; index >= 0; index-- {
		// 进行匹配
		if algorithm.Match(names[index], text) {
			// 选出数据
			tempInfo := info.AnimeInfo{}
			conn.Get(&tempInfo, "SELECT name, url, year, description, picurl FROM bangumi WHERE name = ?", names[index])
			conn.Select(&tempInfo.Source, "SELECT source FROM animesource WHERE anime = ?", names[index])
			conn.Select(&tempInfo.Urls, "SELECT urls FROM animesource WHERE anime = ?", names[index])

			result[strconv.Itoa(index)] = tempInfo

			// 数量 + 1
			idNums = append(idNums, strconv.Itoa(index))
			// 超过 1000 时自动停止搜索
			if len(idNums) >= 1000 {
				break
			}
		}
	}
	// 得到数量
	result["count"] = idNums

	// 返回数据
	ctx.JSON(http.StatusOK, result)
}

// 选出新番
func SearchNewAnime(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 检查权限，非登录状态直接返回空
	if !user.IsLogin(ctx) {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo+"spider")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()
	// 获取名称
	names := make([]string, 0)
	// 计数
	idNums := []string{}
	// isNew 为 1 则为当季动漫
	conn.Select(&names, "SELECT name FROM bangumi WHERE isNew = 1")
	// 逻辑同上
	for index := len(names) - 1; index >= 0; index-- {
		tempInfo := info.AnimeInfo{}
		conn.Get(&tempInfo, "SELECT name, url, year, description, picurl FROM bangumi WHERE name = ?", names[index])
		conn.Select(&tempInfo.Source, "SELECT source FROM animesource WHERE anime = ?", names[index])
		conn.Select(&tempInfo.Urls, "SELECT urls FROM animesource WHERE anime = ?", names[index])

		result[strconv.Itoa(index)] = tempInfo
		idNums = append(idNums, strconv.Itoa(index))
	}
	result["count"] = idNums

	// 返回数据
	ctx.JSON(http.StatusOK, result)
}

// 选出指定年份番剧
func SearchByYear(ctx *gin.Context) {
	result := map[string]interface{}{}
	// 检查权限，非登录状态直接返回空
	if !user.IsLogin(ctx) {
		result["msg"] = "尚未登陆！"
		ctx.JSON(http.StatusOK, result)
		return
	}
	// 选取的年份
	year := ctx.PostForm("year")
	// 连接数据库
	conn, err := sqlx.Connect("mysql", config.MySQLInfo+"spider")
	if err != nil {
		result["msg"] = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	defer conn.Close()
	// 获取名称
	names := make([]string, 0)
	// 计数
	idNums := []string{}
	// 看年份大于 2000 还是小于等于 2000
	if y, _ := strconv.Atoi(year); y > 2000 {
		conn.Select(&names, "SELECT name FROM bangumi WHERE year = ?", year)
	} else {
		conn.Select(&names, "SELECT name FROM bangumi WHERE year <= ?", year)
	}
	// 逻辑同上
	for index := len(names) - 1; index >= 0; index-- {
		tempInfo := info.AnimeInfo{}
		conn.Get(&tempInfo, "SELECT name, url, year, description, picurl FROM bangumi WHERE name = ?", names[index])
		conn.Select(&tempInfo.Source, "SELECT source FROM animesource WHERE anime = ?", names[index])
		conn.Select(&tempInfo.Urls, "SELECT urls FROM animesource WHERE anime = ?", names[index])

		result[strconv.Itoa(index)] = tempInfo
		idNums = append(idNums, strconv.Itoa(index))
	}
	result["count"] = idNums

	ctx.JSON(http.StatusOK, result)
}
