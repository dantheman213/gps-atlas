package controller

import (
    "bufio"
    libGPS "github.com/dantheman213/gps-atlas/pkg/gps"
    "github.com/dantheman213/gps-atlas/pkg/serial"
    "log"
    "os"
)

var gps *libGPS.GPS
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

    gps = libGPS.NewGPS()
    for true {
        processData(d)
    }
}

func processData(device *serial.GPSDevice) {
    d, err := serial.ReadSerialData(device.Port)
    if err != nil {
        log.Printf("couldn't read data stream")
        return
    }

    gps.IngestNMEASentences(d)
}
