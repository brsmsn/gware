package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/brsmsn/gware/pkg/diceware"
	"github.com/brsmsn/gware/pkg/filefmt"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "./%s [-h] [-l N] [-e N] wordList\n", os.Args[0])
		flag.PrintDefaults()
	}

	numWords := flag.Int("l", 7, "Generate passphrases with N amount of words")
	numPass := flag.Int("e", 10, "Extend number of generated passphrases to N passphrases")
	numOnly := flag.Bool("n", false, "Number only, output only diceware numbers (a diceware number is a number from 11111 to 66666), this flag will ignore wordlist")
	blockFmt := flag.Bool("b", false, "Block format only, each line corresponds to a passphrase")

	flag.Parse()
	args := flag.Args()

	if flag.NArg() == 0 && !*numOnly {
		fmt.Fprintf(os.Stderr, "\nNo word list specified\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *numPass <= 0 || *numWords <= 0 {
		fmt.Fprintf(os.Stderr, "\nCan not specify negative numbers\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *numOnly {
		phrases, err := diceware.GeneratePassphraseNums(*numPass, *numWords)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		printResults(phrases, *blockFmt)

	} else {
		words := filefmt.LoadWordList(args[0])
		phrases, err := diceware.GeneratePassphrases(*numPass, *numWords, words)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
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
