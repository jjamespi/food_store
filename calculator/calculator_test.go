package calculator_test

import (
	"food_store/calculator"
	"testing"
)

func TestCalculatorIsMemberAndDoubleOrder(t *testing.T) {
	want := ((1 * 50) + ((2 * 40) - (2 * 40 * 0.05)))
	want -= want * 0.1
	calculator := calculator.HaveMember(true)
	order := map[string]int{
		"Red set":   1,
		"Green set": 2,
	}
	totalPrice := calculator.CalculatePrice(order)
	if want != totalPrice {
		t.Errorf("totalPrice should be %.2f  but have %.2f", want, totalPrice)
	}
}

func TestCalculatorIsNotMemberAndDoubleOrder(t *testing.T) {
	want := ((1 * 50) + ((2 * 40) - (2 * 40 * 0.05)))
	calculator := calculator.HaveMember(false)
	order := map[string]int{
		"Red set":   1,
		"Green set": 2,
	}
	totalPrice := calculator.CalculatePrice(order)
	if want != totalPrice {
		t.Errorf("totalPrice should be %.2f  but have %.2f", want, totalPrice)
	}
}

func TestCalculatorIsMemberAndDoubleOrderIs3GreenSet(t *testing.T) {
	want := ((1 * 50) + ((2 * 40) - (2 * 40 * 0.05)) + (1 * 40))
	want -= want * 0.1
	calculator := calculator.HaveMember(true)
	order := map[string]int{
		"Red set":   1,
		"Green set": 3,
	}
	totalPrice := calculator.CalculatePrice(order)
	if want != totalPrice {
		t.Errorf("totalPrice should be %.2f  but have %.2f", want, totalPrice)
	}
}

func TestCalculatorIsNotMemberAndDoubleOrderIs3GreenSet(t *testing.T) {
	want := ((1 * 50) + ((2 * 40) - (2 * 40 * 0.05)) + (1 * 40))
	calculator := calculator.HaveMember(false)
	order := map[string]int{
		"Red set":   1,
		"Green set": 3,
	}
	totalPrice := calculator.CalculatePrice(order)
	if want != totalPrice {
		t.Errorf("totalPrice should be %.2f  but have %.2f", want, totalPrice)
	}
}
