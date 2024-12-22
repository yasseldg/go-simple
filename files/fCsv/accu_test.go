package fCsv

import (
	"os"
	"testing"
)

func TestAccu(t *testing.T) {
	filePath := "test_accu.csv"
	defer os.Remove(filePath)

	accu, err := NewAccu(filePath, true, 2)
	if err != nil {
		t.Fatalf("Failed to create Accu: %s", err)
	}

	header := []string{"Column1", "Column2", "Column3"}
	accu.AddHeader(header)

	data1 := []string{"Data1", "Data2", "Data3"}
	accu.AddData(data1)

	data2 := []string{"Data4", "Data5", "Data6"}
	accu.AddData(data2)

	// Check if data is saved correctly
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		t.Fatalf("Failed to get file stats: %s", err)
	}

	if stat.Size() == 0 {
		t.Error("Expected data to be written to file, but file is empty")
	}
}
