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

	router.GET("/api/ping", ping)
	router.POST("/api/test", auth.TokenAuthMiddleware(), auth.Test)

	router.POST("/api/login", auth.Login)
	router.POST("/api/logout", auth.TokenAuthMiddleware(), auth.Logout)
	// post the refresh token to refresh the access token
	router.POST("/api/refresh", auth.Refresh)

	router.POST("/api/bkms/s/name", curd.FindBookbyName)
	router.POST("/api/bkms/s/author", curd.FindBookbyAuthor)
	router.POST("/api/bkms/s/isbn", curd.FindBookbyISBN)

	router.POST("/api/bkms/add", curd.AddOneBook)
	router.POST("/api/bkms/delete", curd.DeleteOneBook)
	router.POST("/api/bkms/modify", curd.ModifyOneBook)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello BKMS",
		})
	})

	router.Run("localhost:3040")
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}
