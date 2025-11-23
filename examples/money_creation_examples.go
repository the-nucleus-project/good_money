package main

import (
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	// Example: New
	fmt.Println("=== New ===")
	m1, err := goodmoney.New(100.50, goodmoney.USD)
	if err != nil {
		log.Fatalf("Error creating money: %v", err)
	}
	fmt.Printf("Created: %s\n", m1)

	// Example: New with different currencies
	m2, _ := goodmoney.New(50.25, goodmoney.EUR)
	fmt.Printf("EUR: %s\n", m2)

	m3, _ := goodmoney.New(1000.0, goodmoney.JPY)
	fmt.Printf("JPY: %s\n", m3)

	// Example: NewZero
	fmt.Println("\n=== NewZero ===")
	zero, err := goodmoney.NewZero(goodmoney.USD)
	if err != nil {
		log.Fatalf("Error creating zero money: %v", err)
	}
	fmt.Printf("Zero amount: %s\n", zero)
	fmt.Printf("IsZero: %v\n", zero.IsZero())

	// Example: MustNew
	fmt.Println("\n=== MustNew ===")
	// MustNew panics if there's an error, so use only when you're certain inputs are valid
	m4 := goodmoney.MustNew(200.75, goodmoney.GBP)
	fmt.Printf("MustNew result: %s\n", m4)

	// Example: Error handling with New
	fmt.Println("\n=== Error Handling ===")
	invalid, err := goodmoney.New(100.123, goodmoney.USD) // Too many decimal places
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	} else {
		fmt.Printf("Unexpected success: %s\n", invalid)
	}

	invalidCurrency, err := goodmoney.New(100.50, "INVALID")
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	} else {
		fmt.Printf("Unexpected success: %s\n", invalidCurrency)
	}
}
