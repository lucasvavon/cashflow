package sql

import (
	"cashflow-go/internal/core/entities"
	"gorm.io/gorm"
)

type GormFrequencyRepository struct {
	db *gorm.DB
}

func NewGormFrequencyRepository(db *gorm.DB) *GormFrequencyRepository {
	return &GormFrequencyRepository{db}
}

func (r *GormFrequencyRepository) FindAllFrequencies() (*entities.Frequencies, error) {
	var frequencies entities.Frequencies

	req := r.db.Find(&frequencies)
	if req.Error != nil {
		return nil, req.Error
	}

	return &frequencies, nil
}
