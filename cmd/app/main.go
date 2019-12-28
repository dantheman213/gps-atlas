package main

import (
	"fmt"
	"github.com/dantheman213/gps-usb-serial-reader/pkg/controller"
	"os"
)

func main() {
	controller.ParseOptions()
	o := controller.GetOptions()
	if *o.Help == true {
		controller.PrintHelpSheet()
		os.Exit(0)
	}

	if len(os.Args) == 1 {
		fmt.Println("Info: use '--help' flag for additional options that can be used")
	}

	controller.SanitizeOptions(o)
	if err := controller.ValidateOptions(o); err != nil {
		fmt.Printf("Parse exception: %s", err)
		os.Exit(1)
	}

	controller.Start()
}
