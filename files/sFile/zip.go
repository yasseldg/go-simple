package sFile

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// UnZip descomprime un archivo .zip en el directorio especificado.
func UnZip(file_path, dir_path string) error {
	// Abrir el archivo .zip
	r, err := zip.OpenReader(file_path)
	if err != nil {
		return fmt.Errorf("error al abrir archivo zip: %v", err)
	}
	defer r.Close()

	// Iterar sobre los archivos del zip
	for _, file := range r.File {
		// Construir la ruta del archivo
		file_path := filepath.Join(dir_path, file.Name)

		// Comprobar que no intente escribir fuera del directorio de destino
		if !strings.HasPrefix(file_path, filepath.Clean(dir_path)+string(os.PathSeparator)) {
			return fmt.Errorf("archivo %s intenta escribir fuera del directorio de destino", file_path)
		}

		// Crear directorios necesarios
		if file.FileInfo().IsDir() {
			os.MkdirAll(file_path, os.ModePerm)
			continue
		}

		// Crear el archivo si no es un directorio
		if err := os.MkdirAll(filepath.Dir(file_path), os.ModePerm); err != nil {
			return fmt.Errorf("error al crear directorios: %v", err)
		}

		// Abrir el archivo dentro del zip
		destFile, err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("error al crear archivo descomprimido: %v", err)
		}
		defer destFile.Close()

		// Leer el contenido del archivo comprimido
		zipFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("error al abrir archivo en zip: %v", err)
		}
		defer zipFile.Close()

		// Copiar el contenido al nuevo archivo
		_, err = io.Copy(destFile, zipFile)
		if err != nil {
			return fmt.Errorf("error al copiar contenido al archivo descomprimido: %v", err)
		}
	}

	return nil
}
