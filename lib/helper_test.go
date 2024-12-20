package lib_test

import (
	"testing"

	"github.com/faizalom/go-web/lib"
)

func TestIsEmailValidFalse(t *testing.T) {
	actualString := lib.IsEmailValid("test")
	expectedString := false
	if actualString != expectedString {
		t.Error("Expected Result false is not same as actual Result true", expectedString, actualString)
	}
}

func TestIsEmailValidTrue(t *testing.T) {
	actualString := lib.IsEmailValid("test@test.com")
	expectedString := true
	if actualString != expectedString {
		t.Error("Expected Result false is not same as actual Result true", expectedString, actualString)
	}
}

func TestRandToken(t *testing.T) {
	token1 := lib.RandToken()
	token2 := lib.RandToken()

	if len(token1) == 0 {
		t.Error("Expected non-empty token, got empty string")
	}

	if token1 == token2 {
		t.Error("Expected different tokens, got the same token")
	}
}
