package storage

import "github.com/jmilne22/coffee-dial/internal/models"

type GrindStorage interface {
	Add(grind *models.Grind) error
	Get(fieldName string, value string) ([]*models.Grind, error)
	GetAll() ([]*models.Grind, error)
	Edit(fieldName string, value string, updates map[string]string) error
	Delete(fieldName string, value string) error
}
