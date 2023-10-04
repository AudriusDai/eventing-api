package route

import "github.com/gin-gonic/gin"

func SetupRoutes(group *gin.RouterGroup) {
	postEvent(group)
}
