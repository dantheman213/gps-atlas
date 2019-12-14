package main

import (
    "fmt"
    "github.com/dantheman213/gps-test/pkg/gps"
    "github.com/tarm/serial"
    "log"
    "strings"
)

func main() {
    c := &serial.Config{
        Name: "COM4",
        Baud: 9600,
        ReadTimeout: 1,
        Size: 8,
    }

    s, err := serial.OpenPort(c)
    if err != nil {
        log.Fatal(err)
    }
    defer s.Close()

    buf := make([]byte, 1024)

    for {
        n, err := s.Read(buf)
        if err != nil {
            log.Fatal(err)
        }
        s := string(buf[:n])
        if strings.Index(s, "GPGGA") > -1 {
            //fmt.Print(s)
            gpsPayload, err := gps.ParseGPGGA(s)
            if err != nil {
                // TODO
            } else {
                lat, err := gpsPayload.GetLatitude()
                if err != nil {
                    // TODO
                }
                long, err := gpsPayload.GetLongitude()
                if err != nil {
                    // TODO
                }

                fmt.Println(fmt.Sprintf("%f,%f", lat, long))
            }
        }
    }
}