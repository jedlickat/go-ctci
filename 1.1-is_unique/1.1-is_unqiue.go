// 1.1 - is_unique
package main

import "fmt"

// is_unique returns true iff the str does not contain duplicit characters.
// Internal implementation uses map of runes for that purpose.
func is_unique(str string) bool {
	var charmap map[rune]bool

	charmap = make(map[rune]bool)

	for _, r := range []rune(str) {
		if charmap[r] == true {
			return false
		} else {
			charmap[r] = true
		}

	}
	return true
}

func main() {
	var input string
	var result string

	fmt.Scanln(&input)
	if is_unique(input) {
		result = "unique"
	} else {
		result = "non-unique"
	}

	fmt.Printf("String is %s\n", result)
}
