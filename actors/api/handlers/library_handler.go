package handlers

import (
	"strconv"

	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
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
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	library, err := handler.libraryService.GetById(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{"library": library})

}

func (handler LibraryHandler) Delete(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	err = handler.libraryService.Delete(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(204, nil)

}

func (handler LibraryHandler) Update(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	var params adapters.UpdateLibraryParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
	}

	err = handler.libraryService.Update(id, &params)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(204, nil)
}

func GetId(ctx *gin.Context) (int, error) {
	id := ctx.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return intId, nil
}
