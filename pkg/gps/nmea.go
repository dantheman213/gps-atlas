package gps

import (
    "errors"
    "math"
    "strconv"
    "strings"
)

type GPGGA struct {
    fixTimestamp       string
    latitude           string
    latitudeDirection  string
    longitude          string
    longitudeDirection string
    fixQuality         string
    satellites         string
}

func ParseGPGGA(s string) (*GPGGA, error) {
    tokens := strings.Split(s, ",")
    if len(tokens) >= 8 {
        return &GPGGA{
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

func (g *GPGGA) GetLatitude() (float32, error) {
    return g.parseDegrees(g.latitude, g.latitudeDirection)
}

func (g *GPGGA) GetLongitude() (float32, error) {
    return g.parseDegrees(g.longitude, g.longitudeDirection)
}

func (g *GPGGA) parseDegrees(value string, direction string) (float32, error) {
    if value == "" || direction == "" {
        return 0, errors.New("the location and / or direction value does not exist")
    }

    lat, _ := strconv.ParseFloat(value, 64)
    degrees := math.Floor(lat / 100)
    minutes := ((lat / 100) - math.Floor(lat/100)) * 100 / 60
    decimal := degrees + minutes

    if direction == "W" || direction == "S" {
        decimal *= -1
    }

    return float32(decimal), nil
}
