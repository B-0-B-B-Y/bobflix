package main

import (
	"github.com/B-0-B-B-Y/bobflix/downloader"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	downloader.DownloaderRoutes(v1.Group("/download"))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Pong",
		})
	})

	router.Run()
}
