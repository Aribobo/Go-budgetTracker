package main

import (
	"fmt"
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

// Budgettracker Struct to manage transaction
type Budgettracker struct {
	transactions []Transaction
	nextID       int
}

// interface for common behaviour
// interface is a key to achieve polymorphic behaviour
type FinacialRecord interface {
	GetAmount() float64
	GetType() string
}

// implement interface method for transaction
func (t Transaction) GetAmount() float64 {
	return t.Amount
}
func (t Transaction) GetType() string {
	return t.Type

}

// add a new transaction
func (bt *Budgettracker) AddTransaction(amount float64, category, tType string) {
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
func (bt Budgettracker) DisplayTransactions() {
	fmt.Println("ID\tAmount\tCategory\tDate\tType")
	for _, transaction := range bt.transactions {
		fmt.Printf("%d\t%.2f\t%s\t%s\t%s\n",
			transaction.ID, transaction.Amount, transaction.Category, transaction.Date.Format("2006-01-02"), transaction.Type)

	}
}

//Get total income or expense

func (bt Budgettracker) CalculateTotal(tType string) float64 {
	var total float64
	for _, transaction := range bt.transactions {
		if transaction.Type == tType {
			total += transaction.Amount
		}
	}
	return total
}

// save the transaction to a csv file
func main() {

}
