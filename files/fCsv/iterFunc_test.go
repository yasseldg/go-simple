package fCsv

import (
	"os"
	"testing"
)

func TestIterFunc(t *testing.T) {
	filePath := "test_iter_func.csv"
	defer os.Remove(filePath)

	// Create a CSV file with test data
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"Column1", "Column2", "Column3"}
	writer.Write(header)

	data1 := []string{"Data1", "Data2", "Data3"}
	writer.Write(data1)

	data2 := []string{"Data4", "Data5", "Data6"}
	writer.Write(data2)

	writer.Flush()

	// Test IterFunc
	iterFunc, err := NewIterFunc(filePath, ',', func(line []string) error {
		if len(line) != 3 {
			t.Errorf("Expected 3 columns, got %d", len(line))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("Failed to create IterFunc: %s", err)
	}

	iterFunc.Run()

	if iterFunc.Error() != nil {
		t.Errorf("IterFunc encountered an error: %s", iterFunc.Error())
	}
}
