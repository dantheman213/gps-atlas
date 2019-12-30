package controller

import (
    "github.com/dantheman213/gps-usb-serial-reader/pkg/serial"
    "log"
)

func Start() {
    if *opts.AutoDetect {
        _, err := serial.DetectGPSDevice()
        if err != nil {
            log.Fatalf("[error] %s", err)
        }
    }
}