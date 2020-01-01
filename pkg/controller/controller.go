package controller

import (
    "bufio"
    "fmt"
    "github.com/dantheman213/gps-atlas/pkg/gps"
    "github.com/dantheman213/gps-atlas/pkg/serial"
    "log"
    "os"
)

var writeBufferCSV *bufio.Writer = nil
var writeBufferGPS *bufio.Writer = nil
var writeBufferKML *bufio.Writer = nil
var writeBufferNMEA *bufio.Writer = nil
var fileHandles []*os.File // TODO: Close all file handles before app exits

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
    coordsPrettyStr := gps.GenerateGPSCoordsPretty(loc) // will only be populated if received message was GGA

    if *opts.PrintNMEAToCLI {
        fmt.Print(nmea)
    }

    if *opts.WriteNMEAFilePath != "" {
        if writeBufferNMEA == nil {
            log.Printf("[info] opening and writing NMEA data to file: %s", *opts.WriteNMEAFilePath)
            f, err := os.Create(*opts.WriteNMEAFilePath)
            if err != nil {
                log.Fatalf("[error] %s", err)
            }

            fileHandles = append(fileHandles, f)
            writeBufferNMEA = bufio.NewWriter(f)
        }

        if err := writeData(*writeBufferNMEA, nmea); err != nil {
            log.Fatalf("[error] %s", err)
        }
    }

    if *opts.PrintGPSCoordsToCLI && coordsPrettyStr != "" {
        fmt.Print(coordsPrettyStr)
    }

    if *opts.PrintGPSExtraInfoToCLI {
        // TODO
    }

    if *opts.WriteGPSCoordsFilePath != "" && coordsPrettyStr != "" {
        if writeBufferGPS == nil {
            log.Printf("[info] opening and writing GPS coordinate data to file: %s", *opts.WriteGPSCoordsFilePath)
            f, err := os.Create(*opts.WriteGPSCoordsFilePath)
            if err != nil {
                log.Fatalf("[error] %s", err)
            }

            fileHandles = append(fileHandles, f)
            writeBufferGPS = bufio.NewWriter(f)
        }

        if err := writeData(*writeBufferGPS, coordsPrettyStr); err != nil {
            log.Fatalf("[error] %s", err)
        }
    }

    if *opts.WriteCSVFilePath != "" {
        // TODO
    }

    if *opts.WriteKMLFilePath != "" {
        // TODO
    }
}
