/*
 * @Author: ChZheng
 * @Date: 2021-12-30 15:40:58
 * @LastEditTime: 2022-01-12 22:16:52
 * @LastEditors: ChZheng
 * @Description:
 * @FilePath: /go-programming-tour-book/blog-service/internal/routers/router.go
 */
package routers

import (
	"go-programming-tour-book/blog-service/internal/middleware"
	"go-programming-tour-book/blog-service/internal/routers/api"
	v1 "go-programming-tour-book/blog-service/internal/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()
	upload := api.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
