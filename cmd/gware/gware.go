package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/brsmsn/gware/pkg/diceutil"
	"github.com/brsmsn/gware/pkg/filefmt"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "./%s [-h] [-l N] [-e N] wordList\n", os.Args[0])
		flag.PrintDefaults()
	}

	numWords := flag.Int("l", 7, "Generate passphrases with N amount of words")
	numPass := flag.Int("e", 10, "Extend number of generated passphrases to N passphrases")

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "\nNo word list specified\n\n\n")
		flag.Usage()
		os.Exit(1)
	}

	phrases := generatePhrases(*numPass, *numWords, args[0])
	printResults(phrases)

	os.Exit(0)
}

func printResults(res []string) {
	intro := "The following " + strconv.Itoa(len(res)) + " passphrase(s) were/was generated on " + time.Now().String() +
		"\nPassphrases should be copied as is (including spaces) to ensure proper amount of entropy \nhas been generated"

	fmt.Println(intro)
	fmt.Println(" ")
	for i, val := range res {
		fmt.Println("********* Passphrase " + strconv.Itoa(i+1) + " *********")
		fmt.Println(" ")
		fmt.Println(val)
		fmt.Println(" ")
		fmt.Println("********************************")
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
