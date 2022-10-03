package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/philmish/gono/database"
	"github.com/philmish/gono/database/models"
)

func (r Router)ProjectsRoutes(rg *gin.RouterGroup) {
    projects := rg.Group("/projects")

    projects.GET("/", allProjects)

}

func allProjects(c *gin.Context) {
    db, err := database.GetDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    projects, err := models.GetAll(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, projects)
}

