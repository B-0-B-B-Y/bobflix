package downloader

import (
	"path/filepath"

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

	client, err := NewTorrentClient()
	if err != nil {
		c.AbortWithError(500, err)
	}
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

func torrentDownload(c *gin.Context) {
	file, err := c.FormFile("torrent")
	if err != nil {
		c.AbortWithError(400, err)
	}

	filename := filepath.Base(file.Filename)
	client, err := NewTorrentClient()
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer client.Close()

	torrent, err := client.AddTorrentFromFile(filename)
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

func request(c *gin.Context) {
	var body RequestBody
	c.Bind(&body)
	items := body.Items

	for _, item := range items {
		// WIP : Write these to a store or db, perhaps redis or an sqlite db
	}
}
