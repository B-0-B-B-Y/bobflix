package downloader

import "github.com/anacrolix/torrent"

func NewTorrentClient() (*torrent.Client, error) {
	config := torrent.NewDefaultClientConfig()
	config.DataDir = "/data"

	return torrent.NewClient(config)
}
