package downloader

import "github.com/anacrolix/torrent"

type MagnetDownloadBody struct {
	MagnetLink string `json:"magnetLink" binding:"required"`
}

type RequestBody struct {
	Items []RequestItem `json:"items" binding:"required"`
}

type RequestItem struct {
	Title       string `json:"title" binding:"required"`
	Category    string `json:"category" binding:"required"`
	Description string `json:"description"`
}

func NewTorrentClient() (*torrent.Client, error) {
	config := torrent.NewDefaultClientConfig()
	config.DataDir = "/data"

	return torrent.NewClient(config)
}
