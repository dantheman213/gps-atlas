package serial

import (
    "errors"
    "fmt"
    "github.com/dantheman213/gps-atlas/pkg/utility"
    libSerial "github.com/tarm/serial"
    "strings"
    "time"
)

type GPSDevice struct {
    PortNumber int
    PortName string
    BaudRate int
    Port *libSerial.Port
}

var buf = make([]byte, 1024) // re-usable / better performance

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
    str := ""
    for true {
        n, err := port.Read(buf)
        if err != nil {
            return "", err
        }
        str += string(buf[:n])

        last := str[len(str)-1:]
        if strings.Index(last, "\r") > -1 || strings.Index(last, "\n") > -1 {
            return fmt.Sprintf("%s\n", strings.TrimSpace(str)), nil
        }
    }

    return "", errors.New("couldn't read serial data")
}
