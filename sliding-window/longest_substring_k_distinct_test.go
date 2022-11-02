package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/jalvaydev/grokking-go/sliding-window"
)

type longestSubstringWithKDistinctTest struct {
	k        int
	str      string
	expected int
}

var longestSubstringWithKDistinctTests = []longestSubstringWithKDistinctTest{
	{2, "araaci", 4},
	{1, "araaci", 2},
	{3, "cbbebi", 5},
}

func TestFindLongestSubstringWithKDistinctTest(t *testing.T) {
	for _, test := range longestSubstringWithKDistinctTests {
		if output := slidingwindow.FindLongestSubstringWithKDistinct(test.str, test.k); output != test.expected {
			t.Errorf("Output %v not equal to expected %v | case: %v", output, test.expected, test)
		}
	}
}
