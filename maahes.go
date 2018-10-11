package main

import (
	"fmt"
	"os"

	"github.com/bramz/maahes/lib"
)

func main() {
	if len(os.Args) > 1 {
		lib.StartSession(os.Args[1])
	} else {
		fmt.Print("you must provide a token\n")
		os.Exit(0)
	}
}
