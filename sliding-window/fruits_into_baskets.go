package slidingwindow

import "math"

func FruitsIntoBaskets(fruit []byte) int {
	fruitFrequency := make(map[byte]int)
	maxFruit := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(fruit); windowEnd++ {
		fruitFrequency[fruit[windowEnd]]++

		for len(fruitFrequency) > 2 {
			fruitFrequency[fruit[windowStart]]--
			if fruitFrequency[fruit[windowStart]] == 0 {
				delete(fruitFrequency, fruit[windowStart])
			}
			windowStart++
		}

		maxFruit = int(math.Max(float64(maxFruit), float64(windowEnd)-float64(windowStart)+1))
	}

	return maxFruit
}
