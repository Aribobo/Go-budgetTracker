package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// transaction struct to hold each transaction's details
type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     time.Time
	Type     string
}

// BudgetTracker Struct to manage transaction
type BudgetTracker struct {
	transactions []Transaction
	nextID       int
}

// interface for common behaviour
// interface is a key to achieve polymorphic behaviour
type FinacialRecord interface {
	GetAmount() float64
	GetType() string
}

// implement interface method for transactions
func (t Transaction) GetAmount() float64 {
	return t.Amount
}
func (t Transaction) GetType() string {
	return t.Type

}

// add a new transaction
func (bt *BudgetTracker) AddTransaction(amount float64, category string, tType string) {
	newTransaction := Transaction{
		ID:       bt.nextID,
		Amount:   amount,
		Category: category,
		Date:     time.Now(),
		Type:     tType,
	}
	bt.transactions = append(bt.transactions, newTransaction)
	bt.nextID++
}

// Creating display transaction method
func (bt BudgetTracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")
	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n",
			transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("2006-01-02"), transaction.Type)

	}
}

//Get total income or expense

func (bt BudgetTracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}
func (bt BudgetTracker) saveToCsv(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file) // creating a new CSV file
	defer writer.Flush()          //flush is very important to make sure that data is written before the file is closed
	//write the csv header
	writer.Write([]string{"ID", "Amount", "Category", "Date", "Type"})
	// write data
	for _, t := range bt.transactions {
		record := []string{
			strconv.Itoa(t.ID),
			fmt.Sprintf("%.2f", t.Amount),
			t.Type,
		}
		writer.Write(record)
	}
	fmt.Println("Transaction saved to ", filename)
	return nil
}

// save the transaction to a csv file
func main() {
	// instantiation of budget tracker struct
	bt := BudgetTracker{}
	for {
		fmt.Println("\n-- personal Budget Tracker ---")
		fmt.Println("1. Add Transaction")
		fmt.Println("2. Display Transactions")
		fmt.Println("3. Show Total Income")
		fmt.Println("4. Show Total Expenses")
		fmt.Println("5. Save Transactions to CSV")
		fmt.Println("6. Exit")
		fmt.Println("choose an option: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter Amount : ")
			var amount float64
			fmt.Scanln(&amount)

			fmt.Print("Enter Category : ")
			var category string
			fmt.Scanln(&category)

			fmt.Print("Enter type (Income / Expense):")
			var tType string
			fmt.Scanln(&tType)

			bt.AddTransaction(amount, category, tType)
			fmt.Println("Transaction Added!")
		case 2:
			bt.DisplayTransactions()
		case 3:
			fmt.Printf("total Income: %.2f\n", bt.CalculateTotal("Income"))
		case 4:
			fmt.Printf("total Expenses: %.2f\n", bt.CalculateTotal("Expense"))
		case 5:
			fmt.Printf("Enter filename (e.g transactiond.csv)")
			var filename string
			fmt.Scanln(&filename)
			if err := bt.saveToCsv(filename); err != nil {
				fmt.Println("Enter saving transactions:", err)
			}
		case 6:
			fmt.Println("Exiting.....")
			return
		default:
			fmt.Println("Invalid Choice! Try Again!")

		}
	}
}
