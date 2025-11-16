package goodmoney

import (
	"encoding/json"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = New(100.50, USD)
	}
}

func BenchmarkNewZero(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = NewZero(USD)
	}
}

func BenchmarkAmount(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Amount()
	}
}

func BenchmarkCurrency(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Currency()
	}
}

func BenchmarkAbsolute(b *testing.B) {
	m, _ := New(-100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Absolute()
	}
}

func BenchmarkNegative(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Negative()
	}
}

func BenchmarkIsZero(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.IsZero()
	}
}

func BenchmarkIsPositive(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.IsPositive()
	}
}

func BenchmarkIsNegative(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.IsNegative()
	}
}

func BenchmarkCompare(b *testing.B) {
	m1, _ := New(100.50, USD)
	m2, _ := New(50.25, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m1.Compare(m2)
	}
}

func BenchmarkEquals(b *testing.B) {
	m1, _ := New(100.50, USD)
	m2, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m1.Equals(m2)
	}
}

func BenchmarkGreaterThan(b *testing.B) {
	m1, _ := New(100.50, USD)
	m2, _ := New(50.25, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m1.GreaterThan(m2)
	}
}

func BenchmarkLessThan(b *testing.B) {
	m1, _ := New(50.25, USD)
	m2, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m1.LessThan(m2)
	}
}

func BenchmarkMultiply(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Multiply(2, 3)
	}
}

func BenchmarkDivide(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Divide(2, 3)
	}
}

func BenchmarkRound(b *testing.B) {
	m, _ := New(100.55, USD)
	scheme := RoundHalfUp
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Round(&scheme)
	}
}

func BenchmarkAdd(b *testing.B) {
	m1, _ := New(100.50, USD)
	m2, _ := New(50.25, USD)
	m3, _ := New(25.75, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Add(m1, m2, m3)
	}
}

func BenchmarkSubtract(b *testing.B) {
	m, _ := New(100.50, USD)
	m1, _ := New(50.25, USD)
	m2, _ := New(25.75, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Subtract(m1, m2)
	}
}

func BenchmarkAllocate(b *testing.B) {
	m, _ := New(100.00, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Allocate(1, 2, 3, 4)
	}
}

func BenchmarkAllocateByPercentage(b *testing.B) {
	m, _ := New(100.00, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.AllocateByPercentage(25.0, 25.0, 25.0, 25.0)
	}
}

func BenchmarkString(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.String()
	}
}

func BenchmarkMarshalJSON(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.MarshalJSON()
	}
}

func BenchmarkUnmarshalJSON(b *testing.B) {
	data := []byte(`{"amount":100.50,"currency":"USD"}`)
	var m Money
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.UnmarshalJSON(data)
	}
}

func BenchmarkValue(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = m.Value()
	}
}

func BenchmarkScan(b *testing.B) {
	data := []byte(`{"amount":100.50,"currency":"USD"}`)
	var m Money
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Scan(data)
	}
}

func BenchmarkMarshalUnmarshalRoundTrip(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, _ := m.MarshalJSON()
		var unm Money
		_ = unm.UnmarshalJSON(data)
	}
}

func BenchmarkValueScanRoundTrip(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val, _ := m.Value()
		var scanned Money
		_ = scanned.Scan(val)
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	m, _ := New(100.50, USD)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(m)
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	data := []byte(`{"amount":100.50,"currency":"USD"}`)
	var m Money
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &m)
	}
}
