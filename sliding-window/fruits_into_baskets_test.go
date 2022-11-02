package slidingwindow_test

import (
	"testing"

	slidingwindow "github.com/jalvaydev/grokking-go/sliding-window"
)

type fruitsIntoBasketsTest struct {
	fruit    []byte
	expected int
}

var fruitsIntoBasketsTests = []fruitsIntoBasketsTest{
	{[]byte{'A', 'B', 'C', 'A', 'C'}, 3},
	{[]byte{'A', 'B', 'C', 'B', 'B', 'C'}, 5},
}

func TestFruitsIntoBaskets(t *testing.T) {
	for _, test := range fruitsIntoBasketsTests {
		if output := slidingwindow.FruitsIntoBaskets(test.fruit); output != test.expected {
			t.Errorf("Output %v not equal to expected %v | case: %v", output, test.expected, test)
		}
	}
}
