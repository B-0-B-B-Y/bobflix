package main

import (
	"github.com/B-0-B-B-Y/bobflix/downloader"
	"github.com/B-0-B-B-Y/bobflix/explorer"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	downloader.DownloaderRoutes(v1.Group("/download"))
	explorer.ExplorerRoutes(v1.Group("/explore", CORSMiddleware()))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Pong",
		})
	})

	router.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
