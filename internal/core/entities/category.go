package entities

type Category struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"unique;not null"`
	Logo        string `gorm:"not null"`
	Description string
}
