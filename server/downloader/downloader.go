package downloader

import (
	"github.com/anacrolix/torrent"
	"github.com/gin-gonic/gin"
)

func DownloaderRoutes(router *gin.RouterGroup) {
	router.POST("/magnet", magnetDownload)
	router.POST("/torrent", torrentDownload)
	router.POST("/request", request)
}

func magnetDownload(c *gin.Context) {
	var body MagnetDownloadBody
	c.Bind(&body)
	magnetLink := body.MagnetLink

	if magnetLink == "" {
		c.JSON(400, gin.H{
			"Error": "You need to specify a magnet link",
		})
	}

	config := torrent.NewDefaultClientConfig()
	config.DataDir = "/data"

	client, _ := torrent.NewClient(config)
	defer client.Close()

	torrent, err := client.AddMagnet(magnetLink)
	if err != nil {
		c.AbortWithError(500, err)
	}

	<-torrent.GotInfo()
	torrent.DownloadAll()
	client.WaitAll()

	c.JSON(200, gin.H{
		"name":   torrent.Info().Name,
		"status": "download complete",
	})
}

func torrentDownload(c *gin.Context) {}

func request(c *gin.Context) {}
