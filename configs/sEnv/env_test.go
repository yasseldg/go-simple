package sEnv

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	os.Setenv("TEST_ENV", "test_value")
	defer os.Unsetenv("TEST_ENV")

	if value := Get("TEST_ENV", "default_value"); value != "test_value" {
		t.Errorf("Expected 'test_value', got '%s'", value)
	}

	if value := Get("NON_EXISTENT_ENV", "default_value"); value != "default_value" {
		t.Errorf("Expected 'default_value', got '%s'", value)
	}
}

func TestGetSlice(t *testing.T) {
	os.Setenv("TEST_ENV_SLICE", "value1,value2,value3")
	defer os.Unsetenv("TEST_ENV_SLICE")

	expected := []string{"value1", "value2", "value3"}
	if value := GetSlice("TEST_ENV_SLICE"); !equalSlices(value, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, value)
	}

	defaultExpected := []string{"default1", "default2"}
	if value := GetSlice("NON_EXISTENT_ENV_SLICE", "default1", "default2"); !equalSlices(value, defaultExpected) {
		t.Errorf("Expected '%v', got '%v'", defaultExpected, value)
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("TEST_ENV_INT", "42")
	defer os.Unsetenv("TEST_ENV_INT")

	if value := GetInt("TEST_ENV_INT", 0); value != 42 {
		t.Errorf("Expected 42, got %d", value)
	}

	if value := GetInt("NON_EXISTENT_ENV_INT", 100); value != 100 {
		t.Errorf("Expected 100, got %d", value)
	}
}

func TestGetInt64(t *testing.T) {
	os.Setenv("TEST_ENV_INT64", "64")
	defer os.Unsetenv("TEST_ENV_INT64")

	if value := GetInt64("TEST_ENV_INT64", 0); value != 64 {
		t.Errorf("Expected 64, got %d", value)
	}

	if value := GetInt64("NON_EXISTENT_ENV_INT64", 100); value != 100 {
		t.Errorf("Expected 100, got %d", value)
	}
}

func TestGetBool(t *testing.T) {
	os.Setenv("TEST_ENV_BOOL", "true")
	defer os.Unsetenv("TEST_ENV_BOOL")

	if value := GetBool("TEST_ENV_BOOL", false); value != true {
		t.Errorf("Expected true, got %v", value)
	}

	if value := GetBool("NON_EXISTENT_ENV_BOOL", true); value != true {
		t.Errorf("Expected true, got %v", value)
	}
}

func TestGetSliceInt(t *testing.T) {
	os.Setenv("TEST_ENV_SLICE_INT", "1,2,3")
	defer os.Unsetenv("TEST_ENV_SLICE_INT")

	expected := []int{1, 2, 3}
	if value := GetSliceInt("TEST_ENV_SLICE_INT"); !equalIntSlices(value, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, value)
	}

	defaultExpected := []int{10, 20}
	if value := GetSliceInt("NON_EXISTENT_ENV_SLICE_INT", 10, 20); !equalIntSlices(value, defaultExpected) {
		t.Errorf("Expected '%v', got '%v'", defaultExpected, value)
	}
}

func TestGetFloat64(t *testing.T) {
	os.Setenv("TEST_ENV_FLOAT64", "3.14")
	defer os.Unsetenv("TEST_ENV_FLOAT64")

	if value := GetFloat64("TEST_ENV_FLOAT64", 0.0); value != 3.14 {
		t.Errorf("Expected 3.14, got %f", value)
	}

	if value := GetFloat64("NON_EXISTENT_ENV_FLOAT64", 1.23); value != 1.23 {
		t.Errorf("Expected 1.23, got %f", value)
	}
}

func TestLoadYaml(t *testing.T) {
	// Create a temporary YAML file for testing
	file, err := os.CreateTemp("", "test.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	content := `
key1: value1
key2: value2
`
	if _, err := file.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	file.Close()

	var result map[string]string
	if err := LoadYaml(file.Name(), &result); err != nil {
		t.Fatalf("Failed to load YAML: %v", err)
	}

	expected := map[string]string{"key1": "value1", "key2": "value2"}
	if !equalMaps(result, expected) {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}

// Helper functions for comparing slices and maps
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func equalMaps(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}
