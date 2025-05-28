package financialdomain

type AccountRepository interface {
	FindById(id int) (*Account, error)
	Create(account *Account) (*Account, error)
	Delete(account *Account) error
}
