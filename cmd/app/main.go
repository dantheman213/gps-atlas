package main

import "os"

func main() {
	o := parseFlags()
	if len(os.Args) == 1 || *o.Help == true {
		printHelpSheet()
	}
}
