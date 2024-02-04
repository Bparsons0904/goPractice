package runes

import (
	"log"
	"strings"
	"unicode"
)

func main() {
	test := "   kajhekjas fkjl;al;kj fakjnkasnfjasnflisabas clcdksajflkas f   "
	test2 := "This is more of a standard list of words"

	log.Println("Test1:", test)
	log.Println("Test2:", test2)

	filteredWhiteSpace(test)
	filterSymbols(test)
	sortedString(test)
	countSorted(test)
}

func countSorted(s string) {
	sortedString := sortedString(s)

	countMap := make(map[string]int)

	for _, rune := range sortedString {
		countMap[string(rune)]++
	}

	log.Println(countMap)
}

func sortedString(s string) string {
	var runes []rune
	for _, rune := range filterSymbols(s) {
		runes = append(runes, rune)
	}

	for i := 0; i < len(runes)-1; i++ {
		for k := i + 1; k < len(runes); k++ {
			if runes[i] > runes[k] {
				temp := runes[i]
				runes[i] = runes[k]
				runes[k] = temp
			}
		}
	}

	log.Println(runes)
	log.Println(string(runes))
	return string(runes)
}

func filterSymbols(s string) string {
	var result []rune
	for _, rune := range s {
		if unicode.IsLetter(rune) {
			result = append(result, rune)
		}
	}

	log.Println(string(result))
	return string(result)
}

func filteredWhiteSpace(s string) {
	filtered := strings.Trim(s, " ")
	log.Println("Filtered:", filtered)
}
