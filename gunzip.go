package main

import (
	"compress/gzip"
	"io"
	"os"
)

func gunzipFile(src, dest string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	reader, err := gzip.NewReader(in)
	if err != nil {
		return err
	}
	defer reader.Close()

	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.CopyBuffer(out, reader, make([]byte, 32*1024))
	return err
}
