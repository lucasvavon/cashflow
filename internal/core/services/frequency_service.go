package services

import (
	"cashflow-go/internal/core/entities"
	"cashflow-go/internal/ports"
)

type FrequencyService struct {
	repo ports.FrequencyRepository
}

func NewFrequencyService(repo ports.FrequencyRepository) *FrequencyService {
	return &FrequencyService{
		repo: repo,
	}
}

func (fs *FrequencyService) FindAllFrequencies() (entities.Frequencies, error) {
	return fs.repo.FindAllFrequencies()
}
