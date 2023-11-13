package handlers

import (
	"strconv"

	"github.com/Natan5533/library-api-go/core/ports"
	"github.com/gin-gonic/gin"
)

type LibraryHandler struct {
	libraryService ports.LibrariesService
}

type LibraryCreateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func NewHandler(libraryService ports.LibrariesService) *LibraryHandler {
	return &LibraryHandler{
		libraryService: libraryService,
	}
}

// Actions
func (handler LibraryHandler) Create(ctx *gin.Context) {
	var request LibraryCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
	}

	id, err := handler.libraryService.Create(request.Name, request.Address)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"id": id,
	})

}

func (handler LibraryHandler) GetById(ctx *gin.Context) {
	id := ctx.Param("id")

	intId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	library, err := handler.libraryService.GetById(intId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"library": library})

}
