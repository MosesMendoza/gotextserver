package text

import (
	"bytes"
	"regexp"
	"strings"
)

// Deduplicate returns a string containing only one
// instance of any individual rune in the submitted string.
func Deduplicate(text string) string {
	// strings are read-only in Go so any concatenation generates a new string
	// instead use a buffer for efficiency
	var deduplicated bytes.Buffer
	// use a set implementation to track known characters
	seenCharacters := map[rune]bool{}

	for _, character := range text {
		if !seenCharacters[character] {
			deduplicated.WriteRune(character)
			seenCharacters[character] = true
		}
	}

	return deduplicated.String()
}

// RemoveIntegers returns a substring match of all characters that are not
// integers in the supplied string
func RemoveIntegers(text string) string {
	re := regexp.MustCompile("[^0-9]")
	matched := re.FindAllString(text, -1)
	return strings.Join(matched, "")
}
