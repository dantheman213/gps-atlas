package gps

import (
    "errors"
    "fmt"
    "strings"
)

type Location struct {
    Latitude float32
    Longitude float32
}

func CheckForLocationInfo(sentenceNMEA string) (*GGA, error) {
    if strings.Index(sentenceNMEA, "GGA") > -1 {
        g, err := ParseGPGGA(sentenceNMEA)
        if err != nil {
            return nil, err
        } else {
            return g, nil
        }
    }

    return nil, nil
}

func GenerateGPSCoordsPretty(loc *Location) string {
    str := ""
    if loc != nil {
        str = fmt.Sprintf("%f, %f\n", loc.Latitude, loc.Longitude)
    }

    return str
}

func ParseGPGGA(sentenceNMEA string) (*GGA, error) {
    tokens := strings.Split(sentenceNMEA, ",")
    if len(tokens) >= 8 {
        return &GGA{
            fixTimestamp:       tokens[1],
            latitude:           tokens[2],
            latitudeDirection:  tokens[3],
            longitude:          tokens[4],
            longitudeDirection: tokens[5],
            fixQuality:         tokens[6],
            satellites:         tokens[7],
        }, nil
    }

    return nil, errors.New("malformed NMEA sentence")
}
