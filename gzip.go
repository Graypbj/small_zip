package main

import (
	"compress/gzip"
	"io"
	"os"
)

func gzipFile(src, dest string, buf []byte) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := gzip.NewWriter(out)
	defer writer.Close()

	_, err = io.CopyBuffer(writer, in, buf)
	return err
}
