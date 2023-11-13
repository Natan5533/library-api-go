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
	}

	if err := engine.Run(); err != nil {
		slog.Error("Error on run server")
		panic(err)
	}
}

func library(engine *gin.RouterGroup, lh handlers.LibraryHandler) {
	engine.POST("/library", lh.Create)
	engine.GET("/library/:id", lh.GetById)

}
