package controllers

import (
	"net/http"

	"github.com/fabiokusaba/go-categories-msvc/internal/repositories"
	usecases "github.com/fabiokusaba/go-categories-msvc/internal/use-cases"
	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := usecases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"categories": categories})
}
