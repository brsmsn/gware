package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestGeneratePhrases(t *testing.T) {
	list := generatePhrases(5, 3, "../../_worldlists/eff_large_wordlist.txt")

	fmt.Println(list)

	if len(list) != 5 {
		t.Errorf("Did not generate the right amount of passphrases")
		fmt.Println(list)
	}

	for _, val := range list {
		if strings.Count(val, " ") != 2 {
			t.Errorf("Number of words do not match")
		}
	}
}
