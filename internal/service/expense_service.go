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

func New(store storage.Store) *ExpenseService {
	return &ExpenseService{store: store}
}

// Return the biggest ID + 1
func nextID(expenses []models.Expense) int {
	max := 0
	for _, e := range expenses {
		if e.ID > max {
			max = e.ID
		}
	}
	return max + 1
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

func (s *ExpenseService) Delete(id int) error {
	expenses, err := s.store.Load()
	if err != nil {
		return err
	}

	index := -1
	for i, e := range expenses {
		if e.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("expense not found")
	}

	expenses = append(expenses[:index], expenses[index+1:]...)
	return s.store.Save(expenses)
}

// Return the total spend for the month given
// If month == 0, return the total spend (all months added)
func (s *ExpenseService) Summary(month int) (float64, error) {
	expenses, err := s.store.Load()

	if err != nil {
		return 0, err
	}

	total := 0.0
	for _, e := range expenses {
		if month == 0 || int(e.Date.Month()) == month {
			total += e.Amount
		}
	}
	return total, nil
}

// Return the tab of expenses
func (s *ExpenseService) List() ([]models.Expense, error) {
	return s.store.Load()
}
