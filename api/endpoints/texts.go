package endpoints

import (
	"fmt"
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)

func (r Router)TextRoutes(rg *gin.RouterGroup) {
    texts := rg.Group("/texts")

    texts.GET("/", getAllTexts)
    texts.POST("/", createText)
    texts.GET("/:id", getTextByID)

}

type CreateTextReq struct {
    Title string `json:"title"`
    Content string `json:"content"`
    Stid int `json:"stid"`
}

func createText(c *gin.Context) {
    var req CreateTextReq

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
    err = models.CreateText(req.Title, req.Content, req.Stid, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to create text: %s\n", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": msg,})
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "created",
    })
}

func getAllTexts(c *gin.Context) {
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    texts, err := models.GetAllTexts(db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch all texts: %s\n", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": msg,})
        return
    }
    c.JSON(http.StatusOK, texts)
}

func getTextByID(c *gin.Context) {
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
    text, err := models.GetTextByID(id, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch Text with ID %d\n%s\n", id, err.Error())
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    c.JSON(http.StatusOK, text)
}
