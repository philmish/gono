package endpoints

import (
    "fmt"
    "net/http"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)


func (r Router)TopicRoutes(rg *gin.RouterGroup) {
    topics := rg.Group("/topics")

    topics.POST("/", createTopic)
}

type CreateTopicReq struct {
    Name string `json:"name"`
    Description string `json:"description"`
    Prid int `json:"prid"`
}

func createTopic(c *gin.Context) {
    var req CreateTopicReq
    err := c.BindJSON(&req)
    if err != nil {
        msg := fmt.Sprintf("Failed to parse json request: %s\n", err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
        return
    }
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    err = models.CreateTopic(req.Name, req.Description, req.Prid, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to create Topic %s: %s\n", req.Name, err.Error())
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "created",
    })
}
