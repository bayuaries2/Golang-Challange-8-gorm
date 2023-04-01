package service

import (
	"Challange-7/model"
)

type BookService interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBooks() (res []model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (s *Service) CreateBook(in model.Book) (res model.Book, err error) {
	// call repository
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}

	return res, nil
}
func (s *Service) GetBooks() (res []model.Book, err error) {
	return s.repo.GetBooks()
}

func (s *Service) GetBookById(id int64) (res model.Book, err error) {
	return s.repo.GetBookById(id)
}

func (s *Service) UpdateBook(in model.Book) (res model.Book, err error) {
	return s.repo.UpdateBook(in)
}

func (s *Service) DeleteBook(id int64) (err error) {
	return s.repo.DeleteBook(id)
}
