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
	numOnly := flag.Bool("n", false, "Number only, output only diceware numbers (a diceware number is a number from 11111 to 66666)")
	blockFmt := flag.Bool("b", false, "Block format only, each line corresponds to a passphrase")

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "\nNo word list specified\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *numOnly {
		phrases := generatePhrases(*numPass, *numWords, nil)
		printResults(phrases, *blockFmt)

	} else {
		words := filefmt.LoadWordList(args[0])
		phrases := generatePhrases(*numPass, *numWords, words)
		printResults(phrases, *blockFmt)
	}

	os.Exit(0)
}

func printResults(res []string, toblock bool) {
	intro := "The following " + strconv.Itoa(len(res)) + " passphrase(s) were/was generated on " + time.Now().String() +
		", Passphrases should be copied as is (including spaces) to ensure proper amount of entropy has been generated"

	fmt.Println(intro)
	fmt.Println(" ")
	for i, val := range res {
		if toblock {
			fmt.Println(val)
		} else {
			fmt.Println("******************  Passphrase " + strconv.Itoa(i+1) + " ******************")
			fmt.Println(val)
			fmt.Println(" ")
		}
	}
}

func generatePhrases(nPhrases, nWords int, words map[int]string) []string {
	passphrases := make([]string, nPhrases)

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
				if words != nil {
					passph = words[flatten(roll)]
				} else {
					passph = strconv.Itoa(flatten(roll))
				}
			} else {
				if words != nil {
					passph = passph + " " + words[flatten(roll)]
				} else {
					passph = passph + " " + strconv.Itoa(flatten(roll))
				}
			}
		}

		passphrases[i] = passph
	}

	return passphrases
}
