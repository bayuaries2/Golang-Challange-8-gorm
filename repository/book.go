package repository

import (
	"Challange-7/model"
	"time"
)

// handler>service>repo

// interface book
type BookRepo interface {
	CreateBook(in model.Book) (res model.Book, err error)
	GetBooks() (res []model.Book, err error)
	GetBookById(id int64) (res model.Book, err error)
	UpdateBook(in model.Book) (res model.Book, err error)
	DeleteBook(id int64) (err error)
}

func (r Repo) CreateBook(in model.Book) (res model.Book, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}
	// tidak error
	return res, nil
}

func (r Repo) GetBooks() (res []model.Book, err error) {
	err = r.gorm.Where("deleted_at is null").Find(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBookById(id int64) (res model.Book, err error) {

	err = r.gorm.Where("deleted_at is null").First(&res, "id = ?", id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(in model.Book) (res model.Book, err error) {
	// r.gorm.Model(&res).Where("id = ?", in.ID).Update()
	// update 1 column jika update tampa s

	err = r.gorm.Model(&res).Where("id = ?", in.ID).Updates(model.Book{
		NameBook:    in.NameBook,
		Author:      in.Author,
		Description: in.Description,
	}).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(id int64) (err error) {
	book := model.Book{}
	// hard deleted
	// err = r.gorm.Where("id = ?", id).Delete(book).Error

	err = r.gorm.Model(&book).Where("id = ?", id).Where("deleted_at is null").Update("deleted_at", time.Now()).Error // soft deleted
	if err != nil {
		return err
	}

	return nil
}
