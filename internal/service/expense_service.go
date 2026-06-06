package service

import (
	"errors"
	"time"

	"github.com/cduffaut/expense-tracker/internal/models"
	"github.com/cduffaut/expense-tracker/internal/storage"
)

type ExpenseService struct {
	store storage.Store
}

// Add a new expense to the total
func (s *ExpenseService) Add(description string, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if description == "" {
		return errors.New("description cannot be empty")
	}

	expenses, err := s.store.Load()

	if err != nil {
		return err
	}

	newExpense := models.Expense{
		ID:          nextID(expenses),
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)
	return s.store.Save(expenses)
}
