package gps

import (
    "fmt"
    "github.com/dantheman213/gps-atlas/pkg/gps/nmea"
    "strings"
)

// Decimal Degrees
type LocationDD struct {
    Latitude float32
    Longitude float32
}

// Degrees Minutes Seconds
type LocationDMS struct {
    Latitude float32
    LatitudeDirection string
    Longitude float32
    LongitudeDirection string
}

func CheckForLocationInfo(nmeaSentence string) (*nmea.GGA, error) {
    if strings.Index(nmeaSentence, "GGA") > -1 {
        message, err := nmea.ParseGGA(nmeaSentence)
        if err != nil {
            return nil, err
        } else {
            return message, nil
        }
    }

    return nil, nil
}

func GetGPSLocation(ggaMessage nmea.GGA) (*LocationDD, error) {
    lat, err := ggaMessage.GetLatitudeDD()
    if err != nil {
        return nil, err
    }

    long, err := ggaMessage.GetLongitudeDD()
    if err != nil {
        return nil, err
    }

    return &LocationDD{
        Latitude:  lat,
        Longitude: long,
    }, nil
}

func GetGPSLocationPretty(loc *LocationDD) string {
    str := ""
    if loc != nil {
        str = fmt.Sprintf("%f, %f\n", loc.Latitude, loc.Longitude)
    }

    return str
}

func IngestNMEASentences(rawStr string) error {
    return nil
}
