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
        processData(d)
    }
}

func processData(device *serial.GPSDevice) {
    sentenceNMEA, err := serial.ReadSerialData(device.Port)
    if err != nil {
        log.Printf("couldn't read data stream")
        return
    }

    interpretNMEA(sentenceNMEA)
}

func interpretNMEA(sentence string) {
    var loc *gps.Location = nil
    sentenceGGA, err := gps.CheckForLocationInfo(sentence)
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

    shareData(sentence, loc)
}

func shareData(nmea string, loc *gps.Location) {
    if *opts.PrintNMEAToCLI {
        fmt.Print(nmea)
    }

    if *opts.WriteNMEAFilePath != "" {
        // TODO
    }

    if *opts.PrintGPSCoordsToCLI && loc != nil {
        fmt.Printf("%f, %f\n", loc.Latitude, loc.Longitude)
    }

    if *opts.PrintGPSExtraInfoToCLI {
        // TODO
    }

    if *opts.WriteGPSCoordsFilePath != "" {
        // TODO
    }

    if *opts.WriteCSVFilePath != "" {
        // TODO
    }

    if *opts.WriteKMLFilePath != "" {
        // TODO
    }
}
