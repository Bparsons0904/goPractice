package main

import (
	"fmt"
	"strings"
)

// TestCase struct for find and replace function.
type TestCase struct {
	Source         string
	Search         string
	Replace        string
	ExpectedOutput string
}

// findAndReplace function placeholder (implement your solution here).
func findAndReplace(source, search, replace string) string {
	index := strings.Index(source, search)
	if index == -1 {
		return source
	}

	updatedString := source[:index] + replace + source[index+len(search):]

	return findAndReplace(updatedString, search, replace)
}

// return strings.ReplaceAll(source, search, replace) // Example implementation.

func main() {
	// Slice of test cases for find and replace functionality.
	testCases := []TestCase{
		{"Hello, world!", "world", "Gopher", "Hello, Gopher!"},
		{"foo bar foo", "foo", "baz", "baz bar baz"},
		{"case-sensitive", "case", "CASE", "CASE-sensitive"},
		{"repeat repeat repeat", "repeat", "again", "again again again"},
		{"", "empty", "test", ""},
		{"no match here", "xyz", "abc", "no match here"},
		{"leading space", "leading", "trailing", "trailing space"},
		{"trailing space", "space", "SPACES", "trailing SPACES"},
		{"case-insensitive", "Case", "CASE", "case-insensitive"},
		{"123-456-789", "-", ":", "123:456:789"},
	}

	// Loop through each test case.
	for _, testCase := range testCases {
		result := findAndReplace(testCase.Source, testCase.Search, testCase.Replace) // Call your find and replace function.
		if result == testCase.ExpectedOutput {
			fmt.Printf("PASS: For input \"%s\" with search \"%s\" and replace \"%s\", expected \"%s\" and got \"%s\".\n", testCase.Source, testCase.Search, testCase.Replace, testCase.ExpectedOutput, result)
		} else {
			fmt.Printf("FAIL: For input \"%s\" with search \"%s\" and replace \"%s\", expected \"%s\" but got \"%s\".\n", testCase.Source, testCase.Search, testCase.Replace, testCase.ExpectedOutput, result)
		}
	}
}
