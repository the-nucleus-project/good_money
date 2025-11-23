package goodmoney

import "errors"

var (
	// ErrCurrencyDoesNotExist happens when the provided currency code does not exist
	// in the ISO 4217 Currency codes list
	ErrCurrencyCodeDoesNotExist = errors.New("currency doesn't exist")
)

type Currency struct {
	NumericCode string
	MinorUnit   int
}

// retrive currency by code
func getCurrency(code string) (Currency, error) {
	res, ok := CurrencyMap[code]
	if !ok {
		return Currency{}, ErrCurrencyCodeDoesNotExist
	}
	return res, nil
}

// if currency exists return true, otherwise false
func ValidateCurrency(code string) bool {
	_, ok := CurrencyMap[code]
	return ok
}

// GetCurrency retrieves a Currency by its ISO 4217 currency code.
// Returns nil if the currency code doesn't exist.
//
// Example:
//
//	currency := GetCurrency("USD")
//	if currency != nil {
//	    // currency.NumericCode == "840"
//	    // currency.MinorUnit == 2
//	}
func GetCurrency(code string) *Currency {
	c, err := getCurrency(code)
	if err != nil {
		return nil
	}
	return &c
}

// GetCurrencyByNumericCode retrieves a currency by its ISO 4217 numeric code.
// It returns the Currency, the currency code (e.g., "USD"), and an error if the numeric code doesn't exist.
//
// Example:
//
//	currency, code, err := GetCurrencyByNumericCode("840")
//	// Returns USD currency, "USD", nil
func GetCurrencyByNumericCode(numericCode string) (Currency, string, error) {
	for code, currency := range CurrencyMap {
		if currency.NumericCode == numericCode {
			return currency, code, nil
		}
	}
	return Currency{}, "", ErrCurrencyCodeDoesNotExist
}
