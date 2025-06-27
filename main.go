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
	flag.Parse()

	switch *mode {
	case "zip":
		err := zipFile(*src, *dest)
		checkErr(err)
	case "unzip":
		err := unzipFile(*src, *dest)
		checkErr(err)
	case "gzip":
		err := gzipFile(*src, *dest)
		checkErr(err)
	case "gunzip":
		err := gunzipFile(*src, *dest)
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
