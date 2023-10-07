package betterjson

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type Config struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestUnmarshalWithComments(t *testing.T) {
	testFilesDir := "./test_files/"

	// Map to store the content of the expected files
	expectedFiles := make(map[string][]byte)

	// Read all files in the test_files directory
	files, err := os.ReadDir(testFilesDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	// Separate the files into expected and input
	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, "_expected.json") {
			content, err := os.ReadFile(filepath.Join(testFilesDir, filename))
			if err != nil {
				t.Errorf("Failed to read file %s: %v", filename, err)
				continue
			}
			testName := strings.TrimSuffix(filename, "_expected.json")
			expectedFiles[testName] = content
		}
	}

	// Process each input file and compare with the expected content
	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, "_input.json") {
			content, err := os.ReadFile(filepath.Join(testFilesDir, filename))
			if err != nil {
				t.Errorf("Failed to read file %s: %v", filename, err)
				continue
			}

			testName := strings.TrimSuffix(filename, "_input.json")
			expectedContent, exists := expectedFiles[testName]
			if !exists {
				t.Errorf("No expected file found for %s", testName)
				continue
			}

			var inputConfig, expectedConfig Config
			if err := Unmarshal(content, &inputConfig); err != nil {
				t.Errorf("Failed to unmarshal input from %s: %v", filename, err)
				continue
			}
			if err := json.Unmarshal(expectedContent, &expectedConfig); err != nil {
				t.Errorf("Failed to unmarshal expected content for %s: %v", testName, err)
				continue
			}

			if inputConfig != expectedConfig {
				t.Errorf("Mismatch in %s. Expected: %+v, Got: %+v", testName, expectedConfig, inputConfig)
			}
		}
	}
}
