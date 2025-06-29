package repositories

import "github.com/fabiokusaba/go-categories-msvc/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
}
