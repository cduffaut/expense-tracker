package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/cduffaut/expense-tracker/internal/service"
)

type CLI struct {
	service *service.ExpenseService
}

func New(service *service.ExpenseService) *CLI {
	return &CLI{service: service}
}

func (c *CLI) handleAdd() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	description := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0, "Expense amount")

	addCmd.Parse(os.Args[2:])

	if *description == "" {
		fmt.Println("Error: --description cannot be empty")
		os.Exit(1)
	}

	if *amount <= 0 {
		fmt.Println("Error: --amount must be greater than zero")
		os.Exit(1)
	}

	if err := c.service.Add(*description, *amount); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Info: Expense added successfully!")
}

func (c *CLI) handleList() {
	expenses, err := c.service.List()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if len(expenses) == 0 {
		fmt.Println("No expenses found")
		return
	}

	fmt.Printf("%-4s %-12s %-20s %-10s\n", "ID", "Date", "Description", "Amount")
	for _, e := range expenses {
		fmt.Printf("%-4d %-12s %-20s $%-10.2f\n",
			e.ID,
			e.Date.Format("2006-01-02"),
			e.Description,
			e.Amount,
		)
	}
}

func (c *CLI) handleDelete() {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "Expense ID to delete")
	deleteCmd.Parse(os.Args[2:])

	if *id <= 0 {
		fmt.Println("Error: ID must be greater than zero")
		os.Exit(1)
	}

	if err := c.service.Delete(*id); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Info: Expense %d has been successfully deleted\n", *id)
}

func monthName(month int) string {
	months := []string{
		"January", "February", "March", "April",
		"May", "June", "July", "August",
		"September", "October", "November", "December",
	}

	if month < 1 || month > 12 {
		return "Unknown"
	}

	return months[month-1]
}

func (c *CLI) handleSummary() {
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.Int("month", 0, "Month number (1-12)")
	summaryCmd.Parse(os.Args[2:])

	total, err := c.service.Summary(*month)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if *month == 0 {
		fmt.Printf("Total expense: $%.2f\n", total)
	} else {
		fmt.Printf("Total expenses for %s: $%.2f\n", monthName(*month), total)
	}
}

func (c *CLI) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: expense-tracker <command> [option]")
		fmt.Println("Commands: add, list, delete, summary")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		c.handleAdd()
	case "list":
		c.handleList()
	case "delete":
		c.handleDelete()
	case "summary":
		c.handleSummary()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
