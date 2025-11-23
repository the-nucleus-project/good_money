package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	// Example: MarshalJSON
	fmt.Println("=== MarshalJSON ===")
	m, _ := goodmoney.New(100.50, goodmoney.USD)
	jsonBytes, err := m.MarshalJSON()
	if err != nil {
		log.Fatalf("Error marshaling: %v", err)
	}
	fmt.Printf("JSON: %s\n", string(jsonBytes))

	// Using standard json.Marshal (which calls MarshalJSON)
	jsonBytes2, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error marshaling: %v", err)
	}
	fmt.Printf("Using json.Marshal: %s\n", string(jsonBytes2))

	// Example: UnmarshalJSON
	fmt.Println("\n=== UnmarshalJSON ===")
	var unmarshaled goodmoney.Money
	err = unmarshaled.UnmarshalJSON(jsonBytes)
	if err != nil {
		log.Fatalf("Error unmarshaling: %v", err)
	}
	fmt.Printf("Unmarshaled: %s\n", &unmarshaled)

	// Using standard json.Unmarshal
	var unmarshaled2 goodmoney.Money
	jsonStr := `{"amount":200.75,"currency":"EUR"}`
	err = json.Unmarshal([]byte(jsonStr), &unmarshaled2)
	if err != nil {
		log.Fatalf("Error unmarshaling: %v", err)
	}
	fmt.Printf("Unmarshaled from JSON string: %s\n", &unmarshaled2)

	// Example: Round trip
	fmt.Println("\n=== Round Trip ===")
	original, _ := goodmoney.New(123.45, goodmoney.GBP)
	jsonData, _ := original.MarshalJSON()
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("JSON: %s\n", string(jsonData))

	var restored goodmoney.Money
	restored.UnmarshalJSON(jsonData)
	fmt.Printf("Restored: %s\n", &restored)

	// Verify they're equal
	equal, _ := original.Equals(&restored)
	fmt.Printf("Are equal: %v\n", equal)

	// Example: Error handling - invalid JSON
	fmt.Println("\n=== Error Handling ===")
	var invalid goodmoney.Money
	invalidJSON := []byte(`{"amount":100.50}`) // Missing currency
	err = invalid.UnmarshalJSON(invalidJSON)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	invalidJSON2 := []byte(`{"amount":100.50,"currency":""}`) // Empty currency
	err = invalid.UnmarshalJSON(invalidJSON2)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}

	invalidJSON3 := []byte(`{"amount":100.123,"currency":"USD"}`) // Too many decimals
	err = invalid.UnmarshalJSON(invalidJSON3)
	if err != nil {
		fmt.Printf("Expected error: %v\n", err)
	}
}

