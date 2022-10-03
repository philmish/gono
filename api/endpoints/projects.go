package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)

func (r Router)ProjectsRoutes(rg *gin.RouterGroup) {
    projects := rg.Group("/projects")

    projects.GET("/", allProjects)
    projects.POST("/", createProject)
    projects.GET("/:id", projectByID)

}

func allProjects(c *gin.Context) {
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    projects, err := models.GetAllProjects(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, projects)
}

type ProjectCreateReq struct {
    Name string `json:"name"`
}

func createProject(c *gin.Context) {
    var req ProjectCreateReq
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
    err = models.CreateProject(req.Name, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to create Project: %s\n", err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": msg})
        return
    }
    c.JSON(http.StatusCreated, gin.H{
        "status": "created",
    })
}

func projectByID(c *gin.Context) {
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
    project, err := models.ProjectByID(id, db)
    if err != nil {
        msg := fmt.Sprintf("Failed to fetch project with id %d\n", id)
        c.JSON(http.StatusBadRequest, gin.H{
            "error": msg,
        })
        return
    }
    c.JSON(http.StatusOK, project)
}

