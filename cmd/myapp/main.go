package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	branchName string
	commitId   string
	buildTime  string

	showVer = flag.Bool("v", false, "show version")
)

func main() {
	if *showVer {
		fmt.Printf("%s: %s\t%s\n", branchName, commitId, buildTime)
		os.Exit(0)
	}
}