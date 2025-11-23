package main

import (
	"fmt"
	"log"

	"github.com/nucleus-proj/goodmoney/goodmoney"
)

func main() {
	// Example: IsNegative
	fmt.Println("=== IsNegative ===")
	positive, _ := goodmoney.New(100.50, goodmoney.USD)
	negative, _ := goodmoney.New(-100.50, goodmoney.USD)
	zero, _ := goodmoney.NewZero(goodmoney.USD)

	fmt.Printf("%s is negative: %v\n", positive, positive.IsNegative())
	fmt.Printf("%s is negative: %v\n", negative, negative.IsNegative())
	fmt.Printf("%s is negative: %v\n", zero, zero.IsNegative())

	// Example: IsPositive
	fmt.Println("\n=== IsPositive ===")
	fmt.Printf("%s is positive: %v\n", positive, positive.IsPositive())
	fmt.Printf("%s is positive: %v\n", negative, negative.IsPositive())
	fmt.Printf("%s is positive: %v\n", zero, zero.IsPositive())

	// Example: IsZero
	fmt.Println("\n=== IsZero ===")
	fmt.Printf("%s is zero: %v\n", positive, positive.IsZero())
	fmt.Printf("%s is zero: %v\n", negative, negative.IsZero())
	fmt.Printf("%s is zero: %v\n", zero, zero.IsZero())

	// Example: IsValid
	fmt.Println("\n=== IsValid ===")
	fmt.Printf("%s is valid: %v\n", positive, positive.IsValid())
	invalid := goodmoney.Money{}
	fmt.Printf("Invalid money is valid: %v\n", invalid.IsValid())

	// Example: Absolute
	fmt.Println("\n=== Absolute ===")
	abs := negative.Absolute()
	fmt.Printf("Absolute of %s = %s\n", negative, abs)

	abs2 := positive.Absolute()
	fmt.Printf("Absolute of %s = %s\n", positive, abs2)

	// Example: Negative
	fmt.Println("\n=== Negative ===")
	neg := positive.Negative()
	fmt.Printf("Negative of %s = %s\n", positive, neg)

	neg2 := negative.Negative()
	fmt.Printf("Negative of %s = %s\n", negative, neg2)

	// Example: Amount
	fmt.Println("\n=== Amount ===")
	amount := positive.Amount()
	fmt.Printf("Amount of %s = %.2f\n", positive, amount)

	// Example: Round
	fmt.Println("\n=== Round ===")
	toRound, _ := goodmoney.New(100.55, goodmoney.USD)
	scheme := goodmoney.RoundHalfUp
	rounded := toRound.Round(&scheme)
	fmt.Printf("Rounded %s (HalfUp) = %s\n", toRound, rounded)

	scheme2 := goodmoney.RoundHalfDown
	rounded2 := toRound.Round(&scheme2)
	fmt.Printf("Rounded %s (HalfDown) = %s\n", toRound, rounded2)

	scheme3 := goodmoney.RoundHalfEven
	rounded3 := toRound.Round(&scheme3)
	fmt.Printf("Rounded %s (HalfEven) = %s\n", toRound, rounded3)

	// Round with nil scheme (defaults to RoundTowardZero)
	rounded4 := toRound.Round(nil)
	fmt.Printf("Rounded %s (default) = %s\n", toRound, rounded4)

	// Example: Allocate
	fmt.Println("\n=== Allocate ===")
	total, _ := goodmoney.New(100.00, goodmoney.USD)
	parts, err := total.Allocate(3, 2, 1)
	if err != nil {
		log.Fatalf("Error allocating: %v", err)
	}
	fmt.Printf("Allocated %s in 3:2:1 ratio:\n", total)
	for i, part := range parts {
		fmt.Printf("  Part %d: %s\n", i+1, part)
	}

	// Example: AllocateByPercentage
	fmt.Println("\n=== AllocateByPercentage ===")
	payment, _ := goodmoney.New(1000.00, goodmoney.USD)
	shares, err := payment.AllocateByPercentage(60.0, 25.0, 15.0)
	if err != nil {
		log.Fatalf("Error allocating by percentage: %v", err)
	}
	fmt.Printf("Allocated %s by percentages (60%%, 25%%, 15%%):\n", payment)
	for i, share := range shares {
		percentages := []float64{60.0, 25.0, 15.0}
		fmt.Printf("  Share %d (%.0f%%): %s\n", i+1, percentages[i], share)
	}

	// Example: MajorUnit and MinorUnit
	fmt.Println("\n=== MajorUnit and MinorUnit ===")
	m, _ := goodmoney.New(100.50, goodmoney.USD)
	fmt.Printf("%s\n", m)
	fmt.Printf("Major unit (dollars): %d\n", m.MajorUnit())
	fmt.Printf("Minor unit (cents): %d\n", m.MinorUnit())

	jpy, _ := goodmoney.New(1234.0, goodmoney.JPY)
	fmt.Printf("\n%s\n", jpy)
	fmt.Printf("Major unit (yen): %d\n", jpy.MajorUnit())
	fmt.Printf("Minor unit: %d\n", jpy.MinorUnit())
}
