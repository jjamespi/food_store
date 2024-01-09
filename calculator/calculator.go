package calculator

var (
	menu = map[string]float64{
		"Red set":    50,
		"Green set":  40,
		"Blue set":   30,
		"Yellow set": 50,
		"Pink set":   80,
		"Purple set": 90,
		"Orange set": 120,
	}
	discountDoubleOrder = map[string]bool{
		"Orange set": true,
		"Pink set":   true,
		"Green set":  true,
	}
)

type member struct {
	isMember bool
}

func HaveMember(isMember bool) *member {
	return &member{isMember}
}

func (m *member) CalculatePrice(order map[string]int) float64 {
	var (
		result     float64
		priceOrder float64
	)

	for item, quantity := range order {
		price, exists := menu[item]
		if exists {
			if discountDoubleOrder[item] && quantity >= 2 {
				priceOrder = price * float64(int(quantity/2)*2)
				priceOrder -= priceOrder * 0.05

				if quantity%2 != 0 {
					priceOrder += price * float64(1)
				}
				result += priceOrder
			} else {
				result += price * float64(quantity)
			}
		}
	}

	if m.isMember {
		result -= result * 0.1
	}

	return result
}
