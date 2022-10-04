package endpoints

import (
	"fmt"
	"net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)

func (r Router)NoteRoutes(rg *gin.RouterGroup) {
    notes := rg.Group("/notes")

    notes.GET("/",getAllNotes)
    notes.POST("/", createNote)
    notes.GET("/:id", getNoteByID)
}

func getAllNotes(c *gin.Context) {
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    notes, err := models.GetAllNotes(db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch all notes: %s\n", err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": msg,})
        return
    }
    c.JSON(http.StatusOK, notes)
}

type CreateNoteReq struct {
    Content string `json:"content"`
    Stid int `json:"stid"`
}

func createNote(c *gin.Context) {
    var req CreateNoteReq
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    err = c.BindJSON(&req)
    if err != nil {
        msg := fmt.Sprintf("Failed to parse json request: %s\n", err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
        return
    }
    err = models.CreateNote(req.Content, req.Stid, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to create new Note: %s\n", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": msg})
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "created",
    })
}

func getNoteByID(c *gin.Context) {
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
    note, err := models.GetNoteByID(id, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch Note with ID %d\n%s\n", id, err.Error())
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
    }
    c.JSON(http.StatusOK, note)
}
