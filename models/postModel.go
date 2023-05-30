package models

type Post struct {
	// gorm.Model definition

	ID uint `gorm:"primaryKey"`

	Title string
	Body  string
}
