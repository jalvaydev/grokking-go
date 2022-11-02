package slidingwindow

import "math"

func FindMaxSumSubArray(k int, arr []int) int {
	var windowSum, maxSum int
	var windowStart int

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		// add to windowSum every iteration
		windowSum += arr[windowEnd]

		// enter once k elements are in the sum
		if windowEnd >= k-1 {
			// find if current max is greater than current window
			maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
			// subtract first element of k elements
			windowSum -= arr[windowStart]
			// next element of k elements k = 2
			// [2, 1, 5, 6] => [2, 1] => [1, 5]
			windowStart++
		}
	}

	return maxSum
}
