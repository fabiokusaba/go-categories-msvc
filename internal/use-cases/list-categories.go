package usecases

import (
	"github.com/fabiokusaba/go-categories-msvc/internal/entities"
	"github.com/fabiokusaba/go-categories-msvc/internal/repositories"
)

type listCategoriesUseCase struct {
	repository repositories.ICategoryRepository
}

func NewListCategoriesUseCase(repository repositories.ICategoryRepository) *listCategoriesUseCase {
	return &listCategoriesUseCase{repository: repository}
}

func (u *listCategoriesUseCase) Execute() ([]*entities.Category, error) {
	categories, err := u.repository.List()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
