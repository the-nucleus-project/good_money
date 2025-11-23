package main

import (
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	// Example: Add
	fmt.Println("=== Add ===")
	m1, _ := goodmoney.New(100.50, goodmoney.USD)
	m2, _ := goodmoney.New(50.25, goodmoney.USD)
	m3, _ := goodmoney.New(25.75, goodmoney.USD)

	sum, err := goodmoney.Add(m1, m2, m3)
	if err != nil {
		log.Fatalf("Error adding: %v", err)
	}
	fmt.Printf("%s + %s + %s = %s\n", m1, m2, m3, sum)

	// Example: Subtract
	fmt.Println("\n=== Subtract ===")
	diff, err := m1.Subtract(m2)
	if err != nil {
		log.Fatalf("Error subtracting: %v", err)
	}
	fmt.Printf("%s - %s = %s\n", m1, m2, diff)

	// Subtract multiple amounts
	diff2, err := m1.Subtract(m2, m3)
	if err != nil {
		log.Fatalf("Error subtracting: %v", err)
	}
	fmt.Printf("%s - %s - %s = %s\n", m1, m2, m3, diff2)

	// Example: Multiply
	fmt.Println("\n=== Multiply ===")
	product, err := m1.Multiply(3)
	if err != nil {
		log.Fatalf("Error multiplying: %v", err)
	}
	fmt.Printf("%s * 3 = %s\n", m1, product)

	// Multiply by multiple factors
	product2, err := m1.Multiply(2, 3)
	if err != nil {
		log.Fatalf("Error multiplying: %v", err)
	}
	fmt.Printf("%s * 2 * 3 = %s\n", m1, product2)

	// Example: Divide
	fmt.Println("\n=== Divide ===")
	quotient, err := m1.Divide(2)
	if err != nil {
		log.Fatalf("Error dividing: %v", err)
	}
	fmt.Printf("%s / 2 = %s\n", m1, quotient)

	// Divide by multiple divisors
	quotient2, err := m1.Divide(2, 2)
	if err != nil {
		log.Fatalf("Error dividing: %v", err)
	}
	fmt.Printf("%s / 2 / 2 = %s\n", m1, quotient2)

	// Example: Error handling - currency mismatch
	fmt.Println("\n=== Error Handling ===")
	eur, _ := goodmoney.New(100.50, goodmoney.EUR)
	_, err = goodmoney.Add(m1, eur)
	if err != nil {
		fmt.Printf("Expected error (currency mismatch): %v\n", err)
	}

	// Example: Error handling - division by zero
	_, err = m1.Divide(0)
	if err != nil {
		fmt.Printf("Expected error (division by zero): %v\n", err)
	}
}
