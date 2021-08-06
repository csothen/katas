package kata19

import (
	"reflect"
	"testing"
)

func TestExecute_Cat_Dog(t *testing.T) {
	assertTest(t, Execute("cat", "dog"), []string{"cat", "cot", "cog", "dog"})
}

func TestExecute_Lead_Gold(t *testing.T) {
	assertTest(t, Execute("lead", "gold"), []string{"lead", "load", "goad", "gold"})
}

func TestExecute_Ruby_Code(t *testing.T) {
	assertTest(t, Execute("ruby", "code"), []string{"ruby", "rube", "robe", "rode", "code"})
}

func TestExecute_Loric_Poled(t *testing.T) {
	assertTest(t, Execute("loric", "poled"), []string{"loric", "loris", "lores", "pores", "poles", "poled"})
}

func TestExecute_Polemizes_Syntonize(t *testing.T) {
	assertTest(t, Execute("brushing", "cheating"), []string{"brushing", "crushing", "crusting", "cresting", "creating", "cheating"})
}

func assertTest(t *testing.T, got, expected []string) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incorrect result, got: %v, expected: %v.", got, expected)
	}
}
