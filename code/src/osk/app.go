package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"osk/auth"
	"osk/core"
	"osk/model"
)

func main() {
	core.SetupEnv()
	core.SetupConfig()
	core.SetupLogger()

	core.Logger.Info("initialize model")
	model.Initialize()

	if "release" == core.Config.GinMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// 文件服务
	router.Use(static.Serve("/files", static.LocalFile("./files", false)))

	// 跨域调用
	corsConf := cors.Config{
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"content-type", "access-control-allow-headers", "origin", "authorization", "access-control-allow-origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	corsConf.AllowAllOrigins = true
	router.Use(cors.New(corsConf))

	// 需要认证
	authGroup := auth.BindAuthHandler(router, "/signin", "/auth")

	route(authGroup)

	core.Logger.Infof("serve at %v", core.Config.Host)
	router.Run(core.Config.Host)
}
