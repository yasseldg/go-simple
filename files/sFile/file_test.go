package sFile

import (
	"os"
	"testing"
)

func TestDeletePath(t *testing.T) {
	filePath := "testfile.txt"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	file.Close()

	err = DeletePath(filePath)
	if err != nil {
		t.Errorf("Failed to delete file: %s", err)
	}

	exists, err := ExistingPath(filePath)
	if err != nil {
		t.Errorf("Failed to check if file exists: %s", err)
	}
	if exists {
		t.Errorf("File should not exist")
	}
}

func TestExistingPath(t *testing.T) {
	filePath := "testfile.txt"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	file.Close()
	defer os.Remove(filePath)

	exists, err := ExistingPath(filePath)
	if err != nil {
		t.Errorf("Failed to check if file exists: %s", err)
	}
	if !exists {
		t.Errorf("File should exist")
	}

	nonExistentPath := "nonexistentfile.txt"
	exists, err = ExistingPath(nonExistentPath)
	if err != nil {
		t.Errorf("Failed to check if file exists: %s", err)
	}
	if exists {
		t.Errorf("File should not exist")
	}
}

func TestGetDirsNames(t *testing.T) {
	dirPath := "testdir"
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %s", err)
	}
	defer os.RemoveAll(dirPath)

	subDirPath := dirPath + "/subdir"
	err = os.Mkdir(subDirPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %s", err)
	}

	dirs, err := GetDirsNames(dirPath)
	if err != nil {
		t.Errorf("Failed to get directory names: %s", err)
	}
	if len(dirs) != 1 || dirs[0] != "subdir" {
		t.Errorf("Expected subdir, got %v", dirs)
	}
}

func TestGetFilesNames(t *testing.T) {
	dirPath := "testdir"
	err := os.Mkdir(dirPath, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory: %s", err)
	}
	defer os.RemoveAll(dirPath)

	filePath := dirPath + "/testfile.txt"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	file.Close()

	files, err := GetFilesNames(dirPath)
	if err != nil {
		t.Errorf("Failed to get file names: %s", err)
	}
	if len(files) != 1 || files[0] != "testfile.txt" {
		t.Errorf("Expected testfile.txt, got %v", files)
	}
}

func TestGetDir(t *testing.T) {
	dirPath := "testdir"
	err := GetDir(dirPath)
	if err != nil {
		t.Errorf("Failed to get directory: %s", err)
	}
	defer os.RemoveAll(dirPath)

	exists, err := ExistingPath(dirPath)
	if err != nil {
		t.Errorf("Failed to check if directory exists: %s", err)
	}
	if !exists {
		t.Errorf("Directory should exist")
	}
}
