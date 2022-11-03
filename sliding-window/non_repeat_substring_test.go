package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/jalvaydev/grokking-go/sliding-window"
)

type LongestNonRepeatSubstringTest struct {
	str      string
	expected int
}

var LongestNonRepeatSubstringTests = []LongestNonRepeatSubstringTest{
	{"aabccbb", 3},
	{"abbbb", 2},
	{"abccde", 3},
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"abba", 2},
}

func TestFindLongestNonRepeatSubstringTest(t *testing.T) {
	for _, test := range LongestNonRepeatSubstringTests {
		if output := slidingwindow.FindLongestNonRepeatSubstring(test.str); output != test.expected {
			t.Errorf("Output %v not equal to expected %v | case: %v", output, test.expected, test)
		}
	}
}
