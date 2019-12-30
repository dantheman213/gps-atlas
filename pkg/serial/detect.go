package serial

import (
    "log"
)

type GPSDevice struct {
    PortNumber int
    PortName string
    BaudRate int
}

func DetectGPSDevice() (*GPSDevice, error) {
    log.Print("[info] auto detect starting...")



    return nil, nil
}