package endpoints

import "github.com/gin-gonic/gin"

type Router struct {
    engine *gin.Engine
}

func NewRouter() Router {
    r := Router{
        engine: gin.Default(),
    }

    api := r.engine.Group("/api")
    r.PingRoutes(api)
    r.ProjectsRoutes(api)

    return r
}

func (r Router)Run(addr string) error {
    return r.engine.Run(addr)
}
