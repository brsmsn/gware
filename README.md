# gware [![Build Status](https://travis-ci.org/brsmsn/gware.svg?branch=master)](https://travis-ci.org/brsmsn/gware) [![Go Report Card](https://goreportcard.com/badge/github.com/brsmsn/gware)](https://goreportcard.com/report/github.com/brsmsn/gware)

gware (previously gopherware) is a [diceware](http://world.std.com/~reinhold/diceware.html) passphrase generator implemented in pure go. gware can except any wordlist that follows the same format as the standard diceware list (take a look at [worldlists](https://github.com/brsmsn/gware/tree/master/test/_worldlists)).

## Is this Safe?

gware's number generator uses go's `crypto/rand` package which is a cryptographically secure pseudo random number generator (CSPRNG), the true nature of go's CSPRNG is actually a wrapper for your operating system's CSPRNG. Numbers generated from a CSPRNG are suitable for many cryptographic operations. Knowing this is it safe? Yes! If you ultra paranoid stick with the standard method of a set of casino die and paper.

## Installation

* `git clone git@github.com:brsmsn/gware.git`
* `go build cmd/gware/gware.go`

## Usage
```
Usage : ./gware [-h] [-l N] [-e N] [-o] [-b] wordList
Parameters:
        -h                 Help, get cli parameters
        -l N               Generate passphrases with N amount of words (Default is 7 words)
        -e N               Extend number of generated passphrases to N passphrases (Default is 10 passphrases)
        -n                 Number only, output only diceware numbers (a diceware number is a number from 11111 to 66666)
        -b                 Block format only, each line corresponds to a passphrase.
```
Passphrases are outputed to stdout.



