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

// currencySymbols maps ISO 4217 currency codes to their symbols
var currencySymbols = map[string]string{
	USD: "$",
	EUR: "€",
	GBP: "£",
	JPY: "¥",
	CNY: "¥",
	INR: "₹",
	KRW: "₩",
	BRL: "R$",
	RUB: "₽",
	ZAR: "R",
	AUD: "A$",
	CAD: "C$",
	CHF: "CHF",
	DKK: "kr",
	NOK: "kr",
	SEK: "kr",
	PLN: "zł",
	TRY: "₺",
	MXN: "$",
	ARS: "$",
	CLP: "$",
	COP: "$",
	PEN: "S/",
	PHP: "₱",
	THB: "฿",
	VND: "₫",
	IDR: "Rp",
	MYR: "RM",
	SGD: "S$",
	HKD: "HK$",
	NZD: "NZ$",
	ETB: "Br",
	EGP: "E£",
	ILS: "₪",
	SAR: "﷼",
	AED: "د.إ",
	QAR: "﷼",
	KWD: "د.ك",
	BHD: ".د.ب",
	OMR: "﷼",
	JOD: "د.ا",
	LBP: "£",
	IQD: "ع.د",
	IRR: "﷼",
	PKR: "₨",
	BDT: "৳",
	LKR: "₨",
	NPR: "₨",
	MMK: "K",
	KHR: "៛",
	LAK: "₭",
	MNT: "₮",
	KZT: "₸",
	UZS: "so'm",
	AZN: "₼",
	GEL: "₾",
	AMD: "֏",
	BYN: "Br",
	UAH: "₴",
	MDL: "L",
	RON: "lei",
	BGN: "лв",
	RSD: "дин",
	MKD: "ден",
	ALL: "L",
	BAM: "КМ",
	HUF: "Ft",
	CZK: "Kč",
	ISK: "kr",
	// Add more as needed
}

// symbolPosition maps currency codes to their default symbol position
// true = before amount, false = after amount
var symbolPosition = map[string]bool{
	USD: true,
	EUR: false, // €100,50 in many European locales
	GBP: true,
	JPY: true,
	CNY: true,
	INR: true,
	KRW: true,
	BRL: true,
	RUB: false,
	ZAR: true,
	AUD: true,
	CAD: true,
	CHF: true,
	DKK: false,
	NOK: false,
	SEK: false,
	PLN: false,
	TRY: true,
	MXN: true,
	ARS: true,
	CLP: true,
	COP: true,
	PEN: true,
	PHP: true,
	THB: true,
	VND: false,
	IDR: true,
	MYR: true,
	SGD: true,
	HKD: true,
	NZD: true,
	ETB: true,
	EGP: true,
	ILS: true,
	SAR: true,
	AED: true,
	QAR: true,
	KWD: true,
	BHD: true,
	OMR: true,
	JOD: true,
	LBP: true,
	IQD: false,
	IRR: true,
	PKR: true,
	BDT: true,
	LKR: true,
	NPR: true,
	MMK: true,
	KHR: false,
	LAK: false,
	MNT: true,
	KZT: false,
	UZS: false,
	AZN: true,
	GEL: true,
	AMD: false,
	BYN: false,
	UAH: true,
	MDL: false,
	RON: false,
	BGN: false,
	RSD: false,
	MKD: false,
	ALL: false,
	BAM: false,
	HUF: false,
	CZK: false,
	// Default to true (before) for currencies not in this map
}

// formatNumber formats a number using locale-aware formatting
func formatNumber(locale language.Tag, amount float64, minorUnit int) string {
	// Create a decimal formatter with the appropriate scale
	formatter := number.Decimal(amount, number.Scale(minorUnit))

	// Format using the locale
	p := message.NewPrinter(locale)
	return p.Sprintf("%v", formatter)
}

// getCurrencySymbol returns the symbol for a currency code
func getCurrencySymbol(currencyCode string) string {
	if symbol, ok := currencySymbols[currencyCode]; ok {
		return symbol
	}
	// Fallback to currency code if symbol not found
	return currencyCode
}

// getSymbolPosition returns whether the symbol should be placed before (true) or after (false) the amount
func getSymbolPosition(currencyCode string, locale language.Tag) bool {
	// Check if we have a specific position for this currency
	if position, ok := symbolPosition[currencyCode]; ok {
		// Some currencies have locale-specific positioning
		// For now, we use the default, but this can be enhanced with locale-specific rules
		return position
	}
	// Default to before for unknown currencies
	return true
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
