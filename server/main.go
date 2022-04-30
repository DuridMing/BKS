package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// local import
	"server/auth"
	"server/database/curd"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/ping", ping)

	router.POST("/api/test", auth.TokenAuthMiddleware(), auth.Test)
	router.POST("/api/login", auth.Login)
	router.POST("/api/logout", auth.TokenAuthMiddleware(), auth.Logout)
	// post the refresh token to refresh the access token
	router.POST("/api/refresh", auth.Refresh)

	router.POST("/api/bkms/u", curd.FindoneUser)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello Gin",
		})
	})

	router.Run("localhost:13040")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
