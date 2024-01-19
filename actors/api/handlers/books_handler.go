package handlers

import (
	"net/http"
	"strconv"

	"github.com/Natan5533/library-api-go/core/domain/adapters"
	"github.com/Natan5533/library-api-go/core/ports"
	"github.com/gin-gonic/gin"
)

type (
	BooksHandler struct {
		booksService ports.BooksService
	}

	CreateBookRequest struct {
		Title string `json:"title"`
	}
)

func NewBooksHandler(booksService ports.BooksService) *BooksHandler {
	return &BooksHandler{booksService: booksService}
}

func (handler BooksHandler) CreateBook(ctx *gin.Context) {
	id := ctx.Param("author_id")
	authorID, _ := strconv.Atoi(id)

	var request CreateBookRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	bookID, err := handler.booksService.Create(request.Title, authorID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": bookID})
}

func (handler BooksHandler) GetBookById(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	book, err := handler.booksService.GetById(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"book": book})
}

func (handler BooksHandler) Delete(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.booksService.Delete(id)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (handler BooksHandler) UpdateBook(ctx *gin.Context) {
	id, err := GetId(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var params adapters.UpdateBookParams
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = handler.booksService.Update(id, &params)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
