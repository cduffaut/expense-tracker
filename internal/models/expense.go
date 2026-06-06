package models

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
