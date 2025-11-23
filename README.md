# Good Money

A minimal Money and currency package.

## Features

- **Precise arithmetic** - Stores amounts as `int64` minor units (cents) to avoid floating-point errors
- **ISO 4217 support** - Full currency support with proper decimal handling per currency
- **Money allocation** - Split money by ratios or percentages without losing pennies
- **8 rounding schemes** - HalfUp, HalfDown, HalfEven (banker's), and more
- **Overflow protection** - Automatic detection and error reporting for arithmetic overflow/underflow
- **Zero-allocation performance** - Most operations allocate 0 bytes, optimized for speed

## Quick-start

### Creating Money

```go
import "github.com/the-nucleus-project/goodmoney"

// Create money instances
m1, _ := goodmoney.New(100.50, goodmoney.ETB)  // 100.50 ETB
m2, _ := goodmoney.New(50.25, goodmoney.ETB)   // 50.25 ETB
m3, _ := goodmoney.NewZero(goodmoney.ETB)      // 0.00 ETB
```

### Arithmetic Operations

```go
// Addition
sum, _ := goodmoney.Add(m1, m2)
fmt.Println(sum)  // 150.75 ETB

// Subtraction
diff, _ := m1.Subtract(m2)
fmt.Println(diff)  // 50.25 ETB

// Multiplication
product, _ := m1.Multiply(3)
fmt.Println(product)  // 301.50 ETB

// Division
quotient, _ := m1.Divide(2)
fmt.Println(quotient)  // 50.25 ETB
```

### Comparisons

```go
if m1.GreaterThan(m2) {
    fmt.Println("m1 is greater")
}

if m1.Equals(m2) {
    fmt.Println("amounts are equal")
}
```

### Splitting Money

```go
// Split by ratios
total, _ := goodmoney.New(100.00, goodmoney.ETB)
parts, _ := total.Allocate(3, 2, 1)  // 3:2:1 ratio
// Result: 50.00 ETB, 33.34 ETB, 16.66 ETB

// Split by percentages
payment, _ := goodmoney.New(1000.00, goodmoney.ETB)
shares, _ := payment.AllocateByPercentage(60.0, 25.0, 15.0)
// Result: 600.00 ETB, 250.00 ETB, 150.00 ETB
```

### Rounding

```go
amount, _ := goodmoney.New(100.55, goodmoney.ETB)
scheme := goodmoney.RoundHalfUp
rounded := amount.Round(&scheme)
fmt.Println(rounded)  // 101.00 ETB
```

### Utility Methods

```go
m, _ := goodmoney.New(100.50, goodmoney.ETB)

m.IsZero()       // false
m.IsPositive()   // true
m.IsNegative()   // false

m.Absolute()  // 100.50 ETB
m.Negative()  // -100.50 ETB
```

### Formatting

```go
m, _ := goodmoney.New(100.50, goodmoney.ETB)
fmt.Println(m)        // 100.50 ETB
fmt.Printf("%s", m)   // 100.50 ETB
```

### JSON Serialization

```go
// Marshal to JSON
m, _ := goodmoney.New(100.50, goodmoney.ETB)
jsonBytes, _ := m.MarshalJSON()
// {"amount":100.5,"currency":"ETB"}

// Unmarshal from JSON
var unmarshaled goodmoney.Money
unmarshaled.UnmarshalJSON(jsonBytes)
fmt.Println(&unmarshaled)  // 100.50 ETB
```

### Currency Validation

```go
goodmoney.ValidateCurrency(goodmoney.ETB)  // true
goodmoney.ValidateCurrency("INVALID")      // false
```

## Upcoming
    
- Post-1.0 (1.x)
    - **Enhanced formatting options**
        - **Currency conversion**: Based on exchange rates convert currencies
        - **Currency symbol formatting**: Display amounts with symbols (`$100.50`, `€100,50`, `£100.50`) with proper symbol positioning (before/after) based on currency rules
        - **Locale-aware number formatting**: Thousand separators (`,` or `.`) and decimal separators (`.` or `,`) according to locale conventions (e.g., `1,234.56 USD` vs `1.234,56 EUR`)
        - **Accounting format**: Display negative amounts in accounting notation `(100.50)` instead of `-100.50`
        - **Compact notation**: Abbreviated formats for large amounts (`$1.5K`, `$1.2M`, `$5.3B`, `€2.4K`)
        - **Custom format strings**: Fine-grained control via format patterns (e.g., `Format("$#,###.00")`, `Format("€#.##0,00")`)
        - **Format modes**: Multiple display modes (`Standard`, `Accounting`, `Compact`, `Minimal`, `Symbol`, `Code`)
        - **Internationalization**: Support for different locales (`en-US`, `de-DE`, `fr-FR`, `it-IT`, etc.) with locale-specific formatting rules
        - **Separate component access**: Access major and minor units independently (e.g., `100 dollars` and `50 cents`)
    - **Money parsing** - Parse from formatted strings ("$100.50", "100.50 USD", "€100,50")
    - **Percentage operations** - Calculate percentage of money (e.g., 15% of $100)
    - **Locale-aware formatting** - Format with currency symbols, locale-specific separators ($100.50 vs €100,50)
    - **Human-readable formatting** - "one hundred dollars and fifty cents"
    - **Money ranges/intervals** - Check if money falls within a range (between two amounts)
    - **Money scaling by float** - Multiply/divide by float64 (for ratios, percentages, exchange rates)
    - **Currency symbol formatting** - Format with symbol positioning (before/after based on currency)
    - **Accounting format** - Format negative amounts as (100.50) instead of -100.50
    - **Major/Minor unit access** - Get separate major and minor unit components
    - **Money aggregation** - Min(), Max(), Average() operations for slices of Money
    - **Money parsing validation** - Validate and parse money from various string formats
    - **Tolerance-based comparison** - Compare money within a tolerance range (for floating-point conversion)
    - **Banknote/coin breakdown** - Split money into currency denominations
    - **Format as different display modes** - Standard, accounting, compact formats    


## API

Currency
- `func GetCurrency(code string) *Currency`
- `func GetCurrencyByNumericCode(numericCode string) (Currency, string, error)`
- `func ValidateCurrency(code string) bool`

Money
- `func New(amount float64, code string) (*Money, error)`
- `func NewZero(code string) (*Money, error)`
- `func MustNew(amount float64, code string) *Money`
- `func (m Money) Absolute() *Money`
- `func Add(ms ...*Money) (*Money, error)`
- `func (m Money) Allocate(rs ...int) ([]*Money, error)`
- `func (m Money) AllocateByPercentage(ps ...float64) ([]*Money, error)`
- `func (m Money) Amount() float64`
- `func (m Money) Compare(om *Money) (int, error)`
- `func (m Money) Currency() string`
- `func (m Money) Equals(om *Money) (bool, error)`
- `func (m Money) GreaterThan(om *Money) (bool, error)`
- `func (m Money) GreaterThanOrEqual(om *Money) (bool, error)`
- `func (m Money) LessThan(om *Money) (bool, error)`
- `func (m Money) LessThanOrEqual(om *Money) (bool, error)`
- `func (m Money) IsNegative() bool`
- `func (m Money) IsPositive() bool`
- `func (m Money) IsValid() bool`
- `func (m Money) IsZero() bool`
- `func (m Money) Multiply(ms ...int64) (*Money, error)`
- `func (m Money) Divide(ds ...int64) (*Money, error)`
- `func (m Money) Negative() *Money`
- `func (m Money) Round(scheme *RoundScheme) *Money`
- `func (m Money) Subtract(ms ...*Money) (*Money, error)`

- `func (m Money) String() string`
- `func (m Money) MarshalJSON() ([]byte, error)`
- `func (m Money) UnmarshalJSON(b []byte) error`

- `func (m *Money) Scan(src interface{}) error`
- `func (m Money) Value() (driver.Value, error)`

### Benchmark

Benchmark results (go 1.x, 2s per benchmark, 4 CPUs):

| Function | Time (ns/op) | Memory (B/op) | Allocations (allocs/op) |
|----------|--------------|---------------|------------------------|
| **Simple Operations** | | | |
| `IsNegative()` | 0.48 | 0 | 0 |
| `IsPositive()` | 0.50 | 0 | 0 |
| `IsZero()` | 0.57 | 0 | 0 |
| `Negative()` | 0.47 | 0 | 0 |
| `Absolute()` | 0.54 | 0 | 0 |
| `Amount()` | 4.46 | 0 | 0 |
| `Multiply()` | 2.19 | 0 | 0 |
| `Divide()` | 3.61 | 0 | 0 |
| `Compare()` | 4.39 | 0 | 0 |
| `Equals()` | 4.06 | 0 | 0 |
| `LessThan()` | 5.48 | 0 | 0 |
| `GreaterThan()` | 6.06 | 0 | 0 |
| `Subtract()` | 12.89 | 0 | 0 |
| **Creation & Conversion** | | | |
| `New()` | 171.5 | 40 | 2 |
| `NewZero()` | 269.3 | 40 | 2 |
| `Round()` | 51.52 | 16 | 1 |
| `Currency()` | 2,163 | 0 | 0 |
| `String()` | 2,807 | 64 | 6 |
| **Arithmetic** | | | |
| `Add()` (3 values) | 80.52 | 16 | 1 |
| `Allocate()` (4 ratios) | 309.0 | 96 | 5 |
| `AllocateByPercentage()` (4 percentages) | 280.0 | 96 | 5 |
| **Serialization** | | | |
| `MarshalJSON()` | 3,012 | 72 | 2 |
| `UnmarshalJSON()` | 3,776 | 888 | 22 |
| `json.Marshal()` | 3,407 | 120 | 3 |
| `json.Unmarshal()` | 5,519 | 1,040 | 24 |
| **Database** | | | |
| `Value()` | 3,513 | 72 | 2 |
| `Scan()` | 4,866 | 912 | 23 |
| **Round Trips** | | | |
| `Marshal → Unmarshal` | 7,756 | 960 | 24 |
| `Value → Scan` | 8,211 | 984 | 25 |

**Notes:**
- Fastest operations: Simple boolean checks and arithmetic (0.5-6 ns/op, zero allocations)
- Memory-efficient: Most operations allocate 0-16 bytes
- JSON operations: ~3-5 μs per operation (acceptable for API use)
- Database operations: ~3.5-4.9 μs per operation (JSON format)
- `Currency()` is slower due to map lookup, but still < 3 μs

### Comparison with Other Go Money Packages

> **Disclaimer:** This comparison was generated using AI tools and is based on publicly available documentation and code analysis as of the latest review. Feature support may vary by version. Performance metrics for other libraries are estimates and should be verified with actual benchmarks. Users are encouraged to verify claims independently.

| Feature | **goodmoney** | govalues | rhymond | bojanz/currency |
|---------|----------------|----------|---------|-----------------|
| **Storage** | `int64` minor units | Floating point (`float64`) | Fixed point (decimal) | Floating point (`decimal.Decimal`) |
| **Precision** | Currency-dependent (0-4 decimals) | 19 digits | 18 digits | 39 digits |
| **Performance** | | | | |
| - Addition | ~80 ns/op (measured) | Unknown | Unknown | Unknown |
| - Comparison | ~4 ns/op (measured) | Unknown | Unknown | Unknown |
| - Multiply | ~2 ns/op (measured) | Unknown | Unknown | Unknown |
| - Divide | ~4 ns/op (measured) | Unknown | Unknown | Unknown |
| **Rounding** | 8 schemes (HalfUp, HalfDown, HalfEven, etc.) | Half to even | Not supported | Half up |
| **Allocation** | ✅ `Allocate()` & `AllocateByPercentage()` | ❓ Unknown | ❓ Unknown | ❓ Unknown |
| **ISO 4217** | ✅ Full support | ❌ Not listed | ❌ Not listed | ✅ Full support |
| **Division** | ✅ | ✅ | ❌ | ✅ |
| **Currency Conversion** | ❌ | ✅ | ❌ | ✅ |
| **Overflow Control** | ✅ (int64 bounds) | ✅ | ❌ | ✅ |
| **JSON Support** | ✅ Native | ✅ | ✅ | ✅ |
| **Database Support** | ✅ `Scan`/`Value` | ✅ | ✅ | ✅ |
| **Immutability** | ✅ | ✅ | ✅ | ✅ |

**Key Advantages of goodmoney:**
- ✅ **Fast arithmetic** - Simple operations are 2-6 ns/op with zero allocations
- ✅ **Memory efficient** - Most operations zero-allocation (0 B/op)
- ✅ **Currency-aware** - Full ISO 4217 support with proper decimal handling per currency
- ✅ **Allocation methods** - `Allocate()` and `AllocateByPercentage()` for splitting money without losing pennies
- ✅ **Multiple rounding schemes** - 8 different rounding modes (HalfUp, HalfDown, HalfEven, etc.)
- ✅ **Division support** - Fast division operation (~4 ns/op) with overflow protection
- ✅ **Type safety** - Currency mismatch detection with clear error messages
- ✅ **Simple API** - Clean, focused interface following Go conventions
- ✅ **Overflow protection** - Automatic detection and error reporting for arithmetic overflow/underflow

### Credits & Inspiration

This project was inspired by [Rhymond/go-money](https://github.com/Rhymond/go-money) and implements principles from Martin Fowler's [Money pattern](https://martinfowler.com/eaaCatalog/money.html).
