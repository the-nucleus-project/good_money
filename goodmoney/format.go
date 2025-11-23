package goodmoney

import (
	"fmt"
	"math"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// FormatMode defines different formatting modes for currency display
type FormatMode int

const (
	// FormatStandard formats with currency symbol in standard position
	FormatStandard FormatMode = iota
	// FormatAccounting formats negative amounts with parentheses: (100.50)
	FormatAccounting
	// FormatCompact formats large amounts in compact notation: $1.5K, $1.2M, $5.3B
	FormatCompact
	// FormatMinimal formats without thousand separators
	FormatMinimal
	// FormatSymbol formats with symbol only (no currency code)
	FormatSymbol
	// FormatCode formats with currency code only (same as String() method)
	FormatCode
)

// FormatOptions holds options for formatting money
type FormatOptions struct {
	Locale        language.Tag
	Mode          FormatMode
	CustomPattern string // Reserved for future custom format string support
}

// formatNumber formats a number using locale-aware formatting
func formatNumber(locale language.Tag, amount float64, minorUnit int) string {
	// Create a decimal formatter with the appropriate scale
	formatter := number.Decimal(amount, number.Scale(minorUnit))

	// Format using the locale
	p := message.NewPrinter(locale)
	return p.Sprintf("%v", formatter)
}

// getCurrencySymbol returns the symbol for a currency.
// Falls back to currency code if symbol is not set.
func getCurrencySymbol(currency *Currency, currencyCode string) string {
	if currency != nil && currency.Symbol != "" {
		return currency.Symbol
	}
	// Fallback to currency code if symbol not found
	return currencyCode
}

// getSymbolPosition returns whether the symbol should be placed before (true) or after (false) the amount.
// Uses locale-aware positioning when available, otherwise falls back to currency default.
func getSymbolPosition(currency *Currency) bool {
	if currency == nil {
		return true // Default to before for unknown currencies
	}

	// If Symbol is empty, the currency doesn't have symbol data, default to before
	if currency.Symbol == "" {
		return true
	}

	return currency.SymbolPosition
}

// formatWithSymbol combines formatted number with currency symbol
func formatWithSymbol(formattedNumber, symbol string, position bool, isNegative bool) string {
	if position {
		return symbol + formattedNumber
	}
	return formattedNumber + " " + symbol
}

// formatAccounting formats negative amounts with parentheses
func formatAccounting(formattedNumber string, isNegative bool) string {
	if isNegative {
		// Remove the negative sign and wrap in parentheses
		cleaned := strings.TrimPrefix(formattedNumber, "-")
		return "(" + cleaned + ")"
	}
	return formattedNumber
}

// formatCompact formats large amounts in compact notation (K, M, B)
func formatCompact(amount float64, symbol string, position bool, minorUnit int) string {
	var compactAmount float64
	var suffix string

	absAmount := math.Abs(amount)

	if absAmount >= 1_000_000_000 {
		compactAmount = amount / 1_000_000_000
		suffix = "B"
	} else if absAmount >= 1_000_000 {
		compactAmount = amount / 1_000_000
		suffix = "M"
	} else if absAmount >= 1_000 {
		compactAmount = amount / 1_000
		suffix = "K"
	} else {
		// For amounts less than 1000, use standard formatting
		return formatStandard(amount, symbol, position, minorUnit, false)
	}

	// Format compact amount with 1 decimal place
	compactStr := fmt.Sprintf("%.1f%s", compactAmount, suffix)

	if position {
		return symbol + compactStr
	}
	return compactStr + " " + symbol
}

// formatStandard formats with standard symbol positioning
func formatStandard(amount float64, symbol string, position bool, minorUnit int, isNegative bool) string {
	// Use default locale (English) for standard formatting
	locale := language.English
	formattedNumber := formatNumber(locale, amount, minorUnit)
	return formatWithSymbol(formattedNumber, symbol, position, isNegative)
}

// formatMinimal formats without thousand separators
func formatMinimal(amount float64, symbol string, position bool, minorUnit int, isNegative bool) string {
	// Format without separators by using a simple format string
	formatStr := fmt.Sprintf("%%.%df", minorUnit)
	formattedNumber := fmt.Sprintf(formatStr, amount)
	return formatWithSymbol(formattedNumber, symbol, position, isNegative)
}
