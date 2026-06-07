package main

import (
	"github.com/cduffaut/expense-tracker/internal/cli"
	"github.com/cduffaut/expense-tracker/internal/service"
	"github.com/cduffaut/expense-tracker/internal/storage"
)

func main() {
	store := storage.New("data/expenses.json")

	svc := service.New(store)

	c := cli.New(svc)

	c.Run()
}
