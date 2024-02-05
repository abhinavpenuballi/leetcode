package main

func finalPrices(prices []int) []int {
	finalPrices := []int{}

	for i, price := range prices {
		discount := 0

		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= price {
				discount = prices[j]
				break
			}
		}

		finalPrices = append(finalPrices, price-discount)
	}

	return finalPrices
}
