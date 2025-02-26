package entities

type Frequency struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"unique;not null"`
}

type Frequencies []Frequency
