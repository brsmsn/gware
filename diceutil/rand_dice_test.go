package diceutil

import (
	"testing"
)

// TestGetRandom, test if
func TestGetRandom(t *testing.T) {
	val, err := getRandom()
	if err != nil || val == nil {
		t.Errorf("Error failed to generate random number")
	}
}

// TestRollDice test for simulated die rolls,
func TestRollDice(t *testing.T) {
	val, err := RollDice(6)
	if err == nil && len(val) != 6 {
		t.Errorf("Error failed to generate die roll")
	}
	for i, num := range val {
		if num <= 0 || num > 6 {
			t.Errorf("Dice number %d is out range, expected a number between"+
				"1 and 6", i)
		}
	}
}
