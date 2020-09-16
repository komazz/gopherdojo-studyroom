package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai4/komazz/server"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
