package goodmoney

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"math"
)

var (
	// ErrCurrencyMismatch happens when two compared Money don't have the same currency.
	ErrCurrencyMismatch = errors.New("currencies don't match")

	// ErrTooManyDecimalPlaces happens when a float has more decimal places than the currency supports.
	ErrTooManyDecimalPlaces = errors.New("too many decimal places for currency")

	// ErrNeedAtLeastOneMoney happens when Add or Subtract is called without any Money arguments.
	ErrNeedAtLeastOneMoney = errors.New("need at least one money")

	// ErrOverflow happens when an arithmetic operation would exceed int64 maximum value.
	ErrOverflow = errors.New("amount overflow")

	// ErrUnderflow happens when an arithmetic operation would exceed int64 minimum value.
	ErrUnderflow = errors.New("amount underflow")
)

// RoundScheme defines different rounding schemes
type RoundScheme int

const (
	// RoundHalfUp rounds 0.5 up (standard rounding)
	RoundHalfUp RoundScheme = iota
	// RoundHalfDown rounds 0.5 down
	RoundHalfDown
	// RoundTowardZero truncates toward zero
	RoundTowardZero
	// RoundAwayFromZero rounds away from zero
	RoundAwayFromZero
	// RoundHalfEven (Banker's rounding) rounds 0.5 to nearest even number
	RoundHalfEven
	// RoundCeiling always rounds up
	RoundCeiling
	// RoundFloor always rounds down
	RoundFloor
)

type Money struct {
	amount   int64
	currency *Currency
}

// New creates a new Money instance from a float64 amount and currency code.
// Returns an error if the currency code is invalid or the amount has too many decimal places.
//
// Example:
//
//	m, err := New(100.50, "USD")
//	if err != nil {
//	    // handle error
//	}
func New(amount float64, currencyCode string) (*Money, error) {
	c, err := getCurrency(currencyCode)
	if err != nil {
		return nil, err
	}

	// validate minor unit length of float
	multiplier := math.Pow10(c.MinorUnit)
	amountNormalized := amount * multiplier
	// A. Truncate the normalized amount to get the integer part.
	integerPart := math.Trunc(amountNormalized)

	// B. Calculate the absolute difference (the fractional part).
	// e.g., 123.45 - 123.0 = 0.45
	fractionalPart := math.Abs(amountNormalized - integerPart)

	// C. Compare the fractional part against a small tolerance (epsilon).
	// This is crucial for robust floating-point comparisons.
	// We use a small epsilon (1e-9) to account for binary float imprecision.
	epsilon := 1e-9

	// If the fractionalPart is greater than the epsilon,
	// it means there's a significant non-zero part remaining,
	// and the original 'amount' had too many decimal places.

	if fractionalPart > epsilon {
		return nil, ErrTooManyDecimalPlaces
	}

	amountInMinorUnits := int64(amount * multiplier)

	return &Money{
		amount:   amountInMinorUnits,
		currency: &c,
	}, nil
}

// NewZero creates a new Money instance with zero amount for the given currency code.
// Returns an error if the currency code is invalid.
//
// Example:
//
//	m, err := NewZero("USD") // Creates $0.00
func NewZero(code string) (*Money, error) {
	return New(0, code)
}

// MustNew creates a new Money instance from a float64 amount and currency code.
// Panics if the currency code is invalid or the amount has too many decimal places.
// Use this function only when you are certain the inputs are valid.
//
// Example:
//
//	m := MustNew(100.50, "USD") // Panics if invalid
func MustNew(amount float64, currencyCode string) *Money {
	m, err := New(amount, currencyCode)
	if err != nil {
		panic(err)
	}
	return m
}

// get absolute value of amount
func (m Money) Absolute() *Money {
	return &Money{
		amount:   absInt64(m.amount),
		currency: m.currency,
	}
}

// Amount returns the amount in float64 (internal storage scheme undisclosed).
// Returns 0.0 if currency is nil.
func (m Money) Amount() float64 {
	if m.currency == nil {
		return 0.0
	}
	// Based on currency get minor units
	multiplier := math.Pow10(m.currency.MinorUnit)
	return float64(m.amount) / multiplier
}

// Compare compares two Money values.
// Returns:
//
//	-1 if om is less than m
//	 0 if om and m are equal
//	 1 if om is greater than m
//
// Returns an error if currencies don't match.
func (m Money) Compare(om *Money) (int, error) {
	//validate currency mismatch
	if m.currency == nil || om.currency == nil || m.currency.NumericCode != om.currency.NumericCode {
		return 0, ErrCurrencyMismatch
	}

	if om.amount < m.amount {
		return -1, nil
	}
	if om.amount == m.amount {
		return 0, nil
	}
	return 1, nil
}

// Currency returns the currency code string (e.g., "USD", "EUR").
// Returns empty string if currency is nil.
func (m Money) Currency() string {
	if m.currency == nil {
		return ""
	}
	_, code, _ := GetCurrencyByNumericCode(m.currency.NumericCode)
	return code
}

// check if equal
func (m Money) Equals(om *Money) (bool, error) {
	//validate currency mismatch
	if m.currency == nil || om.currency == nil || m.currency.NumericCode != om.currency.NumericCode {
		return false, ErrCurrencyMismatch
	}
	if om.amount == m.amount {
		return true, nil
	}
	return false, nil
}

// equals to wrapper for GreaterThan
func (m Money) GreaterThan(om *Money) (bool, error) {
	res, err := m.Equals(om)
	if err != nil {
		return false, err
	}
	if res {
		return false, nil
	}
	return m.amount > om.amount, nil
}

// equals to wrapper for GreaterThanOrEqual
func (m Money) GreaterThanOrEqual(om *Money) (bool, error) {
	res, err := m.Equals(om)
	if err != nil {
		return false, err
	}
	if res {
		return true, nil
	}
	return m.amount >= om.amount, nil
}

// equals to wrapper for LessThan
func (m Money) LessThan(om *Money) (bool, error) {
	res, err := m.Equals(om)
	if err != nil {
		return false, err
	}
	if res {
		return false, nil
	}
	return m.amount < om.amount, nil
}

// equals to wrapper for GreaterThanOrEqual
func (m Money) LessThanOrEqual(om *Money) (bool, error) {
	res, err := m.Equals(om)
	if err != nil {
		return false, err
	}
	if res {
		return true, nil
	}
	return m.amount <= om.amount, nil
}

// check if negative
func (m Money) IsNegative() bool {
	return m.amount < 0
}

// check if positive
func (m Money) IsPositive() bool {
	return m.amount > 0
}

// IsValid returns true if the Money instance has a valid (non-nil) currency.
func (m Money) IsValid() bool {
	return m.currency != nil
}

// IsZero returns true if the money amount is zero.
func (m Money) IsZero() bool {
	return m.amount == 0
}

// Multiply multiplies the money amount by one or more int64 factors.
// Returns an error if the result would overflow int64.
func (m Money) Multiply(ms ...int64) (*Money, error) {
	result := m.amount
	for _, multiplier := range ms {
		// Check for overflow before multiplication
		if result != 0 && multiplier != 0 {
			// Check positive overflow (result * multiplier > MaxInt64)
			if result > 0 && multiplier > 0 {
				if result > math.MaxInt64/multiplier {
					return nil, ErrOverflow
				}
			}
			// Check negative overflow (result * multiplier < MinInt64)
			if result > 0 && multiplier < 0 {
				if multiplier < math.MinInt64/result {
					return nil, ErrUnderflow
				}
			}
			if result < 0 && multiplier > 0 {
				if result < math.MinInt64/multiplier {
					return nil, ErrUnderflow
				}
			}
			// For negative * negative, check if result would overflow MaxInt64
			if result < 0 && multiplier < 0 {
				if result < math.MaxInt64/multiplier {
					return nil, ErrOverflow
				}
			}
		}
		result *= multiplier
	}
	return &Money{
		amount:   result,
		currency: m.currency,
	}, nil
}

// Divide divides the money amount by one or more int64 divisors.
// Returns an error if any divisor is zero.
//
// Example:
//
//	half, err := m.Divide(2)        // Divide by 2
//	quarter, err := m.Divide(2, 2)  // Divide by 2 twice
func (m Money) Divide(ds ...int64) (*Money, error) {
	result := m.amount
	for _, divisor := range ds {
		if divisor == 0 {
			return nil, errors.New("division by zero")
		}
		result /= divisor
	}
	return &Money{
		amount:   result,
		currency: m.currency,
	}, nil
}

// get negative of amount
func (m Money) Negative() *Money {
	return &Money{
		amount:   m.amount * -1,
		currency: m.currency,
	}
}

// round with the specified rounding scheme
// Defaults to RoundTowardZero if scheme is nil
func (m Money) Round(scheme *RoundScheme) *Money {
	if m.currency == nil {
		return &Money{
			amount:   m.amount,
			currency: nil,
		}
	}

	// Default to RoundTowardZero if no scheme provided
	roundScheme := RoundTowardZero
	if scheme != nil {
		roundScheme = *scheme
	}

	// Convert to float64, apply rounding scheme, then convert back
	amountFloat := m.Amount()
	roundedAmount := applyRoundScheme(amountFloat, roundScheme)

	// Convert back to minor units
	multiplier := math.Pow10(m.currency.MinorUnit)
	amountInMinorUnits := int64(roundedAmount * multiplier)

	return &Money{
		amount:   amountInMinorUnits,
		currency: m.currency,
	}
}

// applyRoundScheme applies the specified rounding scheme to a float64 value
func applyRoundScheme(amount float64, scheme RoundScheme) float64 {
	switch scheme {
	case RoundHalfUp:
		if amount >= 0 {
			return math.Floor(amount + 0.5)
		}
		return math.Ceil(amount + 0.5)
	case RoundHalfDown:
		if amount >= 0 {
			return math.Ceil(amount - 0.5)
		}
		return math.Floor(amount - 0.5)
	case RoundTowardZero:
		return math.Trunc(amount)
	case RoundAwayFromZero:
		if amount >= 0 {
			return math.Ceil(amount)
		}
		return math.Floor(amount)
	case RoundHalfEven:
		// Banker's rounding: round to nearest even
		floor := math.Floor(amount)
		frac := amount - floor
		if frac < 0.5 {
			return floor
		}
		if frac > 0.5 {
			return floor + 1
		}
		// Exactly 0.5 - round to nearest even
		if int64(floor)%2 == 0 {
			return floor
		}
		return floor + 1
	case RoundCeiling:
		return math.Ceil(amount)
	case RoundFloor:
		return math.Floor(amount)
	default:
		return math.Trunc(amount)
	}
}

// Add adds two or more Money values together.
// Returns an error if currencies don't match, if any Money is nil, or if overflow occurs.
//
// Example:
//
//	total, err := Add(m1, m2, m3)
func Add(ms ...*Money) (*Money, error) {
	// need at least one money argument
	if len(ms) == 0 {
		return nil, ErrNeedAtLeastOneMoney
	}

	// validate first money and get reference currency
	firstMoney := ms[0]
	if firstMoney == nil || firstMoney.currency == nil {
		return nil, ErrCurrencyMismatch
	}

	referenceCurrency := firstMoney.currency.NumericCode
	result := firstMoney.amount

	// optimize: early return for single Money
	if len(ms) == 1 {
		return &Money{
			amount:   result,
			currency: firstMoney.currency,
		}, nil
	}

	// validate and sum remaining money values
	for i := 1; i < len(ms); i++ {
		money := ms[i]
		if money == nil || money.currency == nil || money.currency.NumericCode != referenceCurrency {
			return nil, ErrCurrencyMismatch
		}
		// Check for overflow before addition
		if money.amount > 0 && result > math.MaxInt64-money.amount {
			return nil, ErrOverflow
		}
		if money.amount < 0 && result < math.MinInt64-money.amount {
			return nil, ErrUnderflow
		}
		result += money.amount
	}

	return &Money{
		amount:   result,
		currency: firstMoney.currency,
	}, nil
}

// Subtract subtracts one or more Money values from this Money.
// Returns an error if currencies don't match, if any Money is nil, or if underflow occurs.
//
// Example:
//
//	diff, err := m1.Subtract(m2, m3)
func (m Money) Subtract(ms ...*Money) (*Money, error) {
	// need at least one money argument
	if len(ms) == 0 {
		return nil, ErrNeedAtLeastOneMoney
	}

	referenceCurrency := m.currency.NumericCode
	result := m.amount

	// validate and subtract all money values
	for i := 0; i < len(ms); i++ {
		money := ms[i]
		if money == nil || money.currency == nil || money.currency.NumericCode != referenceCurrency {
			return nil, ErrCurrencyMismatch
		}
		// Check for overflow before subtraction
		if money.amount < 0 && result > math.MaxInt64+money.amount {
			return nil, ErrOverflow
		}
		if money.amount > 0 && result < math.MinInt64+money.amount {
			return nil, ErrUnderflow
		}
		result -= money.amount
	}

	return &Money{
		amount:   result,
		currency: m.currency,
	}, nil
}

// Allocate splits money by given ratios without losing pennies.
// Leftover pennies are distributed amongst the parties using round-robin principle.
//
// Example:
//
//	parts, err := total.Allocate(3, 2, 1) // Split in 3:2:1 ratio
func (m *Money) Allocate(rs ...int) ([]*Money, error) {
	if len(rs) == 0 {
		return nil, errors.New("no ratios specified")
	}

	// Calculate sum of ratios.
	var sum int64
	for _, r := range rs {
		if r < 0 {
			return nil, errors.New("negative ratios not allowed")
		}
		if int64(r) > (math.MaxInt64 - sum) {
			return nil, errors.New("sum of given ratios exceeds max int")
		}
		sum += int64(r)
	}

	var total int64
	ms := make([]*Money, 0, len(rs))
	for _, r := range rs {
		// Calculate allocated amount: (amount * ratio) / sum
		var allocatedAmount int64
		if sum != 0 {
			allocatedAmount = (m.amount * int64(r)) / sum
		}

		party := &Money{
			amount:   allocatedAmount,
			currency: m.currency,
		}

		ms = append(ms, party)
		total += party.amount
	}

	// if the sum of all ratios is zero, then we just returns zeros and don't do anything
	// with the leftover
	if sum == 0 {
		return ms, nil
	}

	// Calculate leftover value and divide to first parties.
	lo := m.amount - total
	sub := int64(1)
	if lo < 0 {
		sub = -sub
	}

	for p := 0; lo != 0; p = (p + 1) % len(ms) {
		ms[p].amount = ms[p].amount + sub
		lo -= sub
	}

	return ms, nil
}

// AllocateByPercentage splits money by given percentages without losing pennies.
// Leftover pennies are distributed amongst the parties using round-robin principle.
//
// Example:
//
//	shares, err := payment.AllocateByPercentage(60.0, 25.0, 15.0)
func (m *Money) AllocateByPercentage(ps ...float64) ([]*Money, error) {
	if len(ps) == 0 {
		return nil, errors.New("no percentages specified")
	}

	// Calculate sum of percentages.
	var sum float64
	for _, p := range ps {
		if p < 0 {
			return nil, errors.New("negative percentages not allowed")
		}
		if p > (math.MaxFloat64 - sum) {
			return nil, errors.New("sum of given percentages exceeds max float")
		}
		sum += p
	}

	var total int64
	ms := make([]*Money, 0, len(ps))
	for _, p := range ps {
		// Calculate allocated amount: (amount * percentage) / sum
		var allocatedAmount int64
		if sum != 0 {
			// Calculate proportion and multiply by amount
			proportion := p / sum
			allocatedAmount = int64(float64(m.amount) * proportion)
		}

		party := &Money{
			amount:   allocatedAmount,
			currency: m.currency,
		}

		ms = append(ms, party)
		total += party.amount
	}

	// if the sum of all percentages is zero, then we just returns zeros and don't do anything
	// with the leftover
	if sum == 0 {
		return ms, nil
	}

	// Calculate leftover value and divide to first parties.
	lo := m.amount - total
	sub := int64(1)
	if lo < 0 {
		sub = -sub
	}

	for p := 0; lo != 0; p = (p + 1) % len(ms) {
		ms[p].amount = ms[p].amount + sub
		lo -= sub
	}

	return ms, nil
}

// String returns a string representation of Money in the format "amount currency".
// Example: "100.50 USD", "-50.25 ETB", "0.00 EUR"
// Implements fmt.Stringer interface.
func (m Money) String() string {
	if m.currency == nil {
		// If currency is nil, just show the amount in minor units
		return fmt.Sprintf("%d (no currency)", m.amount)
	}

	currencyCode := m.Currency()
	amount := m.Amount()

	// Format amount based on currency's minor unit precision
	formatStr := fmt.Sprintf("%%.%df", m.currency.MinorUnit)
	formattedAmount := fmt.Sprintf(formatStr, amount)

	return fmt.Sprintf("%s %s", formattedAmount, currencyCode)
}

// moneyJSON represents the JSON structure for Money serialization
type moneyJSON struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// MarshalJSON implements json.Marshaler interface.
// It serializes Money to JSON in the format: {"amount": 100.50, "currency": "USD"}
func (m Money) MarshalJSON() ([]byte, error) {
	if m.currency == nil {
		return nil, errors.New("cannot marshal Money with nil currency")
	}

	return json.Marshal(moneyJSON{
		Amount:   m.Amount(),
		Currency: m.Currency(),
	})
}

// UnmarshalJSON implements json.Unmarshaler interface.
// It deserializes JSON in the format: {"amount": 100.50, "currency": "USD"}
func (m *Money) UnmarshalJSON(data []byte) error {
	var j moneyJSON
	if err := json.Unmarshal(data, &j); err != nil {
		return fmt.Errorf("failed to unmarshal Money: %w", err)
	}

	// Validate that both fields are present by checking if amount was explicitly set
	// This is a workaround since Go's json.Unmarshal doesn't distinguish between
	// missing field and zero value. We check by re-parsing and looking for the field.
	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err == nil {
		if _, ok := raw["amount"]; !ok {
			return errors.New("amount field is required")
		}
		if _, ok := raw["currency"]; !ok {
			return errors.New("currency field is required")
		}
	}

	// Validate currency code exists
	if j.Currency == "" {
		return errors.New("currency code cannot be empty")
	}

	// Create Money using New() to ensure validation
	newMoney, err := New(j.Amount, j.Currency)
	if err != nil {
		return fmt.Errorf("failed to create Money from JSON: %w", err)
	}

	// Copy the values
	*m = *newMoney
	return nil
}

// Value implements driver.Valuer interface.
// It returns Money as JSON bytes for database storage.
// Returns nil if currency is nil.
func (m Money) Value() (driver.Value, error) {
	if m.currency == nil {
		return nil, nil
	}
	return m.MarshalJSON()
}

// Scan implements sql.Scanner interface.
// It reads Money from database value (JSON bytes or string).
func (m *Money) Scan(src interface{}) error {
	if src == nil {
		*m = Money{}
		return nil
	}

	var data []byte
	switch v := src.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("cannot scan %T into Money", src)
	}

	if len(data) == 0 {
		*m = Money{}
		return nil
	}

	return m.UnmarshalJSON(data)
}
