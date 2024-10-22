package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     string `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func (b Book) GetID() string {
	return b.ID
}

func (b *Book) SetID(id string) {
	b.ID = id
}