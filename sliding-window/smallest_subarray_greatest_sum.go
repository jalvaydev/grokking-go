package slidingwindow

import (
	"math"
)

func FindSmallestSubarrayWithGreatestSum(s int, arr []int) int {
	smallestSubarray := math.MaxInt
	var windowSum, windowStart int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd]
		// shrink the window until s is >= to the sum
		for windowSum >= s {
			smallestSubarray = int(math.Min(float64(smallestSubarray), float64(windowEnd-windowStart+1)))
			windowSum -= arr[windowStart]
			windowStart++
		}
	}

	return smallestSubarray
}
