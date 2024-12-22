package fCsv

import (
	"os"
	"testing"
)

func TestIter(t *testing.T) {
	filePath := "test_iter.csv"
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

	// Test Iter
	iter, err := newIter(filePath, 2, ',')
	if err != nil {
		t.Fatalf("Failed to create Iter: %s", err)
	}

	if !iter.open() {
		t.Fatalf("Failed to open Iter: %s", iter.Error())
	}

	// Read and check data
	for iter.Next() {
		item := iter.Item
		if len(item) != 3 {
			t.Errorf("Expected 3 columns, got %d", len(item))
		}
	}

	if iter.Error() != nil {
		t.Errorf("Iter encountered an error: %s", iter.Error())
	}
}
