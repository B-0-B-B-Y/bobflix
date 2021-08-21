package explorer

type ListResponse struct {
	Items []ListItem `json:"items" binding:"required"`
}

type ListItem struct {
	Title       string `json:"title" binding:"required"`
	Runtime     string `json:"runtime" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	Description string `json:"description" binding:"required"`
}
