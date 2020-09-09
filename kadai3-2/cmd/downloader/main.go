package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai3-2/zr/downloader"
)

func main() {
	if err := downloader.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
