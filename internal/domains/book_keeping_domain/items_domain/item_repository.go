package itemsdomain

type ItemRepository interface {
	FindById(id int) (*Item, error)
	Create(item *Item) (*Item, error)
	Delete(item *Item) error
}
