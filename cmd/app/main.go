package main

func main() {
	c := parseFlags()
	if *c.Help == true {
		printHelpSheet()
	}
}
