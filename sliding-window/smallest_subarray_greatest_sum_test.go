package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/jalvaydev/grokking-go/sliding-window"
)

type smallestSubarrayWithGreatestSumTest struct {
	s        int
	arr      []int
	expected int
}

var smallestSubarrayWithGreatestSumTests = []smallestSubarrayWithGreatestSumTest{
	{7, []int{2, 1, 5, 2, 3, 2}, 2},
	{7, []int{2, 1, 5, 2, 8}, 1},
	{8, []int{3, 4, 1, 1, 6}, 3},
}

func TestFindSmallestSubarrayWithGreatestSum(t *testing.T) {
	for _, test := range smallestSubarrayWithGreatestSumTests {
		if output := slidingwindow.FindSmallestSubarrayWithGreatestSum(test.s, test.arr); output != test.expected {
			t.Errorf("Output %v not equal to expected %v | case: %v", output, test.expected, test)
		}
	}
}
