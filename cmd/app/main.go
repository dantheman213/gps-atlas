package main

import "os"

func main() {
	o := parseOptions()
	if len(os.Args) == 1 || *o.Help == true {
		printHelpSheet()
	}
}
