package handler

import (
	app "bingo"
	"bingo/middleware"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Initialize() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), gin.Recovery())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	staticServer := static.Serve("/", app.EmbedFolder())
	r.Use(staticServer)
	r.NoRoute(func(c *gin.Context) {
		if c.Request.Method == http.MethodGet &&
			!strings.ContainsRune(c.Request.URL.Path, '.') &&
			!strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.Request.URL.Path = "/"
			staticServer(c)
		}
	})

	api := r.Group("/api")
	{
		games := api.Group("/games")
		games.POST("", createGame)
		games.GET("/:id", getGame)
		games.PUT("/:id/call_number", callNumber)
	}

	return r
}
