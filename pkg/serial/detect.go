package serial

import (
    "errors"
    "fmt"
    "log"
)

func DetectGPSDevice() (*GPSDevice, error) {
    log.Println("[info] search and auto-detect GPS devices in progress...")

    currentPortNumber := 1
    currentPortName := fmt.Sprintf("%s%d", "COM", currentPortNumber)
    currentBaudRate := 9600

    log.Printf("[info] attemping to contact port %s with baud rate %d", currentPortName, currentBaudRate)
    d, err := Connect(currentPortName, currentBaudRate, 5)

    if err != nil {
        log.Println("[info] couldn't establish a connection")
    } else {
        // Found a valid serial device to connect, now check if it's an NMEA compliant GPS serial device
        s, err := ReadSerialData(d.Port)

        if err != nil {
            fmt.Println(err)
        } else {
            // Got some data, is it in NMEA format?
            if s == "" {
                // TODO
            }
        }
    }

    return nil, errors.New("could not find a valid GPS serial device")
}
