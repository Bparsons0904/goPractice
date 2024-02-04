package main

import (
	"fmt"
	"strings"
)

// LogEntry struct to hold the extracted log level and message.
type LogEntry struct {
	Level   string
	Message string
}

// TestCase struct for holding each test scenario.
type TestCase struct {
	LogLine        string
	ExpectedOutput LogEntry
}

// parseLogEntry function to be implemented.
func parseLogEntry(logLine string) LogEntry {
	sections := strings.SplitN(logLine, " ", 4)

	level := strings.Trim(sections[0], "[]")

	return LogEntry{
		Level:   level,
		Message: sections[3],
	}
}

func main() {
	// Test cases.
	testCases := []TestCase{
		{
			LogLine:        "[ERROR] 2024-02-03 15:04:05 Some error occurred in the system: Error details here.",
			ExpectedOutput: LogEntry{Level: "ERROR", Message: "Some error occurred in the system: Error details here."},
		},
		{
			LogLine:        "[WARN] 2024-02-03 16:15:00 Warning: System is running low on memory.",
			ExpectedOutput: LogEntry{Level: "WARN", Message: "Warning: System is running low on memory."},
		},
		{
			LogLine:        "[INFO] 2024-02-04 09:00:00 System startup completed successfully.",
			ExpectedOutput: LogEntry{Level: "INFO", Message: "System startup completed successfully."},
		},
		// Add more test cases as needed.
	}

	// Run test cases.
	for _, testCase := range testCases {
		result := parseLogEntry(testCase.LogLine)
		if result == testCase.ExpectedOutput {
			fmt.Printf("PASS: For log line \"%s\"\n", testCase.LogLine)
		} else {
			fmt.Printf("FAIL: For log line \"%s\". Expected [%s] \"%s\", got [%s] \"%s\".\n",
				testCase.LogLine, testCase.ExpectedOutput.Level, testCase.ExpectedOutput.Message, result.Level, result.Message)
		}
	}
}
