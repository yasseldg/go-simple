package sFile

import (
	"fmt"
	"os"
)

func DeletePath(file_path string) error {
	exist, err := ExistingPath(file_path)
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}

	err = os.Remove(file_path)
	if err != nil {
		err = fmt.Errorf("os.Remove( %s ): %s", file_path, err.Error())
	}

	return err
}

func ExistingPath(file_path string) (bool, error) {
	_, err := os.Stat(file_path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetDirsNames(dir_path string) ([]string, error) {
	return getEntriesNames(dir_path, true)
}

func GetFilesNames(dir_path string) ([]string, error) {
	return getEntriesNames(dir_path, false)
}

func GetDir(dir_path string) (err error) {
	_, err = os.Stat(dir_path)
	if err == nil {
		// Directory exists
		return nil
	}

	if os.IsNotExist(err) {
		// File or directory does not exist
		return os.MkdirAll(dir_path, mode(0755, os.ModeDir))
	}

	return fmt.Errorf("os.Stat( %q ): %s", dir_path, err.Error())
}

// private functions

// mode returns the file mode masked by the umask
func mode(mode, umask os.FileMode) os.FileMode {
	return mode & ^umask
}

func getEntriesNames(dir_path string, is_dir bool) ([]string, error) {
	exists, err := ExistingPath(dir_path)
	if err != nil {
		return nil, fmt.Errorf("ExistingPath(): %s", err.Error())
	}
	if !exists {
		return nil, fmt.Errorf("%s not exist", dir_path)
	}

	entries, err := os.ReadDir(dir_path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadDir(): %s", err.Error())
	}

	var names []string
	for _, entry := range entries {
		if is_dir == entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	return names, nil
}
