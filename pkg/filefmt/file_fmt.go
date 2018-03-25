// Package filefmt implements IO opperations dealing with specific format for
// diceware wordlists.
package filefmt

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// formatLine formats a given line in the wordlist to a numeric value (key) and a word
// value.
func formatLine(line string) (int, string) {
	// some file maybe formatted with a tab or space for thw whitespace between
	//key and word.
	stringSlice := strings.Fields(line)

	key, err := strconv.Atoi(stringSlice[0])
	if err != nil {
		log.Fatal(err)
	}

	return key, stringSlice[1]
}

// isValidLine returns a true if the line is has 5 digits, N whitespace and M chracters
func isValidLine(line string) bool {
	validRegex := regexp.MustCompile(`^\d{5}\s+.+$`)
	return validRegex.MatchString(line)
}

// LoadWordList loads fileName into memory. The diceware word list is stored as a map
// in memory. The 5 digit int in the file is the key of the map with word being the value.
func LoadWordList(fileName string) map[int]string {

	data, err := os.Open(fileName)
	defer data.Close()
	if err != nil {
		log.Fatal(err)
	}

	wordlist := make(map[int]string)

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		if !isValidLine(line) {
			continue
		}
		key, word := formatLine(line)
		wordlist[key] = word
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wordlist
}
