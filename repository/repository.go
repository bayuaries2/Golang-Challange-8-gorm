package repository

import (
	"gorm.io/gorm"
)

type Repo struct {
	gorm *gorm.DB
}

type RepoInterface interface {
	BookRepo
}

// make contructor function

func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm} // handel dependency
}
