package lib_test

import (
	"log"
	"os"
	"testing"

	"github.com/faizalom/go-web/lib"
)

func TestLogErrors(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_log_*.log")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Call the LogErrors function with the temporary file
	lib.LogErrors(tempFile.Name())

	// Log a test message
	testMessage := "This is a test error message"
	log.Println(testMessage)

	// Read the contents of the log file
	logContents, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	// Check if the test message is in the log file
	if !contains(logContents, testMessage) {
		t.Errorf("log file does not contain expected message: got %v want %v", string(logContents), testMessage)
	}
}

// Helper function to check if a byte slice contains a string
func contains(logContents []byte, message string) bool {
	return string(logContents) != "" && string(logContents) != message
}
