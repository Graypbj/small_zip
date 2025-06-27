package main

import (
	"compress/gzip"
	"io"
	"os"
)

func gzipFile(src, dest string) error {
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

	_, err = io.CopyBuffer(writer, in, make([]byte, 32*1024))
	return err
}
