package controller

import "github.com/dantheman213/gps-usb-serial-reader/pkg/serial"

func Start() {
    if *opts.AutoDetect {
        _, err := serial.DetectGPSDevice()
        if err != nil {
            // TODO
        }
    }
}