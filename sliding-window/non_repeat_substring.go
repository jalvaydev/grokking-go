package slidingwindow

import "math"

func FindLongestNonRepeatSubstring(str string) int {
	windowStart := 0
	maxLength := 0
	charIndexMap := make(map[byte]int)

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		// save current iteration char
		lastChar := str[windowEnd]

		// check if lastChar is already in the index map
		if idx, ok := charIndexMap[lastChar]; ok {
			windowStart = int(math.Max(float64(windowStart), float64(idx+1)))
		}
		charIndexMap[lastChar] = windowEnd

		maxLength = int(math.Max(float64(maxLength), float64(windowEnd-windowStart+1)))
	}

	return maxLength
}
