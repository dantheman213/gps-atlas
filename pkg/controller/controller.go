package controller

import (
    "bufio"
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
    _, err := serial.ReadSerialData(device.Port)
    if err != nil {
        log.Printf("couldn't read data stream")
        return
    }
}
