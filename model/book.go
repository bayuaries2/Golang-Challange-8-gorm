package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Book struct {
	ID          int       `json:"id" gorm:"primaryKey;type:serial" valid:"requires"`
	NameBook    string    `json:"name_book" db:"name_book" gorm:"type:varchar(50); not null" valid:"requires"`
	Author      string    `json:"author" db:"author" gorm:"type:varchar(20); not null" valid:"requires"`
	Description string    `json:"description" db:"description" valid:"requires"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"-" gorm:"default:null"`
}

func (b Book) Validation() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.NameBook, validation.Required),
		validation.Field(&b.Author, validation.Required),
		validation.Field(&b.Description, validation.Required))
}
