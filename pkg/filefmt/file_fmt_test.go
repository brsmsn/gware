package filefmt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadWordList(t *testing.T) {
	testList1 := "../../_worldlists/eff_large_wordlist.txt"
	testList2 := "../../_worldlists/diceware.wordlist.asc"
	wordList1 := LoadWordList(testList1)
	wordList2 := LoadWordList(testList2)

	key1 := 11111
	key2 := 12341
	key3 := 55423

	//first = testlist1, second = testlist2
	expected1 := [2]string{"abacus", "a"}
	expected2 := [2]string{"army", "aok"}
	expected3 := [2]string{"spew", "stall"}

	//test eff wordlist
	if expected1[0] != wordList1[key1] || expected2[0] != wordList1[key2] || expected3[0] != wordList1[key3] {
		t.Errorf("Map for eff word list incorrect")
		fmt.Println(wordList1)
	}

	//test origninal wordlist
	if expected1[1] != wordList2[key1] || expected2[1] != wordList2[key2] || expected3[1] != wordList2[key3] {
		t.Errorf("Map for original word list incorrect")
		fmt.Println(wordList1)
	}
}

func TestIsValidLine(t *testing.T) {
	correctCase := [3]string{"66323	wikipedia", "66323  wikipedia", "66323		wikipedia"}
	wrongCase := [4]string{"66323wikipedia", "66323iki pedia", "6 wikipedia", "-----BEGIN PGP SIGNED MESSAGE-----"}

	for _, val := range correctCase {
		if !isValidLine(val) {
			t.Errorf("Regex failed on cases that were correct")
		}
	}

	for _, val := range wrongCase {
		if isValidLine(val) {
			t.Errorf("Regex failed on cases that were incorrect")
		}
	}
}

func TestFormatLine(t *testing.T) {
	testLine1 := "66323	wikipedia"

	testLine1Key := 66323
	testLine1Word := "wikipedia"

	key, word := formatLine(testLine1)

	if key != testLine1Key && word != testLine1Word {
		t.Errorf("Failed to parse line")
	}

	fmt.Println("******** Start of Debugging values ********")
	fmt.Printf("Line to be Formated was: %s", testLine1)
	fmt.Printf("\nValues: key:%d, word:%s", key, word)
	fmt.Printf("\nKey is a %s", reflect.TypeOf(key))
	fmt.Printf("\nWord is a %s\n", reflect.TypeOf(word))
	fmt.Println("******** End of Debugging values ********")
}
