package main

import (
	"fmt"

	"github.com/brsmsn/gware/pkg/diceutil"
)

func main() {
	value, _ := diceutil.RollDice(6)

	fmt.Println(value)
}
