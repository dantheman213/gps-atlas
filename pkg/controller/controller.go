package controller

import (
    "fmt"
    "github.com/dantheman213/gps-atlas/pkg/gps"
    "github.com/dantheman213/gps-atlas/pkg/serial"
    "log"
)

func Start() {
    var d *serial.GPSDevice = nil

    if *opts.AutoDetect {
        var err error
        d, err = serial.DetectGPSDevice()
        if err != nil {
            log.Fatalf("[error] %s", err)
        }
    }

    for true {
        Process(d)
    }
}

func Process(device *serial.GPSDevice) {
    sentenceNMEA, err := serial.ReadSerialData(device.Port)
    if err != nil {
        log.Printf("couldn't read data stream")
        return
    }

    var loc *gps.Location = nil
    sentenceGGA, err := gps.CheckForLocationInfo(sentenceNMEA)
    if err != nil {
        log.Printf("GGA sentence malformed")
        return
    }
    if sentenceGGA != nil {
        loc, err = sentenceGGA.GetLocationPoint()
        if err != nil {
            // log.Printf("couldn't calculate coordinates from GGA sentence")
            return
        }
    }

    if *opts.PrintNMEAToCLI {
        fmt.Print(sentenceNMEA)
    }

    if *opts.PrintGPSCoordsToCLI && loc != nil {
        fmt.Printf("%f, %f\n", loc.Latitude, loc.Longitude)
    }
}