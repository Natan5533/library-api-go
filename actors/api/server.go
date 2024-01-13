package api

import (
	"log/slog"

	"github.com/Natan5533/library-api-go/actors/api/handlers"
	"github.com/Natan5533/library-api-go/infra"
	"github.com/gin-gonic/gin"
)

func InitServer(container *infra.Container) {
	engine := gin.Default()

	apiV1 := engine.Group("/api/v1")
	{
		library(apiV1, container.LibraryHandler)
		author(apiV1, container.AuthorsHandler)
	}

	if err := engine.Run(); err != nil {
		slog.Error("Error on run server")
		panic(err)
	}
}

func library(engine *gin.RouterGroup, lh handlers.LibraryHandler) {
	engine.POST("/library", lh.Create)
	engine.GET("/library/:id", lh.GetById)
	engine.DELETE("/library/:id", lh.Delete)
	engine.PUT("/library/:id", lh.Update)
}

func author(engine *gin.RouterGroup, ah handlers.AuthorHandler) {
	engine.POST("/author/:library_id", ah.CreateAuthor)
	engine.GET("/author/:id", ah.GetById)
	engine.DELETE("/author/:id", ah.Delete)
	engine.PUT("/author/:id", ah.Update)
}
