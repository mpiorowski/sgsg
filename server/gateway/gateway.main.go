package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

    base "go-svelte-grpc/server/gateway/base"
    users "go-svelte-grpc/server/gateway/users"
)


func main() {
	if base.ENV == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{base.DOMAIN}
	config.AllowCredentials = true
    config.AllowHeaders = []string{"Authorization", "Content-Type"}
	router.Use(cors.New(config))

	router.GET("/auth", base.Auth)
	router.POST("/login", base.Login)

	router.GET("/files/:targetId", base.GetFiles)
	router.POST("/files", base.CreateFile)
	router.DELETE("/files/:targetId/:fileId", base.DeleteFile)

	router.GET("/users", users.GetUsers)
	router.POST("/users", users.CreateUser)
	router.DELETE("/users", users.DeleteUser)

	if err := router.Run(fmt.Sprintf("0.0.0.0:%v", base.PORT)); err != nil {
		panic(err)
	}
}
