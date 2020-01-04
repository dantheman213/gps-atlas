package serial

import (
    "errors"
    "fmt"
    "log"
    "runtime"
    "strings"
)

const (
    autoDetectPortCount = 50
    dataDetectIterationCount = 100
    portDetectTimeout = 5
)

var baudRates = [...]int { 115200, 38400, 19200, 9600, 4800 }

func DetectGPSDevice() (*GPSDevice, error) {
    log.Println("[info] search and auto-detect GPS devices in progress...")

    portPrefix := ""
    switch runtime.GOOS {
    case "linux":
    case "darwin":
        portPrefix = "/dev/ttyS"
        break
    case "windows":
        portPrefix = "COM"
        break
    default:
        portPrefix = ""
    }

    for currentPortNumber := 0; currentPortNumber < autoDetectPortCount; currentPortNumber++ {
        currentPortName := fmt.Sprintf("%s%d", portPrefix, currentPortNumber)
        for bIndex := 0; bIndex < len(baudRates); bIndex++ {
            // TODO deal with MacOS serial ports with baudrate as suffix
            currentBaudRate := baudRates[bIndex]
            log.Printf("[info] attemping to contact port %s with baud rate %d", currentPortName, currentBaudRate)
            d, err := Connect(currentPortName, baudRates[bIndex], portDetectTimeout)

            if err != nil {
                log.Println("[info] couldn't establish a connection")
                continue
            } else {
                // Found a valid serial device to connect, now check if it's an NMEA compliant GPS serial device
                // Check in multiple iterations to ensure not correct protocol
                for commCount := 0; commCount < dataDetectIterationCount; commCount++ {
                    str, err := ReadSerialData(d.Port)
                    if err != nil {
                        log.Printf("[info] %s", err)
                    } else {
                        if len(str) > 7 {
                            // Got some data, is it in NMEA format?
                            if str[0:2] == "$G" {
                                firstCommaPos := strings.Index(str, ",")
                                if firstCommaPos == 6 || firstCommaPos == 7 || firstCommaPos == 9 {
                                    // ex: $GPGGA, $GPPTNL, $GPPFUGDP
                                    log.Printf("[info] found GPS device at port: %s baud rate: %d", d.PortName, d.BaudRate)
                                    return d, nil
                                }
                            }
                        }
                    }
                }

                log.Println("[info] found data but doesn't appear to be NMEA... skipping...")
            }

            log.Printf("[info] closing port %s", currentPortName)
            _ = d.Port.Close()
        }
    }

    return nil, errors.New("could not find a valid GPS serial device")
}
