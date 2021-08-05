package kata02

import (
	"testing"
)

func TestExecute(t *testing.T) {
	assertTest(t, Execute(3, []int{}), -1)
	assertTest(t, Execute(3, []int{1}), -1)
	assertTest(t, Execute(1, []int{1}), 0)

	assertTest(t, Execute(1, []int{1, 3, 5}), 0)
	assertTest(t, Execute(3, []int{1, 3, 5}), 1)
	assertTest(t, Execute(5, []int{1, 3, 5}), 2)
	assertTest(t, Execute(0, []int{1, 3, 5}), -1)
	assertTest(t, Execute(2, []int{1, 3, 5}), -1)
	assertTest(t, Execute(4, []int{1, 3, 5}), -1)
	assertTest(t, Execute(6, []int{1, 3, 5}), -1)

	assertTest(t, Execute(1, []int{1, 3, 5, 7}), 0)
	assertTest(t, Execute(3, []int{1, 3, 5, 7}), 1)
	assertTest(t, Execute(5, []int{1, 3, 5, 7}), 2)
	assertTest(t, Execute(7, []int{1, 3, 5, 7}), 3)
	assertTest(t, Execute(0, []int{1, 3, 5, 7}), -1)
	assertTest(t, Execute(2, []int{1, 3, 5, 7}), -1)
	assertTest(t, Execute(4, []int{1, 3, 5, 7}), -1)
	assertTest(t, Execute(6, []int{1, 3, 5, 7}), -1)
	assertTest(t, Execute(8, []int{1, 3, 5, 7}), -1)
}

func assertTest(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("Incorrect result, got: %d, expected: %d.", got, expected)
	}
}
