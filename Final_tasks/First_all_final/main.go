package main

const (
	minPrice = 99.00
	maxPrice = 20000
)

func ApplyPriceLimits(price float64) float64 {
	switch minPrice <= price && price <= maxPrice {
	case true:
		return price
	case false:
		if price < minPrice {
			return minPrice
		} else {
			return maxPrice
		}
	}
	return 0.0
}

func main() {

}