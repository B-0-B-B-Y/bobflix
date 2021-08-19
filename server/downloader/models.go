package downloader

type MagnetDownloadBody struct {
	MagnetLink string `json:"magnetLink" binding:"required"`
}
