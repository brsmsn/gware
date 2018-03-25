package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/brsmsn/gware/pkg/diceutil"
	"github.com/brsmsn/gware/pkg/filefmt"
)

func main() {
	numWords := flag.Int("l", 7, "Generate passphrases with N amount of words (Default is 7 words)")
	numPass := flag.Int("e", 10, " Extend number of generated passphrases to N passphrases (Default is 10 passphrases)")

	flag.Parse()
	args := flag.Args()

	phrases := generatePhrases(*numPass, *numWords, args[1])

	printResults(phrases)
}

func printResults(res []string) {
	intro := "The following " + string(len(res)) + " passphrase(s) were generated on " + time.Now().String()

	fmt.Println(intro)
	for i, val := range res {
		fmt.Println("********* Passphrase" + string(i) + " *********")
		fmt.Println(val)
		fmt.Println("******************************")
	}
}

func generatePhrases(nPhrases, nWords int, wordlist string) []string {
	passphrases := make([]string, nPhrases)
	words := filefmt.LoadWordList(wordlist)

	for i := 0; i < nPhrases; i++ {
		var passph string
		for k := 0; k < nWords; k++ {
			//diceware require 5 die
			roll, err := diceutil.RollDice(5)

			if err != nil {
				fmt.Println("error:", err)
			}

			//converts array to string
			flatten := func(arr []int) int {
				var res []string
				for _, val := range arr {
					n := strconv.Itoa(val)
					res = append(res, n)
				}

				final, _ := strconv.Atoi(strings.Join(res, ""))

				return final
			}

			if k == 0 {
				passph = words[flatten(roll)]
			} else {
				passph = passph + " " + words[flatten(roll)]
			}
		}

		passphrases[i] = passph
	}

	return passphrases
}
