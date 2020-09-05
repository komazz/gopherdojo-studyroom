package main

import (
	"fmt"
	"os"

	"github.com/gopherdojo/gopherdojo-studyroom/kadai3-1/komazz/typegame"
)

func main() {
	if err := typegame.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
