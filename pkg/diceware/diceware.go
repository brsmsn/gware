//Package diceware contains definitions for the diceware algorithm
package diceware

import (
	"strconv"
	"strings"

	"github.com/brsmsn/gware/pkg/diceutil"
)

//GeneratePassphraseNums generates nPhrases worth of passphrases each containing nWords
func GeneratePassphraseNums(nPhrases, nWords int) ([]string, error) {
	return generatePhrases(nPhrases, nWords, nil)
}

//GeneratePassphrases generates  nPhrases worth of passphrases each containing nWords from a wordlist words
func GeneratePassphrases(nPhrases, nWords int, words map[int]string) ([]string, error) {
	return generatePhrases(nPhrases, nWords, words)
}

//generates passphrases
func generatePhrases(nPhrases, nWords int, words map[int]string) ([]string, error) {
	passphrases := make([]string, nPhrases)

	for i := 0; i < nPhrases; i++ {
		var passph string
		for k := 0; k < nWords; k++ {
			//diceware require 5 die
			roll, err := diceutil.RollDice(5)
			if err != nil {
				return nil, err
			}

			index, err := flatten(roll)
			if err != nil {
				return nil, err
			}

			if k == 0 {
				if words != nil {
					passph = words[index]
				} else {
					passph = strconv.Itoa(index)
				}
			} else {
				if words != nil {
					passph = passph + " " + words[index]
				} else {
					passph = passph + " " + strconv.Itoa(index)
				}
			}
		}

		passphrases[i] = passph
	}

	return passphrases, nil
}

//converts array to string
func flatten(arr []int) (int, error) {
	var res []string
	for _, val := range arr {
		n := strconv.Itoa(val)
		res = append(res, n)
	}

	final, err := strconv.Atoi(strings.Join(res, ""))
	if err != nil {
		return -1, err
	}

	return final, nil
}
