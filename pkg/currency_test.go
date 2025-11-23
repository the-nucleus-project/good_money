package goodmoney

import (
	"testing"
)

func TestGetCurrencyByNumericCode(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		numericCode  string
		wantCurrency Currency
		wantCode     string
		wantErr      error
	}{
		// Valid numeric codes
		{
			name:        "USD numeric code",
			numericCode: "840",
			wantCurrency: Currency{
				NumericCode: "840",
				MinorUnit:   2,
			},
			wantCode: USD,
			wantErr:  nil,
		},
		{
			name:        "EUR numeric code",
			numericCode: "978",
			wantCurrency: Currency{
				NumericCode: "978",
				MinorUnit:   2,
			},
			wantCode: EUR,
			wantErr:  nil,
		},
		{
			name:        "JPY numeric code",
			numericCode: "392",
			wantCurrency: Currency{
				NumericCode: "392",
				MinorUnit:   0,
			},
			wantCode: JPY,
			wantErr:  nil,
		},
		{
			name:        "BHD numeric code (3 decimal places)",
			numericCode: "048",
			wantCurrency: Currency{
				NumericCode: "048",
				MinorUnit:   3,
			},
			wantCode: BHD,
			wantErr:  nil,
		},
		{
			name:        "GBP numeric code",
			numericCode: "826",
			wantCurrency: Currency{
				NumericCode: "826",
				MinorUnit:   2,
			},
			wantCode: GBP,
			wantErr:  nil,
		},
		{
			name:        "ETB numeric code",
			numericCode: "230",
			wantCurrency: Currency{
				NumericCode: "230",
				MinorUnit:   2,
			},
			wantCode: ETB,
			wantErr:  nil,
		},
		{
			name:        "CLF numeric code (4 decimal places)",
			numericCode: "990",
			wantCurrency: Currency{
				NumericCode: "990",
				MinorUnit:   4,
			},
			wantCode: CLF,
			wantErr:  nil,
		},

		// Invalid numeric codes
		{
			name:         "invalid numeric code",
			numericCode:  "777",
			wantCurrency: Currency{},
			wantCode:     "",
			wantErr:      ErrCurrencyCodeDoesNotExist,
		},
		{
			name:         "empty numeric code",
			numericCode:  "",
			wantCurrency: Currency{},
			wantCode:     "",
			wantErr:      ErrCurrencyCodeDoesNotExist,
		},
		{
			name:         "non-existent numeric code",
			numericCode:  "000",
			wantCurrency: Currency{},
			wantCode:     "",
			wantErr:      ErrCurrencyCodeDoesNotExist,
		},
	}

	for _, tt := range tests {
		tt := tt // Capture loop variable for parallel subtests
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			gotCurrency, gotCode, err := GetCurrencyByNumericCode(tt.numericCode)

			if tt.wantErr != nil {
				if err == nil {
					t.Errorf("GetCurrencyByNumericCode() expected error %v, got nil", tt.wantErr)
					return
				}
				if err != tt.wantErr {
					t.Errorf("GetCurrencyByNumericCode() expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Errorf("GetCurrencyByNumericCode() unexpected error: %v", err)
				return
			}

			if gotCurrency.NumericCode != tt.wantCurrency.NumericCode {
				t.Errorf("GetCurrencyByNumericCode() NumericCode = %s, want %s", gotCurrency.NumericCode, tt.wantCurrency.NumericCode)
			}

			if gotCurrency.MinorUnit != tt.wantCurrency.MinorUnit {
				t.Errorf("GetCurrencyByNumericCode() MinorUnit = %d, want %d", gotCurrency.MinorUnit, tt.wantCurrency.MinorUnit)
			}

			if gotCode != tt.wantCode {
				t.Errorf("GetCurrencyByNumericCode() currency code = %s, want %s", gotCode, tt.wantCode)
			}
		})
	}
}

func TestGetCurrencyByNumericCodeCoverage(t *testing.T) {
	t.Parallel()

	// Verify some known mappings
	testCases := map[string]string{
		"840": USD,
		"978": EUR,
		"392": JPY,
		"048": BHD,
		"826": GBP,
		"230": ETB,
		"990": CLF,
	}

	for numericCode, expectedCode := range testCases {
		gotCurrency, gotCode, err := GetCurrencyByNumericCode(numericCode)
		if err != nil {
			t.Errorf("getCurrencyByNumericCode(%s) unexpected error: %v", numericCode, err)
			continue
		}
		if gotCode != expectedCode {
			t.Errorf("getCurrencyByNumericCode(%s) code = %s, want %s", numericCode, gotCode, expectedCode)
		}
		if gotCurrency.NumericCode != numericCode {
			t.Errorf("getCurrencyByNumericCode(%s) NumericCode = %s, want %s", numericCode, gotCurrency.NumericCode, numericCode)
		}
	}

	// Verify all currencies can be retrieved by their numeric code
	for code, currency := range CurrencyMap {
		gotCurrency, gotCode, err := GetCurrencyByNumericCode(currency.NumericCode)
		if err != nil {
			t.Errorf("getCurrencyByNumericCode(%s) unexpected error for currency %s: %v", currency.NumericCode, code, err)
			continue
		}
		if gotCode != code {
			t.Errorf("getCurrencyByNumericCode(%s) code = %s, want %s", currency.NumericCode, gotCode, code)
		}
		if gotCurrency.NumericCode != currency.NumericCode {
			t.Errorf("getCurrencyByNumericCode(%s) NumericCode = %s, want %s", currency.NumericCode, gotCurrency.NumericCode, currency.NumericCode)
		}
	}
}
