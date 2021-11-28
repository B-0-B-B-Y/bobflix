package explorer

import (
	"github.com/gin-gonic/gin"
)

func ExplorerRoutes(router *gin.RouterGroup) {
	router.GET("/list", List)
	router.StaticFS("/stream", gin.Dir(VIDEO_DIR, false))
}
