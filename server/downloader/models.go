package downloader

import "github.com/B-0-B-B-Y/bobflix/db"

type MagnetDownloadBody struct {
	MagnetLink  string `json:"magnetLink" binding:"required"`
	Runtime     string `json:"runtime" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TorrentDownloadBody struct {
	Runtime     string `json:"runtime" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type RequestBody struct {
	RequestedMovie db.MovieRequest `json:"requestedMovie" binding:"required"`
}
