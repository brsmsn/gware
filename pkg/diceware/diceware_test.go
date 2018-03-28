package diceware

import (
	"testing"

	"github.com/brsmsn/gware/pkg/filefmt"
)

func TestGeneratePassphraseNums(t *testing.T) {
	str, err := GeneratePassphraseNums(1, 1)
	if err != nil {
		t.Error(err)
	}

	if len(str) != 1 || str[0] == "" {
		t.Error("Failed to generate list")
	}
}

func TestGeneratePassphrases(t *testing.T) {
	testList1 := "../../test/worldlists/eff_large_wordlist.txt"
	words := filefmt.LoadWordList(testList1)
	phrases, err := GeneratePassphrases(1, 1, words)
	if err != nil {
		t.Error(err)
	}

	if len(phrases) != 1 || phrases[0] == "" {
		t.Error("Failed to generate list")
	}
}

func TestFlatten(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	res, err := flatten(arr)
	if err != nil {
		t.Error(err)
	}

	if res != 12345 {
		t.Error("Failed to flatten number")
	}
}
