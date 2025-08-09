package storage

import (
	"os"

	"github.com/jmilne22/coffee-dial/internal/models"
	"gopkg.in/yaml.v3"
)

type YAMLStorage struct {
	filePath string
}

func NewYAMLStorage(filepath string) *YAMLStorage {
	return &YAMLStorage{
		filePath: filepath,
	}
}

func (y *YAMLStorage) fileExists() bool {
	if _, err := os.Stat(y.filePath); err == nil {
		return true
	}
	return false
}

func (y *YAMLStorage) readAllGrinds() ([]*models.Grind, error) {
	if !y.fileExists() {
		return []*models.Grind{}, nil
	}
	yamlFile, err := os.ReadFile(y.filePath)
	if err != nil {
		return nil, err
	}

	var grinds []*models.Grind

	err = yaml.Unmarshal(yamlFile, &grinds)
	if err != nil {
		return nil, err
	}
	return grinds, nil
}
