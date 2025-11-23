package main

import (
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	// Example: Value (for database storage)
	fmt.Println("=== Value ===")
	m, _ := goodmoney.New(100.50, goodmoney.USD)
	value, err := m.Value()
	if err != nil {
		log.Fatalf("Error getting value: %v", err)
	}
	fmt.Printf("Money: %s\n", m)
	fmt.Printf("Database value type: %T\n", value)
	if bytes, ok := value.([]byte); ok {
		fmt.Printf("Database value (JSON): %s\n", string(bytes))
	}

	// Example: Scan (for reading from database)
	fmt.Println("\n=== Scan ===")
	var scanned goodmoney.Money

	// Simulate reading from database as []byte
	jsonBytes := []byte(`{"amount":200.75,"currency":"EUR"}`)
	err = scanned.Scan(jsonBytes)
	if err != nil {
		log.Fatalf("Error scanning: %v", err)
	}
	fmt.Printf("Scanned from []byte: %s\n", &scanned)

	// Simulate reading from database as string
	var scanned2 goodmoney.Money
	jsonStr := `{"amount":50.25,"currency":"GBP"}`
	err = scanned2.Scan(jsonStr)
	if err != nil {
		log.Fatalf("Error scanning: %v", err)
	}
	fmt.Printf("Scanned from string: %s\n", &scanned2)

	// Example: Scan with nil (NULL value)
	fmt.Println("\n=== Scan with nil ===")
	var nullMoney goodmoney.Money
	err = nullMoney.Scan(nil)
	if err != nil {
		log.Fatalf("Error scanning nil: %v", err)
	}
	fmt.Printf("Scanned nil: %s\n", &nullMoney)
	fmt.Printf("Is valid: %v\n", nullMoney.IsValid())

	// Example: Round trip (Value -> Scan)
	fmt.Println("\n=== Round Trip (Value -> Scan) ===")
	original, _ := goodmoney.New(123.45, goodmoney.JPY)
	dbValue, _ := original.Value()
	fmt.Printf("Original: %s\n", original)

	var restored goodmoney.Money
	restored.Scan(dbValue)
	fmt.Printf("Restored: %s\n", &restored)

	equal, _ := original.Equals(&restored)
	fmt.Printf("Are equal: %v\n", equal)

	// Example: Error handling
	fmt.Println("\n=== Error Handling ===")
	var invalid goodmoney.Money
	invalidValue := 12345 // Wrong type
	err = invalid.Scan(invalidValue)
	if err != nil {
		fmt.Printf("Expected error (wrong type): %v\n", err)
	}

	// Example: Using with sql package (conceptual)
	fmt.Println("\n=== Database Usage Pattern ===")
	fmt.Println("// In your database code:")
	fmt.Println("//")
	fmt.Println("// Storing:")
	fmt.Println("//   money, _ := goodmoney.New(100.50, goodmoney.USD)")
	fmt.Println("//   dbValue, _ := money.Value()")
	fmt.Println("//   db.Exec(\"INSERT INTO transactions (amount) VALUES (?)\", dbValue)")
	fmt.Println("//")
	fmt.Println("// Reading:")
	fmt.Println("//   var money goodmoney.Money")
	fmt.Println("//   var dbValue []byte")
	fmt.Println("//   db.QueryRow(\"SELECT amount FROM transactions WHERE id = ?\", id).Scan(&dbValue)")
	fmt.Println("//   money.Scan(dbValue)")
}

