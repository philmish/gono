package endpoints

import (
	"github.com/gin-gonic/gin"
)

func (r Router)PingRoutes(rg *gin.RouterGroup) {
    ping := rg.Group("/ping")

    ping.GET("/", pingAPI)
}

func pingAPI(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "Ok",
    })
}
