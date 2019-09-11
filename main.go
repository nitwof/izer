package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:len(os.Args)] {
			fmt.Printf("$ %s\n", arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Printf("$ %s\n", scanner.Text())
		}

		if scanner.Err() != nil {
			os.Exit(1)
		}
	}
}
