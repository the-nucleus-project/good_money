package goodmoney

import (
	"testing"

	"golang.org/x/text/language"
)

func TestFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		inputAmount       float64
		inputCurrencyCode string
		locale            language.Tag
		wantContains      []string // Check that result contains these strings
		wantNotContains   []string // Check that result does not contain these strings
	}{
		{
			name:              "USD with American English locale",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			wantContains:      []string{"$", "1,234.56"},
		},
		{
			name:              "EUR with German locale",
			inputAmount:       1234.56,
			inputCurrencyCode: EUR,
			locale:            language.German,
			wantContains:      []string{"1.234", "56"},
		},
		{
			name:              "USD with French locale",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.French,
			wantContains:      []string{"1", "234", "56"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			money, err := New(tt.inputAmount, tt.inputCurrencyCode)
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			got := money.Format(tt.locale)

			for _, want := range tt.wantContains {
				if !contains(got, want) {
					t.Errorf("Format() = %q, should contain %q", got, want)
				}
			}

			for _, notWant := range tt.wantNotContains {
				if contains(got, notWant) {
					t.Errorf("Format() = %q, should not contain %q", got, notWant)
				}
			}
		})
	}
}

func TestFormatWithMode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		inputAmount       float64
		inputCurrencyCode string
		locale            language.Tag
		mode              FormatMode
		wantContains      []string
		wantNotContains   []string
	}{
		{
			name:              "FormatStandard USD",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatStandard,
			wantContains:      []string{"$", "1,234.56"},
		},
		{
			name:              "FormatCode USD",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatCode,
			wantContains:      []string{"1234.56", "USD"},
		},
		{
			name:              "FormatSymbol USD",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatSymbol,
			wantContains:      []string{"$", "1,234.56"},
			wantNotContains:   []string{"USD"},
		},
		{
			name:              "FormatAccounting positive USD",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatAccounting,
			wantContains:      []string{"$", "1,234.56"},
			wantNotContains:   []string{"("},
		},
		{
			name:              "FormatAccounting negative USD",
			inputAmount:       -1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatAccounting,
			wantContains:      []string{"$", "(", "1,234.56", ")"},
		},
		{
			name:              "FormatCompact millions USD",
			inputAmount:       1500000.00,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatCompact,
			wantContains:      []string{"$", "1.5M"},
		},
		{
			name:              "FormatCompact thousands USD",
			inputAmount:       1500.00,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatCompact,
			wantContains:      []string{"$", "1.5K"},
		},
		{
			name:              "FormatCompact billions USD",
			inputAmount:       5300000000.00,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatCompact,
			wantContains:      []string{"$", "5.3B"},
		},
		{
			name:              "FormatMinimal USD",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatMinimal,
			wantContains:      []string{"$", "1234.56"},
			wantNotContains:   []string{","},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			money, err := New(tt.inputAmount, tt.inputCurrencyCode)
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			got := money.FormatWithMode(tt.locale, tt.mode)

			for _, want := range tt.wantContains {
				if !contains(got, want) {
					t.Errorf("FormatWithMode() = %q, should contain %q", got, want)
				}
			}

			for _, notWant := range tt.wantNotContains {
				if contains(got, notWant) {
					t.Errorf("FormatWithMode() = %q, should not contain %q", got, notWant)
				}
			}
		})
	}
}

func TestFormatWithOptions(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		inputAmount       float64
		inputCurrencyCode string
		opts              FormatOptions
		wantContains      []string
	}{
		{
			name:              "FormatOptions with German locale",
			inputAmount:       1234.56,
			inputCurrencyCode: EUR,
			opts: FormatOptions{
				Locale: language.German,
				Mode:   FormatStandard,
			},
			wantContains: []string{"1.234", "56"},
		},
		{
			name:              "FormatOptions with Accounting mode",
			inputAmount:       -100.50,
			inputCurrencyCode: USD,
			opts: FormatOptions{
				Locale: language.AmericanEnglish,
				Mode:   FormatAccounting,
			},
			wantContains: []string{"$", "(", "100.50", ")"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			money, err := New(tt.inputAmount, tt.inputCurrencyCode)
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			got := money.FormatWithOptions(tt.opts)

			for _, want := range tt.wantContains {
				if !contains(got, want) {
					t.Errorf("FormatWithOptions() = %q, should contain %q", got, want)
				}
			}
		})
	}
}

func TestFormatLocaleSpecific(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		inputAmount       float64
		inputCurrencyCode string
		locale            language.Tag
		description       string
	}{
		{
			name:              "American English formatting",
			inputAmount:       1234.56,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			description:       "Should use comma for thousands, period for decimals",
		},
		{
			name:              "German formatting",
			inputAmount:       1234.56,
			inputCurrencyCode: EUR,
			locale:            language.German,
			description:       "Should use period for thousands, comma for decimals",
		},
		{
			name:              "French formatting",
			inputAmount:       1234.56,
			inputCurrencyCode: EUR,
			locale:            language.French,
			description:       "Should use space for thousands, comma for decimals",
		},
		{
			name:              "Italian formatting",
			inputAmount:       1234.56,
			inputCurrencyCode: EUR,
			locale:            language.Italian,
			description:       "Should use period for thousands, comma for decimals",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			money, err := New(tt.inputAmount, tt.inputCurrencyCode)
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			got := money.Format(tt.locale)
			if got == "" {
				t.Errorf("Format() returned empty string for %s", tt.description)
			}
			// Just verify it formats without error and contains the amount
			// Check for various locale-specific formats: "1234", "1.234", "1 234", "1,234"
			hasAmount := contains(got, "1234") || contains(got, "1.234") ||
				contains(got, "1 234") || contains(got, "1,234") ||
				contains(got, "1\u00a0234") // non-breaking space
			if !hasAmount {
				t.Errorf("Format() = %q, should contain amount representation", got)
			}
		})
	}
}

func TestFormatCurrencySymbols(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		inputAmount       float64
		inputCurrencyCode string
		locale            language.Tag
		wantContains      string
	}{
		{
			name:              "USD symbol",
			inputAmount:       100.50,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			wantContains:      "$",
		},
		{
			name:              "EUR symbol",
			inputAmount:       100.50,
			inputCurrencyCode: EUR,
			locale:            language.German,
			wantContains:      "€",
		},
		{
			name:              "GBP symbol",
			inputAmount:       100.50,
			inputCurrencyCode: GBP,
			locale:            language.BritishEnglish,
			wantContains:      "£",
		},
		{
			name:              "JPY symbol",
			inputAmount:       1000.0,
			inputCurrencyCode: JPY,
			locale:            language.Japanese,
			wantContains:      "¥",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			money, err := New(tt.inputAmount, tt.inputCurrencyCode)
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			got := money.FormatWithMode(tt.locale, FormatSymbol)
			if !contains(got, tt.wantContains) {
				t.Errorf("FormatWithMode() = %q, should contain %q", got, tt.wantContains)
			}
		})
	}
}

func TestFormatNegativeAmounts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		inputAmount       float64
		inputCurrencyCode string
		locale            language.Tag
		mode              FormatMode
		wantContains      []string
		wantNotContains   []string
	}{
		{
			name:              "Negative Standard format",
			inputAmount:       -100.50,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatStandard,
			wantContains:      []string{"$", "-", "100.50"},
		},
		{
			name:              "Negative Accounting format",
			inputAmount:       -100.50,
			inputCurrencyCode: USD,
			locale:            language.AmericanEnglish,
			mode:              FormatAccounting,
			wantContains:      []string{"$", "(", "100.50", ")"},
			wantNotContains:   []string{"-"},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			money, err := New(tt.inputAmount, tt.inputCurrencyCode)
			if err != nil {
				t.Fatalf("New() unexpected error: %v", err)
			}

			got := money.FormatWithMode(tt.locale, tt.mode)

			for _, want := range tt.wantContains {
				if !contains(got, want) {
					t.Errorf("FormatWithMode() = %q, should contain %q", got, want)
				}
			}

			for _, notWant := range tt.wantNotContains {
				if contains(got, notWant) {
					t.Errorf("FormatWithMode() = %q, should not contain %q", got, notWant)
				}
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
