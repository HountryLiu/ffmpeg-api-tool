package router

import (
	"github.com/HountryLiu/ffmpeg-api-tool/controller/ffmpeg"
	_ "github.com/HountryLiu/ffmpeg-api-tool/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine) {
	//swagger集成
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := router.Group("/api")

	//goquery学习
	api_ffmpeg := api.Group("/ffmpeg")
	ffmpeg.Router(api_ffmpeg)
}
