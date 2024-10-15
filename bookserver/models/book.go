package models

type Book struct {
	ID   string   `json:"id"`
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