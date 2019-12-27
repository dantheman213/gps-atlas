package main

import (
	"fmt"
	"os"
)

func main() {
	o := parseOptions()
	if len(os.Args) == 1 || *o.Help == true {
		printHelpSheet()
		os.Exit(0)
	}

	if err := checkOptionSanity(o); err != nil {
		fmt.Printf("Parse exception: %s", err)
	}
}
