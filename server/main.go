package main

import (
	"github.com/gin-gonic/gin"

	// local import
	"server/auth"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/ping", ping)

	router.POST("/api/test", auth.TokenAuthMiddleware(), auth.Test)
	router.POST("/api/login", auth.TokenAuthMiddleware(), auth.Login)
	router.POST("/api/logout", auth.TokenAuthMiddleware(), auth.Logout)
	// post the refresh token to refresh the access token
	router.POST("/api/refresh", auth.Refresh)

	router.Run("localhost:3040")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
