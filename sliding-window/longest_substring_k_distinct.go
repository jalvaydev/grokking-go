package slidingwindow

import "math"

func FindLongestSubstringWithKDistinct(str string, k int) int {
	charFrequency := make(map[byte]int)
	maxLength := 0
	windowStart := 0

	for windowEnd := 0; windowEnd < len(str); windowEnd++ {
		charFrequency[str[windowEnd]] += 1
		// charFrequency contains more than k distinct characters
		// shrink the window until charFrequency = k
		for len(charFrequency) > k {
			charFrequency[str[windowStart]] -= 1
			if charFrequency[str[windowStart]] == 0 {
				delete(charFrequency, str[windowStart])
			}
			windowStart++
		}

		maxLength = int(math.Max(float64(maxLength), float64(windowEnd-windowStart+1)))
	}

	return maxLength
}
