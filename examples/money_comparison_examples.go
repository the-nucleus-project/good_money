package main

import (
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	m1, _ := goodmoney.New(100.50, goodmoney.USD)
	m2, _ := goodmoney.New(50.25, goodmoney.USD)
	m3, _ := goodmoney.New(100.50, goodmoney.USD)

	// Example: Compare
	fmt.Println("=== Compare ===")
	result, err := m1.Compare(m2)
	if err != nil {
		log.Fatalf("Error comparing: %v", err)
	}
	if result > 0 {
		fmt.Printf("%s is greater than %s\n", m1, m2)
	} else if result < 0 {
		fmt.Printf("%s is less than %s\n", m1, m2)
	} else {
		fmt.Printf("%s is equal to %s\n", m1, m2)
	}

	result2, _ := m1.Compare(m3)
	if result2 == 0 {
		fmt.Printf("%s is equal to %s\n", m1, m3)
	}

	// Example: Equals
	fmt.Println("\n=== Equals ===")
	equal, err := m1.Equals(m2)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%s equals %s: %v\n", m1, m2, equal)

	equal2, _ := m1.Equals(m3)
	fmt.Printf("%s equals %s: %v\n", m1, m3, equal2)

	// Example: GreaterThan
	fmt.Println("\n=== GreaterThan ===")
	greater, err := m1.GreaterThan(m2)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%s is greater than %s: %v\n", m1, m2, greater)

	greater2, _ := m1.GreaterThan(m3)
	fmt.Printf("%s is greater than %s: %v\n", m1, m3, greater2)

	// Example: GreaterThanOrEqual
	fmt.Println("\n=== GreaterThanOrEqual ===")
	greaterOrEqual, _ := m1.GreaterThanOrEqual(m2)
	fmt.Printf("%s >= %s: %v\n", m1, m2, greaterOrEqual)

	greaterOrEqual2, _ := m1.GreaterThanOrEqual(m3)
	fmt.Printf("%s >= %s: %v\n", m1, m3, greaterOrEqual2)

	// Example: LessThan
	fmt.Println("\n=== LessThan ===")
	less, _ := m2.LessThan(m1)
	fmt.Printf("%s is less than %s: %v\n", m2, m1, less)

	less2, _ := m1.LessThan(m3)
	fmt.Printf("%s is less than %s: %v\n", m1, m3, less2)

	// Example: LessThanOrEqual
	fmt.Println("\n=== LessThanOrEqual ===")
	lessOrEqual, _ := m2.LessThanOrEqual(m1)
	fmt.Printf("%s <= %s: %v\n", m2, m1, lessOrEqual)

	lessOrEqual2, _ := m1.LessThanOrEqual(m3)
	fmt.Printf("%s <= %s: %v\n", m1, m3, lessOrEqual2)

	// Example: Error handling - currency mismatch
	fmt.Println("\n=== Error Handling ===")
	eur, _ := goodmoney.New(100.50, goodmoney.EUR)
	_, err = m1.Compare(eur)
	if err != nil {
		fmt.Printf("Expected error (currency mismatch): %v\n", err)
	}
}
