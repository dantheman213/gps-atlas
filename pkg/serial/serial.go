package serial

import (
    libSerial "github.com/tarm/serial"
    "time"
)

func Connect(portName string, baudRate int, timeout int) (*libSerial.Port, error) {
    c := &libSerial.Config{
        Name: portName,
        Baud: baudRate,
        ReadTimeout: time.Duration(timeout),
        Size: 8,
    }

    return libSerial.OpenPort(c)
}
