package endpoints

import (
	"fmt"
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)

func (r Router)SubtopicRoutes(rg *gin.RouterGroup) {
    subtopics := rg.Group("/subtopics")

    subtopics.GET("/", getAllSubtopics)
    subtopics.POST("/", createSubtopic)
    subtopics.GET("/:id", subtopicByID)
}

type SubtopicCreateReq struct {
    Name string `json:"name"`
    Toid int `json:"toid"`

}

func createSubtopic(c *gin.Context) {
    var req SubtopicCreateReq
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
    err = models.CreateSubtopic(req.Name, req.Toid, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to create Subtopic %s\n%s", req.Name, err.Error())
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "created",
    })
}

func getAllSubtopics(c *gin.Context) {
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    subtopics, err := models.GetAllSubtopics(db)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "errors": err.Error(),
        })
        return
    }
    c.JSON(http.StatusOK, subtopics)
}

func subtopicByID(c *gin.Context) {
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
    subtopic, err := models.SubtopicByID(id, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch Subtopic with id %d\n%s", id, err.Error())
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    c.JSON(http.StatusOK, subtopic)
}
