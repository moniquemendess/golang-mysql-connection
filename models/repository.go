package models

//interface generica
type Repository[T any] interface {
	Create(entity T) (T, error)
	Update(entity T) error
	Delete(id int) error
}
