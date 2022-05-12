package main

import (
	"Project/config"
	"Project/gofiles"
	"Project/gofiles/bbs"
	"Project/gofiles/initCode"
	"Project/gofiles/mail"
	"Project/gofiles/spider/anime"
	"Project/gofiles/spider/sina"
	"Project/gofiles/storage"
	"Project/gofiles/user"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	// 初始化数据库
	initCode.InitDatabase()
	// 开启消费服务
	go mail.ConsumeCode()

	// 启用爬虫
	if config.AllowSpider {
		go func() {
			// 初始化片源地址
			//initCode.InitAnime()
			// 持续追踪更新
			go anime.ContinueGetNewAnime()
			// 抓取新闻
			go sina.CountTime()
			// 推送
			go sina.SendEveryUser()
		}()
	}

	// 新建路由
	router := gin.New()
	// 关闭控制台颜色
	gin.DisableConsoleColor()

	// 设置日志参数、路径
	f, _ := os.OpenFile("logs.log", os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0644)
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	if config.TurnOnHttps {
		// http 重定向 https
		router.Use(TlsHandler())
	}

	// 加载 html 文件
	router.LoadHTMLGlob("html/*")
	// 讨论区帖子
	router.Static("/bbsFile", "./bbsFile")
	// css 文件
	router.Static("/css", "./css")
	// 网站图标（默认为泉此方）
	router.StaticFile("/favicon.ico", "./favicon.ico")
	// js 文件
	router.Static("/js", "./js")
	// 系统图片文件
	router.Static("/picture", "./picture")
	// robots.txt 文件
	router.StaticFile("/robots.txt", "./robots.txt")
	// 用户文件
	router.Static("/userFile", "./userFile")

	// 设置 404 界面
	router.NoRoute(gofiles.ToNotFound)

	// 根路由
	root := router.Group("")
	{
		// 前往根目录
		root.GET("/", gofiles.ToHome)
		// 发送邮件
		root.GET("/sendStock", mail.SendCode)
	}

	// bbs 专区
	bbsRouter := router.Group("/bbs")
	{
		// 前往讨论区
		bbsRouter.GET("/", gofiles.ToBbs)
		// 前往文章编辑页面
		bbsRouter.GET("/CreateText", gofiles.ToCreateText)
		// 查询所有分类
		bbsRouter.GET("/InquireClassification", bbs.GetClassification)
		// 添加评论
		bbsRouter.POST("/AddComment", bbs.AddComment)
		// 作者
		bbsRouter.POST("/Author", bbs.Author)
		// 新建文章
		bbsRouter.POST("/CreateTexts", bbs.CreateText)
		// 删除文章
		bbsRouter.POST("/DeleteBlog", bbs.DeleteFromBlog)
		// 删除评论
		bbsRouter.POST("/DeleteComment", bbs.DeleteComment)
		// 获取评论 ID
		bbsRouter.POST("/GetCommentsID", bbs.GetCommentsID)
		// 获取用户文章
		bbsRouter.POST("/GetUserText", bbs.GetUserText)
		// 获取头像
		bbsRouter.POST("/GetProfile", bbs.GetProfile)
		// 获取最后一次编辑时间
		bbsRouter.POST("/GetLastModify", bbs.GetLastModify)
		// 获取修改文章信息
		bbsRouter.POST("/GetModifyBlog", bbs.GetModifyBlog)
		// 获取图片 url
		bbsRouter.POST("/Getpicurl", bbs.Getpicurl)
		// 获取页面数量
		bbsRouter.POST("/InquirePageNums", bbs.GetPageNums)
		// 获取文章内容
		bbsRouter.POST("/InquireText", bbs.GetText)
		// 编辑文章
		bbsRouter.POST("/ModifyBlog", bbs.ModifyBlog)
		// 赞
		bbsRouter.POST("/Parise", bbs.Parise)
		// 获取赞数
		bbsRouter.POST("/PariseNum", bbs.PariseNum)
		// 搜索
		bbsRouter.POST("/Search", bbs.Search)
		// 评论
		bbsRouter.POST("/TextComment", bbs.TextComment)
		// 浏览量
		bbsRouter.POST("/Views", bbs.Views)
	}

	// 用户路由
	users := router.Group("/user")
	{
		users.GET("/Exit", gofiles.Exit)
		users.GET("/ToLogin", gofiles.ToLogin)
		users.GET("/TochangePassword", gofiles.ToChangePassword)
		users.GET("/signAddScore", user.SignAddScore)
		users.GET("/getUsersName", user.GetUsersNames)
		users.GET("/checkPermission", user.CheckPermission)
		users.GET("/IsSystem", user.IsSystem)

		users.POST("/IsSystems", user.IsSystemOrAuthor)
		users.POST("/changePassword", user.ChangePassWord)
		users.POST("/login", user.Login)
		users.POST("/register", user.Register)
		users.POST("/sendCode", mail.SendCode)
		users.POST("/UploadProfile", bbs.UploadProfile)
	}

	// 收藏路由
	collections := router.Group("/collections")
	{
		collections.GET("/", gofiles.ToCollections)
		collections.GET("/GetWebs", storage.GetWebs)
		collections.POST("/PutWebs", storage.PutWebs)
		collections.POST("/PutPic", storage.PutPic)
	}

	// 资源路由（还没做）
	Resources := router.Group("/resources")
	{
		Resources.GET("/", gofiles.ToResources)
	}

	// 动漫路由
	animes := router.Group("/anime")
	{
		// 根路径
		animes.GET("/", gofiles.ToAnime)
		// 选出新番
		animes.GET("/searchNewAnime", anime.SearchNewAnime)
		// 搜索
		animes.POST("/search", anime.Search)
		// 按年检索
		animes.POST("/searchByYear", anime.SearchByYear)
	}

	// 存储路由
	storages := router.Group("/storage")
	{
		// 根路径
		storages.GET("/", gofiles.ToStorage)

		// 新建文件夹
		storages.POST("/makeDirectory", storage.MakeDirectory)
		// 获取用户文件夹数量
		storages.POST("/getUserFileNums", storage.GetUserFileNums)
		// 计算用户文件夹中图片分页数量
		storages.POST("/getUserStoragePicturePage", storage.GetUserStoragePicturePage)
		// 获取用户存储的图片
		storages.POST("/getUserStoragePicture", storage.GetUserStoragePicture)
		// 上传图片
		storages.POST("/stroageImg", storage.StorageFiles)
		// 批量下载
		storages.POST("/downloadFiles", storage.DownloadFiles)
	}

	// 开启 https
	if config.TurnOnHttps {
		// 监听 http
		go router.Run(":80")
		// 监听 https，自行选择 SSL 证书
		router.RunTLS(":443", config.HttpsCertification1, config.HttpsCertification2)
	} else {
		// 监听 http
		router.Run(":80")
	}
}

// 重定向到 https
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     config.YourName + ":443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
