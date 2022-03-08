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
	initCode.InitDatabase()
	go func() {
		initCode.InitAnime()
		initCode.ContinueGetNewAnime()
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

	// 加载静态资源
	router.LoadHTMLGlob("html/*")
	router.Static("/blog", "./blog")
	router.Static("/css", "./css")
	router.Static("/files", "./files")
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/js", "./js")
	router.Static("/music", "./music")
	router.Static("/picture", "./picture")
	router.StaticFile("/robots.txt", "./robots.txt")
	router.Static("/userFile", "./userFile")

	// 设置 404 界面
	router.NoRoute(gofiles.ToNotFound)

	// 设置根路由
	root := router.Group("")
	{
		root.GET("/", gofiles.ToHead)
		root.Any("/blog", gofiles.ToNotFound)
		root.Any("/css", gofiles.ToNotFound)
		root.Any("/files", gofiles.ToNotFound)
		root.Any("/js", gofiles.ToNotFound)
		root.Any("/music", gofiles.ToNotFound)
		root.Any("/picture", gofiles.ToNotFound)

		root.GET("/sendStock", sina.SendStock)
		root.GET("/serverError", gofiles.ToError)
	}

	// 设置博客路由
	blog := router.Group("/blogs")
	{
		blog.GET("/", gofiles.ToBlog)
		blog.GET("/CreateText", gofiles.ToCreateText)
		blog.GET("/InquireClassification", blogfunc.GetClassification)
		blog.POST("/AddComment", blogfunc.AddComment)
		blog.POST("/Author", blogfunc.Author)
		blog.POST("/CreateTexts", blogfunc.CreateText)
		blog.POST("/DeleteBlog", blogfunc.DeleteFromBlog)
		blog.POST("/DeleteComment", blogfunc.DeleteComment)
		blog.POST("/GetCommentsID", blogfunc.GetCommentsID)
		blog.POST("/GetUserText", blogfunc.GetUserText)
		blog.POST("/GetProfile", blogfunc.GetProfile)
		blog.POST("/GetLastModify", blogfunc.GetLastModify)
		blog.POST("/GetModifyBlog", blogfunc.GetModifyBlog)
		blog.POST("/Getpicurl", blogfunc.Getpicurl)
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

	// 资源路由
	Resources := router.Group("/resources")
	{
		Resources.GET("/", gofiles.ToResources)
	}

	// 动漫路由
	animes := router.Group("/anime")
	{
		animes.GET("/", gofiles.ToAnime)
		animes.GET("/checkPermission", anime.CheckPermission)
		animes.GET("/searchNewAnime", anime.SearchNewAnime)

		animes.POST("/search", anime.Search)
		animes.POST("/searchByYear", anime.SearchByYear)
	}

	// 存储路由
	storages := router.Group("/storage")
	{
		storages.GET("/", gofiles.ToStorage)

		storages.POST("/makeDirectory", storage.MakeDirectory)
		storages.POST("/getUserFileNums", storage.GetUserFileNums)
		storages.POST("/getUserStoragePicturePage", storage.GetUserStoragePicturePage)
		storages.POST("/getUserStoragePicture", storage.GetUserStoragePicture)
		storages.POST("/stroageImg", storage.StorageFiles)
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
