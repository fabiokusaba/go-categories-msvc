package main

import (
	"github.com/fabiokusaba/go-categories-msvc/cmd/api/controllers"
	"github.com/fabiokusaba/go-categories-msvc/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	inMemoryCategoryRepository := repositories.NewInMemoryCategoryRepository()

	categoryRoutes.POST("", func(ctx *gin.Context) {
		controllers.CreateCategory(ctx, inMemoryCategoryRepository)
	})
}
