package lib_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/faizalom/go-web/lib"
)

func TestExeTemplate(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "test_template_*.html")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	fmt.Fprintf(tempFile, "<html><head><title>{{.Title}}</title></head><body></body></html>")

	abs := filepath.Dir(tempFile.Name())
	themePath := abs + "/*.html"
	lib.TemplateParseGlob(themePath)
	if lib.Template == nil {
		t.Errorf("TemplateParseGlob failed, Template is nil")
	}

	tempFileName := filepath.Base(tempFile.Name())

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	data := map[string]string{"Title": "Test Title"}

	// Execute the template
	lib.ExeTemplate(rr, tempFileName, data)

	// Check the response code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ExeTemplate returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "<html><head><title>Test Title</title></head><body></body></html>"
	if rr.Body.String() != expected {
		t.Errorf("ExeTemplate returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
