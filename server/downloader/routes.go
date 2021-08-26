package downloader

import "github.com/gin-gonic/gin"

func DownloaderRoutes(router *gin.RouterGroup) {
	router.POST("/magnet", magnetDownload)
	router.POST("/torrent", torrentDownload)
	router.POST("/request", request)
}
