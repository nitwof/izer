package main

import (
	"fmt"

	"github.com/NightWolf007/izer/cmd"
)

var (
	version = "master"
	commit  = "none"
)

func main() {
	cmd.Execute(fmt.Sprintf("%s (%s)", version, commit))
}
