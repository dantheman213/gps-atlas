package main

func main() {
	o := parseFlags()
	if *o.Help == true {
		printHelpSheet()
	}
}
