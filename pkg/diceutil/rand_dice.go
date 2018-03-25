//Package diceutil contains utility functions for simulating die rolls.
package diceutil

import (
	"crypto/rand"
	"fmt"
)

// getRandom gets cryptographically secured random numbers from rand.Read().
func getRandom() ([]byte, error) {
	b := make([]byte, 1)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return b, nil
}

// RollDice simulates n die rolls. Where n is the amount of dice to
// roll. Returns a slice of integer(s) representing n die roll(s) and an error if
// an exception has been thrown, nil otherwise.
func RollDice(n int) ([]int, error) {
	rolls := make([]int, n)
	for i := range rolls {
		val, err := getRandom()
		if err != nil {
			return nil, err
		}
		//we do mod6 + 1 to shift the numbers from 1 to 6 instead
		//of 0 to 5. A dice does not have a 0 side.
		rolls[i] = int(val[0])%6 + 1
	}
	return rolls, nil
}
