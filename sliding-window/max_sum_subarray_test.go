package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/jalvaydev/grokking-go/sliding-window"
)

type maxSubArrayTest struct {
	k        int
	arr      []int
	expected int
}

var maxSubArrayTests = []maxSubArrayTest{
	{3, []int{2, 1, 5, 1, 3, 2}, 9},
	{2, []int{2, 3, 4, 1, 5}, 7},
}

func TestFindMaxSumSubArray(t *testing.T) {
	for _, test := range maxSubArrayTests {
		if output := slidingwindow.FindMaxSumSubArray(test.k, test.arr); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
