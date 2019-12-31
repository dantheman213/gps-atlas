package serial

import (
    "github.com/dantheman213/gps-atlas/pkg/utility"
    libSerial "github.com/tarm/serial"
    "time"
)

type GPSDevice struct {
    PortNumber int
    PortName string
    BaudRate int
    Port *libSerial.Port
}

var buf = make([]byte, 115200) // re-usable / better performance

func Connect(portName string, baudRate int, timeout int) (*GPSDevice, error) {
    c := &libSerial.Config{
        Name: portName,
        Baud: baudRate,
        ReadTimeout: time.Duration(timeout) * time.Second,
        Size: 8,
    }

    p, err := libSerial.OpenPort(c)
    if err != nil {
        return nil, err
    }

    return &GPSDevice{
        Port: p,
        PortNumber: utility.ExtractIntFromStr(portName),
        PortName:   portName,
        BaudRate:   baudRate,
    }, nil
}

func ReadSerialData(port *libSerial.Port) (string, error) {
    n, err := port.Read(buf)
    if err != nil {
        return "", err
    }

    return string(buf[:n]), nil
}