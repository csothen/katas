package kata19

import (
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	assertTest(t, Execute("cat", "dog"), []string{"cat", "cot", "cog", "dog"})
	assertTest(t, Execute("lead", "gold"), []string{"lead", "load", "goad", "gold"})
	assertTest(t, Execute("ruby", "code"), []string{"ruby", "rubs", "robs", "rods", "rode", "code"})
}

func assertTest(t *testing.T, got, expected []string) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incorrect result, got: %v, expected: %v.", got, expected)
	}
}
