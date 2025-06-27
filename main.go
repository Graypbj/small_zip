package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	mode := flag.String("mode", "", "zip | unzip | gzip | gunzip")
	src := flag.String("src", "", "Source file or directory")
	dest := flag.String("dest", "", "Destination path")
	bufSize := flag.Int("bufsize", 32*1024, "Maximum buffer size in bytes: (default: 32KB)")
	flag.Parse()

	if *bufSize <= 0 {
		fmt.Fprintln(os.Stderr, "Error: buffer size must be > 0")
		os.Exit(1)
	}
	buffer := make([]byte, *bufSize)

	switch *mode {
	case "zip":
		err := zipFile(*src, *dest, buffer)
		checkErr(err)
	case "unzip":
		err := unzipFile(*src, *dest, buffer)
		checkErr(err)
	case "gzip":
		err := gzipFile(*src, *dest, buffer)
		checkErr(err)
	case "gunzip":
		err := gunzipFile(*src, *dest, buffer)
		checkErr(err)
	default:
		fmt.Println("Invalid mode. Use zip, unzip, gzip, or gunzip")
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
