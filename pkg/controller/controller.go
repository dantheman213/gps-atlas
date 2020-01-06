package controller

import (
    "bufio"
    "fmt"
    libGPS "github.com/dantheman213/gps-atlas/pkg/gps"
    "github.com/dantheman213/gps-atlas/pkg/serial"
    "log"
    "os"
)

var device *serial.GPSDevice = nil
var gps *libGPS.GPS = nil

var writeBufferCSV *bufio.Writer = nil
var writeBufferGPS *bufio.Writer = nil
var writeBufferKML *bufio.Writer = nil
var writeBufferNMEA *bufio.Writer = nil

var fileHandles []*os.File // TODO: Close all file handles before app exits

func Start() {
    if *opts.AutoDetect {
        var err error
        device, err = serial.DetectGPSDevice()
        if err != nil {
            log.Fatalf("[error] %s", err)
        }
    }

    gps = libGPS.NewGPS()
    for true {
        processData()
    }
}

func processData() {
    d, err := serial.ReadSerialData(device.Port)
    if err != nil {
        log.Printf("couldn't read data stream")
        return
    }

    gps.IngestNMEASentences(d)
    displayData(d)
    writeData(d)
}

func displayData(sentences string) {
    if *opts.PrintNMEAToCLI {
        fmt.Print(sentences)
    }

    c := gps.GetGPSLocationPretty()
    if *opts.PrintGPSCoordsToCLI && c != "" {
        fmt.Print(c)
    }

    if *opts.PrintGPSExtraInfoToCLI {
        // TODO
    }
}

func writeData(sentences string) {
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

        if err := write(*writeBufferNMEA, sentences); err != nil {
            log.Fatalf("[error] %s", err)
        }
    }

    c := gps.GetGPSLocationPretty()
    if *opts.WriteGPSCoordsFilePath != "" && c != "" {
        if writeBufferGPS == nil {
            log.Printf("[info] opening and writing GPS coordinate data to file: %s", *opts.WriteGPSCoordsFilePath)
            f, err := os.Create(*opts.WriteGPSCoordsFilePath)
            if err != nil {
                log.Fatalf("[error] %s", err)
            }

            fileHandles = append(fileHandles, f)
            writeBufferGPS = bufio.NewWriter(f)
        }

        if err := write(*writeBufferGPS, c); err != nil {
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
