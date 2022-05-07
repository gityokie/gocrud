package store

import "github.com/gityokie/gocrud/internal/models"

type Store interface {
	DeleteBook(b *models.Book, id string) (err error)
	PutOneBook(b *models.Book, id string) (err error)
	GetOneBook(b *models.Book, id string) (err error)
	AddNewBook(b *models.Book) (err error)
	GetAllBook(b *[]models.Book) (err error)
}
