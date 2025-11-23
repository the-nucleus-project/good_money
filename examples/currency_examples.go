package main

import (
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	// Example: GetCurrency
	fmt.Println("=== GetCurrency ===")
	currency := goodmoney.GetCurrency(goodmoney.USD)
	if currency != nil {
		fmt.Printf("Currency USD - NumericCode: %s, Symbol: %s, MinorUnit: %d, SymbolPosition: %v\n",
			currency.NumericCode, currency.Symbol, currency.MinorUnit, currency.SymbolPosition)
	}

	// Example: GetCurrencyByNumericCode
	fmt.Println("\n=== GetCurrencyByNumericCode ===")
	currency2, code, err := goodmoney.GetCurrencyByNumericCode("840")
	if err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Currency Code: %s, Symbol: %s\n", code, currency2.Symbol)
	}

	// Example: ValidateCurrency
	fmt.Println("\n=== ValidateCurrency ===")
	valid := goodmoney.ValidateCurrency(goodmoney.USD)
	fmt.Printf("USD is valid: %v\n", valid)

	invalid := goodmoney.ValidateCurrency("INVALID")
	fmt.Printf("INVALID is valid: %v\n", invalid)

	// Validate multiple currencies
	currencies := []string{goodmoney.USD, goodmoney.EUR, goodmoney.GBP, "XYZ"}
	for _, code := range currencies {
		isValid := goodmoney.ValidateCurrency(code)
		fmt.Printf("%s: %v\n", code, isValid)
	}
}
