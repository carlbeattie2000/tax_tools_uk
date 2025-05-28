package itemsdomain

type CategoryRepository interface {
	FindById(id int) (*Category, error)
	Create(category *Category) (*Category, error)
	Delete(category *Category) error
}
