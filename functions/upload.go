package functions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.String(http.StatusBadRequest, "上传失败")
		return
	}

	ctx.SaveUploadedFile(file, "./userFile"+"/"+file.Filename)

	ctx.String(http.StatusOK, "上传成功")
}
