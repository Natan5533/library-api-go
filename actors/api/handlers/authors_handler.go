package handlers

import (
	"strconv"

	adapters "github.com/Natan5533/library-api-go/core/domain/adpaters"
	"github.com/Natan5533/library-api-go/core/ports"
	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	authorService ports.AuthorsService
}

type AuthorCreateRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewAuthorsHandler(authorService ports.AuthorsService) *AuthorHandler {
	return &AuthorHandler{
		authorService: authorService,
	}
}

func (handler AuthorHandler) CreateAuthor(ctx *gin.Context) {
	id := ctx.Param("library_id")
	libraryId, _ := strconv.Atoi(id)

	var request AuthorCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(500, gin.H{"error": "Internal Server Error"})
	}

	responseId, err := handler.authorService.Create(request.Email, request.Name, libraryId)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"id": responseId,
	})
}

func (handler AuthorHandler) GetById(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	author, err := handler.authorService.GetById(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"author": author,
	})
}

func (handler AuthorHandler) Delete(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = handler.authorService.Delete(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(204, nil)
}

func (handler AuthorHandler) Update(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}

	var params adapters.UpdateAuthorParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error()})
		return
	}

	err = handler.authorService.Update(id, &params)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(204, nil)
}
