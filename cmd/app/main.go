package main

import (
    "fmt"
    "github.com/dantheman213/gps-atlas/pkg/controller"
    "log"
    "os"
)

var Version string = "v0.0.0" // populated at build time

func main() {
    controller.ParseOptions()
    o := controller.GetOptions()

    if *o.Help == true {
        controller.PrintHelpSheet(Version)
        os.Exit(0)
    }

    if *o.Version == true {
        fmt.Println(Version)
        os.Exit(0)
    }

    if len(os.Args) == 1 {
        log.Println("[info] use '--help' flag for additional options that can be used")
    }

    controller.SanitizeOptions()
    if err := controller.ValidateOptions(); err != nil {
        log.Fatalf("[error] %s", err)
    }

    controller.Start()
}
