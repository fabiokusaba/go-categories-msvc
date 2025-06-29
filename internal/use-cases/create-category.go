package usecases

import (
	"log"

	"github.com/fabiokusaba/go-categories-msvc/internal/entities"
	"github.com/fabiokusaba/go-categories-msvc/internal/repositories"
)

type createCategoryUseCase struct {
	// db
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository: repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	log.Println(category)

	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
