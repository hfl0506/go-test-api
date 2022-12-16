package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id uuid.UUID `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
	Rating int `json:"rating"`
}

func (b *Book) BeforeSave(tx *gorm.DB) (err error) {
	id, err := uuid.NewV4()

	if err != nil {
		return
	}

	b.Id = id

	return
}