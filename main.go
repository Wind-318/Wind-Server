package main

import (
	"Project/gofiles"
	"Project/gofiles/anime"
	"Project/gofiles/blogfunc"
	"Project/gofiles/collectionfunc"
	"Project/gofiles/initCode"
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
	go func() {
		// 初始化片源地址
		initCode.InitAnime()
		// 持续追踪更新
		anime.ContinueGetNewAnime()
	}()
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
	// http 重定向 https
	router.Use(TlsHandler())

	// 计时抓取信息到数据库
	go initCode.CountTime()

	// 发送邮件
	go initCode.SendEveryUser()

	// 加载 html 文件
	router.LoadHTMLGlob("html/*")
	// 讨论区帖子
	router.Static("/blog", "./blog")
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
		root.GET("/", gofiles.ToHead)
		// 发送邮件
		root.GET("/sendStock", sina.SendStock)
		// 服务器错误
		root.GET("/serverError", gofiles.ToError)
	}

	// 讨论区路由
	blog := router.Group("/blogs")
	{
		// 前往讨论区
		blog.GET("/", gofiles.ToBlog)
		// 前往文章编辑页面
		blog.GET("/CreateText", gofiles.ToCreateText)
		// 查询所有分类
		blog.GET("/InquireClassification", blogfunc.GetClassification)
		// 添加评论
		blog.POST("/AddComment", blogfunc.AddComment)
		// 作者
		blog.POST("/Author", blogfunc.Author)
		// 新建文章
		blog.POST("/CreateTexts", blogfunc.CreateText)
		// 删除文章
		blog.POST("/DeleteBlog", blogfunc.DeleteFromBlog)
		// 删除评论
		blog.POST("/DeleteComment", blogfunc.DeleteComment)
		// 获取评论 ID
		blog.POST("/GetCommentsID", blogfunc.GetCommentsID)
		// 获取用户文章
		blog.POST("/GetUserText", blogfunc.GetUserText)
		// 获取头像
		blog.POST("/GetProfile", blogfunc.GetProfile)
		// 获取最后一次编辑时间
		blog.POST("/GetLastModify", blogfunc.GetLastModify)
		// 获取修改文章信息
		blog.POST("/GetModifyBlog", blogfunc.GetModifyBlog)
		// 获取图片 url
		blog.POST("/Getpicurl", blogfunc.Getpicurl)
		// 获取页面数量
		blog.POST("/InquirePageNums", blogfunc.GetPageNums)
		blog.POST("/InquireText", blogfunc.GetText)
		blog.POST("/ModifyBlog", blogfunc.ModifyBlog)
		blog.POST("/Parise", blogfunc.Parise)
		blog.POST("/PariseNum", blogfunc.PariseNum)
		blog.POST("/Search", blogfunc.Search)
		blog.POST("/TextComment", blogfunc.TextComment)
		blog.POST("/Views", blogfunc.Views)
	}

	// 用户路由
	users := router.Group("/user")
	{
		users.GET("/Exit", gofiles.Exit)
		users.GET("/ToLogin", gofiles.ToLogin)
		users.GET("/TochangePassword", gofiles.ToChangePassword)
		users.GET("/signAddScore", user.SignAddScore)
		users.GET("/getUsersName", user.GetUsersName)

		users.POST("/changePassword", user.ChangePassWord)
		users.POST("/login", user.Login)
		users.POST("/register", user.Register)
		users.POST("/sendCode", sina.SendCode)
		users.POST("/verificationFind", user.VerificationFind)
		users.POST("/UploadProfile", blogfunc.UploadProfile)
	}

	// 收藏路由
	collections := router.Group("/collections")
	{
		collections.GET("/", gofiles.ToCollections)
		collections.GET("/GetWebs", collectionfunc.GetWebs)
		collections.GET("/IsSystem", collectionfunc.IsSystem)
		collections.POST("/IsSystems", collectionfunc.IsSystems)
		collections.POST("/PutWebs", collectionfunc.PutWebs)
		collections.POST("/PutPic", collectionfunc.PutPic)
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
		animes.GET("/checkPermission", anime.CheckPermission)
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

	// 监听 http
	go router.Run(":80")
	// 监听 https，自行选择 SSL 证书
	router.RunTLS(":443", "windserver.top.pem", "windserver.top.key")
}

// 重定向到 https
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "windserver.top:443",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			return
		}
		c.Next()
	}
}
