package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	first := os.Args[1]
	if first == "help" {
		printHelp()
		return
	} else if first == "inspect" {
		if len(os.Args) < 3 {
			printHelp()
			return
		}

		err := inspect(os.Args[2])
		if err != nil {
			fmt.Println(err)
		}

		return
	}
}
