package kata19

import (
	"reflect"
	"strings"
	"testing"
)

func TestExecute_SameStartAndEnd(t *testing.T) {
	chain, err := Execute("cat", "cat")
	expected := []string{"cat"}

	assertTest(t, chain, expected)
	assertTest(t, err, nil)
}

func TestExecute_DifferentLengthsError(t *testing.T) {
	chain, err := Execute("cat", "gold")
	expected := ErrDifferentLengths

	assertTest(t, len(chain), 0)
	assertTest(t, err, expected)
}

func TestExecute_StartWordNotInDictionaryError(t *testing.T) {
	chain, err := Execute("arduino", "astound")

	assertTest(t, len(chain), 0)
	assertTest(t, strings.Contains(err.Error(), ErrWordNotInDictionary.Error()), true)
}

func TestExecute_EndWordNotInDictionaryError(t *testing.T) {
	chain, err := Execute("astound", "arduino")

	assertTest(t, len(chain), 0)
	assertTest(t, strings.Contains(err.Error(), ErrWordNotInDictionary.Error()), true)
}

func TestExecute_ChainNotFoundError(t *testing.T) {
	chain, err := Execute("astrally", "washered")
	expected := ErrChainNotFound

	assertTest(t, len(chain), 0)
	assertTest(t, err, expected)
}

func TestExecute_CatToDog(t *testing.T) {
	chain, err := Execute("cat", "dog")
	expected := []string{"cat", "cot", "cog", "dog"}

	assertTest(t, chain, expected)
	assertTest(t, err, nil)
}

func TestExecute_LeadToGold(t *testing.T) {
	chain, err := Execute("lead", "gold")
	expected := []string{"lead", "load", "goad", "gold"}

	assertTest(t, chain, expected)
	assertTest(t, err, nil)
}

func TestExecute_RubyToCode(t *testing.T) {
	chain, err := Execute("ruby", "code")
	expected := []string{"ruby", "rube", "robe", "rode", "code"}

	assertTest(t, chain, expected)
	assertTest(t, err, nil)
}

func TestExecute_LoricToPoled(t *testing.T) {
	chain, err := Execute("loric", "poled")
	expected := []string{"loric", "loris", "lores", "pores", "poles", "poled"}

	assertTest(t, chain, expected)
	assertTest(t, err, nil)
}

func TestExecute_PolemizesToSyntonize(t *testing.T) {
	chain, err := Execute("brushing", "cheating")
	expected := []string{"brushing", "crushing", "crusting", "cresting", "creating", "cheating"}

	assertTest(t, chain, expected)
	assertTest(t, err, nil)
}

func assertTest(t *testing.T, got, expected interface{}) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Incorrect result, got: %v, expected: %v.", got, expected)
	}
}
