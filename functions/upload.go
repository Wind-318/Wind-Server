package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	res, err := ctx.MultipartForm()
	if err != nil {
		ctx.String(http.StatusBadRequest, "上传失败")
		return
	}
	files := res.File["file"]

	for _, file := range files {
		ctx.SaveUploadedFile(file, "userFile"+"/"+file.Filename)
	}

	ctx.String(http.StatusOK, "上传成功")
}
