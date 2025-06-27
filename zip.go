package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func zipFile(src, dest string, buf []byte) error {
	outFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		checkErr(err)

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		inFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer inFile.Close()

		fh, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		fh.Name = relPath
		fh.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(fh)
		if err != nil {
			return err
		}

		_, err = io.CopyBuffer(writer, inFile, buf)
		return err
	})
}
