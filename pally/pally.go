package pally

import (
	"fmt"
	"log"
	"unicode"
)

// TestCase struct to hold a single test case for the palindrome checker.
type TestCase struct {
	Input          string
	ExpectedOutput bool
}

// isPalindrome function placeholder (implement your solution here).
func isPalindrome(s string) bool {
	var runes []rune

	for _, rune := range s {
		if unicode.IsLetter(rune) {
			runes = append(runes, unicode.ToLower(rune))
		}
	}

	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}

	log.Println(runes)
	return true // Placeholder return.
}

func main() {
	// Slice of test cases for palindrome checking.
	testCases := []TestCase{
		{"racecar", true},
		{"a", true},
		{"", true},
		{"hello", false},
		{"RaceCar", true},
		{"A man a plan a canal Panama", true},
		{"This is not a palindrome", false},
		{"Madam, I'm Adam", true},
		{"12321", true},
		{"1a2b3b2a1x", false},
	}

	// Loop through each test case.
	for _, testCase := range testCases {
		result := isPalindrome(testCase.Input) // Call your palindrome checking function.
		if result == testCase.ExpectedOutput {
			fmt.Printf("PASS: For input \"%s\", expected %t.\n", testCase.Input, testCase.ExpectedOutput)
		} else {
			fmt.Printf("FAIL: For input \"%s\", expected %t.\n", testCase.Input, testCase.ExpectedOutput)
		}
	}
}
