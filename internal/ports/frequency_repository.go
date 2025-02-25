package ports

import "cashflow-go/internal/core/entities"

type FrequencyRepository interface {
	FindAllFrequencies() (entities.Frequencies, error)
}
