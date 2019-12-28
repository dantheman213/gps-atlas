package serial

import "fmt"

type GPSDevice struct {
    RawPort int
    FullPort string
    BaudRate int
}

func DetectGPSDevice() (*GPSDevice, error) {
    fmt.Println("Info: auto detect starting...")
    return nil, nil
}