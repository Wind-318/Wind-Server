package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 上传文件
func Upload(ctx *gin.Context) {
	res, err := ctx.MultipartForm()
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
		return
	}
	files := res.File["file"]

	for _, file := range files {
		err = ctx.SaveUploadedFile(file, "userFile"+"/"+file.Filename)
		if err != nil {
			ctx.HTML(http.StatusInternalServerError, "serverError.html", nil)
			return
		}
	}

	ctx.HTML(http.StatusOK, "success.html", nil)
}
