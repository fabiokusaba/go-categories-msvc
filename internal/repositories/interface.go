package repositories

type ICategoryRepository interface {
	Save(name string) error
}
