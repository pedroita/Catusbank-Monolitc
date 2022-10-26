package model

type Repository[T any] interface {
	FindAll() []T
	Find(id int) T
	Delete(id int) int
	Update(t T) int
	Create(t T) T
}
