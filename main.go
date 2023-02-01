package main

import (
	"os"

	global_router "github.com/HountryLiu/ffmpeg-api-tool/router"
	"github.com/gin-gonic/gin"
)

func main() {
	//设置系统变量,打开数据库debug
	if err := os.Setenv("DB_DEBUG", "true"); err != nil {
		panic(err)
	}

	router := gin.Default()
	global_router.InitRouter(router)

	router.Run(":8080")
}
