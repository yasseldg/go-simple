package sFile

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// UnGzip descomprime un archivo .gz y lo guarda en el directorio especificado.
func UnGzip(file_path, dir_path string) error {
	// Abrir el archivo .gz
	file, err := os.Open(file_path)
	if err != nil {
		return fmt.Errorf("error al abrir el archivo gzip: %v", err)
	}
	defer file.Close()

	// Crear un lector gzip
	gzReader, err := gzip.NewReader(file)
	if err != nil {
		return fmt.Errorf("error al crear el lector gzip: %v", err)
	}
	defer gzReader.Close()

	output_file_name := filepath.Base(file_path[:len(file_path)-len(filepath.Ext(file_path))])

	// Construir la ruta de salida
	output_path := filepath.Join(dir_path, output_file_name)

	// Crear el archivo de salida
	outFile, err := os.Create(output_path)
	if err != nil {
		return fmt.Errorf("error al crear el archivo de salida: %v", err)
	}
	defer outFile.Close()

	// Copiar el contenido descomprimido al archivo de salida
	_, err = io.Copy(outFile, gzReader)
	if err != nil {
		return fmt.Errorf("error al escribir el contenido descomprimido: %v", err)
	}

	return nil
}
