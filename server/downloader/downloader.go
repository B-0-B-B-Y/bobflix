package downloader

import (
	"path/filepath"

	database "github.com/B-0-B-B-Y/bobflix/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

	database.New().AddMovie(&database.Movie{
		ID:          uuid.NewString(),
		Title:       torrent.Info().Name,
		Runtime:     body.Runtime,
		Genre:       body.Genre,
		Description: body.Description,
	})

	c.JSON(200, gin.H{
		"name":   torrent.Info().Name,
		"status": "download complete",
	})
}

func torrentDownload(c *gin.Context) {
	var body TorrentDownloadBody
	c.Bind(&body)
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

	database.New().AddMovie(&database.Movie{
		ID:          uuid.NewString(),
		Title:       torrent.Info().Name,
		Runtime:     body.Runtime,
		Genre:       body.Genre,
		Description: body.Description,
	})

	c.JSON(200, gin.H{
		"name":   torrent.Info().Name,
		"status": "download complete",
	})
}

func request(c *gin.Context) {
	var body RequestBody
	c.Bind(&body)

	err := database.New().AddMovieRequest(&body.RequestedMovie)
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.Status(200)
}
