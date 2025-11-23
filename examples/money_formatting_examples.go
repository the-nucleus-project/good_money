package main

import (
	"fmt"

	"github.com/nucleus-proj/goodmoney/goodmoney"
	"golang.org/x/text/language"
)

func main() {
	m, _ := goodmoney.New(1234.56, goodmoney.USD)

	// Example: String
	fmt.Println("=== String ===")
	fmt.Printf("String representation: %s\n", m.String())
	fmt.Printf("Using fmt.Printf: %s\n", m)

	// Example: Format
	fmt.Println("\n=== Format ===")
	fmt.Printf("American English: %s\n", m.Format(language.AmericanEnglish))
	fmt.Printf("German: %s\n", m.Format(language.German))
	fmt.Printf("French: %s\n", m.Format(language.French))
	fmt.Printf("Italian: %s\n", m.Format(language.Italian))

	// Different currencies
	eur, _ := goodmoney.New(1234.56, goodmoney.EUR)
	fmt.Printf("\nEUR in German: %s\n", eur.Format(language.German))
	fmt.Printf("EUR in French: %s\n", eur.Format(language.French))

	gbp, _ := goodmoney.New(1234.56, goodmoney.GBP)
	fmt.Printf("GBP in British English: %s\n", gbp.Format(language.BritishEnglish))

	// Example: FormatWithMode - FormatStandard
	fmt.Println("\n=== FormatWithMode - FormatStandard ===")
	fmt.Printf("Standard: %s\n", m.FormatWithMode(language.AmericanEnglish, goodmoney.FormatStandard))

	// Example: FormatWithMode - FormatAccounting
	fmt.Println("\n=== FormatWithMode - FormatAccounting ===")
	positive, _ := goodmoney.New(100.50, goodmoney.USD)
	negative, _ := goodmoney.New(-100.50, goodmoney.USD)
	fmt.Printf("Positive accounting: %s\n", positive.FormatWithMode(language.AmericanEnglish, goodmoney.FormatAccounting))
	fmt.Printf("Negative accounting: %s\n", negative.FormatWithMode(language.AmericanEnglish, goodmoney.FormatAccounting))

	// Example: FormatWithMode - FormatCompact
	fmt.Println("\n=== FormatWithMode - FormatCompact ===")
	thousands, _ := goodmoney.New(1500.00, goodmoney.USD)
	millions, _ := goodmoney.New(1500000.00, goodmoney.USD)
	billions, _ := goodmoney.New(5300000000.00, goodmoney.USD)
	small, _ := goodmoney.New(500.00, goodmoney.USD)

	fmt.Printf("Small amount: %s\n", small.FormatWithMode(language.AmericanEnglish, goodmoney.FormatCompact))
	fmt.Printf("Thousands: %s\n", thousands.FormatWithMode(language.AmericanEnglish, goodmoney.FormatCompact))
	fmt.Printf("Millions: %s\n", millions.FormatWithMode(language.AmericanEnglish, goodmoney.FormatCompact))
	fmt.Printf("Billions: %s\n", billions.FormatWithMode(language.AmericanEnglish, goodmoney.FormatCompact))

	// Example: FormatWithMode - FormatMinimal
	fmt.Println("\n=== FormatWithMode - FormatMinimal ===")
	fmt.Printf("Minimal: %s\n", m.FormatWithMode(language.AmericanEnglish, goodmoney.FormatMinimal))

	// Example: FormatWithMode - FormatSymbol
	fmt.Println("\n=== FormatWithMode - FormatSymbol ===")
	fmt.Printf("Symbol only: %s\n", m.FormatWithMode(language.AmericanEnglish, goodmoney.FormatSymbol))
	fmt.Printf("EUR symbol: %s\n", eur.FormatWithMode(language.German, goodmoney.FormatSymbol))
	fmt.Printf("GBP symbol: %s\n", gbp.FormatWithMode(language.BritishEnglish, goodmoney.FormatSymbol))

	// Example: FormatWithMode - FormatCode
	fmt.Println("\n=== FormatWithMode - FormatCode ===")
	fmt.Printf("Code only: %s\n", m.FormatWithMode(language.AmericanEnglish, goodmoney.FormatCode))
	fmt.Printf("EUR code: %s\n", eur.FormatWithMode(language.German, goodmoney.FormatCode))

	// Example: FormatWithOptions
	fmt.Println("\n=== FormatWithOptions ===")
	opts1 := goodmoney.FormatOptions{
		Locale: language.German,
		Mode:   goodmoney.FormatStandard,
	}
	fmt.Printf("German standard: %s\n", m.FormatWithOptions(opts1))

	opts2 := goodmoney.FormatOptions{
		Locale: language.French,
		Mode:   goodmoney.FormatAccounting,
	}
	fmt.Printf("French accounting (negative): %s\n", negative.FormatWithOptions(opts2))

	opts3 := goodmoney.FormatOptions{
		Locale: language.AmericanEnglish,
		Mode:   goodmoney.FormatCompact,
	}
	fmt.Printf("American compact: %s\n", millions.FormatWithOptions(opts3))
}
