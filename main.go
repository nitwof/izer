package main

import (
	"bufio"
	"fmt"
	"os"

	"iconizer/icons"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:len(os.Args)] {
			fmt.Printf("%s %s\n", icons.Get(arg), arg)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Printf("%s %s\n", icons.Get(scanner.Text()), scanner.Text())
		}

		if scanner.Err() != nil {
			os.Exit(1)
		}
	}
}
