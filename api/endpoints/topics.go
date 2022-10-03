package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)


func (r Router)TopicRoutes(rg *gin.RouterGroup) {
    topics := rg.Group("/topics")

    topics.GET("/", getAllTopics)
    topics.POST("/", createTopic)
    topics.GET("/:id", topicByID)
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

func getAllTopics(c *gin.Context) {
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    topics, err := models.GetAllTopics(db)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, topics)
}

func topicByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        msg := fmt.Sprintf("Failed to convert id to int: %s\n", c.Param("id"))
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    topic, err := models.TopicByID(id, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch Topic with id %d\n%s", id, err.Error())
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    c.JSON(http.StatusOK, topic)
}
