package explorer

import "github.com/gin-gonic/gin"

func ExplorerRoutes(router *gin.RouterGroup) {
	router.GET("/list", List)
}
