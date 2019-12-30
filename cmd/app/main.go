package main

import (
	"github.com/dantheman213/gps-usb-serial-reader/pkg/controller"
	"log"
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
		log.Println("[info] use '--help' flag for additional options that can be used")
	}

	controller.SanitizeOptions(o)
	if err := controller.ValidateOptions(o); err != nil {
		log.Fatalf("[error] %s", err)
	}

	controller.Start()
}
