package main

import (
	"github.com/RGRU/go-memorycache"
	"github.com/gin-gonic/gin"
	"github.com/test-cache/internal/api"
	"github.com/test-cache/internal/api/controllers"
	"net/http"
	"time"
)

func main() {
	cStorage := memorycache.New(5*time.Minute, 5*time.Second)

	gin.SetMode("debug")

	router := gin.Default()
	setupRoutes(router, cStorage)
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":9000")
}

// setupRoutes роутер.
func setupRoutes(router *gin.Engine, cStorage *memorycache.Cache) {
	rg := router.Group("", api.ErrorHandler())
	{
		rg.GET("/", api.Handler(controllers.Cache(cStorage)))
	}
}
