package main

import (
	"fmt"
	"food_store/calculator"
)

func main() {
	calculator := calculator.HaveMember(true)
	order := map[string]int{
		"Red set":   1,
		"Green set": 3,
	}

	totalPrice := calculator.CalculatePrice(order)
	fmt.Printf("Total Price: %.2f THB\n", totalPrice)

}
