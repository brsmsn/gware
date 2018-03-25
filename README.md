# gware [![Build Status](https://travis-ci.org/brsmsn/gware.svg?branch=master)](https://travis-ci.org/brsmsn/gware)

gware (previously gopherware) is a diceware passphrase generator implemented in pure go. 

## Is this Safe?

gware's number generator uses go's `crypto/rand` package which is a cryptographically secure pseudo random number generator (CSPRNG), the true nature of go's CSPRNG is actually a wrapper for your operating system's CSPRNG. Numbers generated from a CSPRNG are suitable for many cryptographic operations. Knowing this is it safe? Yes! If you ultra paranoid stick with the standard method of a set of casino dies and paper.

## Installation

* `git clone git@github.com:brsmsn/gware.git`

## Usage
```
Usage : ./gware [-h] [-l N] [-t N] [-s File] [-w List]
Parameters:
        -h                  Help, get cli parameters
        -l N                Generate passphrases with N amount of words (Default is 7 words)
        -e N                Extend number of generated passphrases to N passphrases (Default is 10 passphrases)
        -s File             Save passphrases to specified File (Default stdout)
        -w List             Use specified list for diceware wordlist
```




