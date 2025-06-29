package controllers

import (
	"net/http"

	"github.com/fabiokusaba/go-categories-msvc/internal/repositories"
	usecases "github.com/fabiokusaba/go-categories-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	useCase := usecases.NewCreateCategoryUseCase(repository)

	err := useCase.Execute(body.Name)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"success": true})
}
