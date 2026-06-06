package storage

import (
	"encoding/json"
	"os"

	"github.com/cduffaut/expense-tracker/internal/models"
)

// App service(s) depends on Store
type Store interface {
	Load() ([]models.Expense, error)
	Save(expenses []models.Expense) error
}

// Where is stored the data
type JsonStore struct {
	filePath string
}

// Return a type initialised
func New(filePath string) *JsonStore {
	return &JsonStore{filePath: filePath}
}

func (s *JsonStore) Load() ([]models.Expense, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []models.Expense{}, nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var expenses []models.Expense

	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}

func (s *JsonStore) Save(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
