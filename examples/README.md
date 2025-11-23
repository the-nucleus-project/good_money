# Good Money Examples

This directory contains comprehensive examples for all public APIs in the goodmoney package.

## Examples

### Currency APIs
- **currency_examples.go** - Examples for `GetCurrency`, `GetCurrencyByNumericCode`, and `ValidateCurrency`

### Money Creation
- **money_creation_examples.go** - Examples for `New`, `NewZero`, and `MustNew`

### Money Arithmetic
- **money_arithmetic_examples.go** - Examples for `Add`, `Subtract`, `Multiply`, and `Divide`

### Money Comparison
- **money_comparison_examples.go** - Examples for `Compare`, `Equals`, `GreaterThan`, `GreaterThanOrEqual`, `LessThan`, and `LessThanOrEqual`

### Money Utilities
- **money_utility_examples.go** - Examples for:
  - Boolean checks: `IsNegative`, `IsPositive`, `IsZero`, `IsValid`
  - Transformations: `Absolute`, `Negative`
  - Accessors: `Amount`, `MajorUnit`, `MinorUnit`
  - Operations: `Round`, `Allocate`, `AllocateByPercentage`

### Money Formatting
- **money_formatting_examples.go** - Examples for:
  - `String()` - Basic string representation
  - `Format()` - Locale-aware formatting
  - `FormatWithMode()` - Formatting with different modes (Standard, Accounting, Compact, Minimal, Symbol, Code)
  - `FormatWithOptions()` - Flexible formatting with options

### Money Serialization
- **money_serialization_examples.go** - Examples for `MarshalJSON` and `UnmarshalJSON`

### Money Database Operations
- **money_database_examples.go** - Examples for `Value()` and `Scan()` for database integration

## Running Examples

To run any example:

```bash
go run examples/currency_examples.go
go run examples/money_creation_examples.go
go run examples/money_arithmetic_examples.go
# ... etc
```

Or run all examples:

```bash
for file in examples/*_examples.go; do
    echo "Running $file"
    go run "$file"
    echo ""
done
```

