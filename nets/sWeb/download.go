package sWeb

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/yasseldg/go-simple/files/sFile"
)

func DownloadFile(dir_path, file_name, url string) error {

	err := sFile.GetDir(dir_path)
	if err != nil {
		return fmt.Errorf("GetDir( %s ): %s ", dir_path, err)
	}

	file_path := fmt.Sprintf("%s/%s", dir_path, file_name)

	exist, err := sFile.ExistingPath(file_path)
	if exist {
		return fmt.Errorf("file ( %s ) exist ", file_path)
	}

	if err != nil {
		return fmt.Errorf("ExistingPath( %s ): %s ", file_path, err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("http.Get( %s ): %s ", url, err)
	}
	defer resp.Body.Close()

	out, err := os.Create(file_path)
	if err != nil {
		return fmt.Errorf("os.Create( %s ): %s ", file_path, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("io.Copy( %s ): %s ", file_path, err)
	}

	return nil
}
