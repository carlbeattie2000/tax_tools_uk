package customer

type CustomerRepository interface {
	FindById(id int) (*Customer, error)
	Create(customer *Customer) (*Customer, error)
	Delete(customer *Customer) error
}
