/*
 * @Author: ChZheng
 * @Date: 2022-01-10 23:37:38
 * @LastEditTime: 2022-01-12 22:13:16
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/routers/api/v1/upload.go
 */
package api

import (
	"go-programming-tour-book/blog-service/global"
	"go-programming-tour-book/blog-service/internal/service"
	"go-programming-tour-book/blog-service/pkg/app"
	"go-programming-tour-book/blog-service/pkg/errcode"
	"go-programming-tour-book/blog-service/pkg/upload"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if fileHeader != nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
