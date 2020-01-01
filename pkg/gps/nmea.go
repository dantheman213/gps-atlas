package gps

import (
    "errors"
    "math"
    "strconv"
)

// Support these NMEA codes at MVP
// GGA
// RMC
// GSA
// GSV

type GGA struct {
    fixTimestamp       string
    latitude           string
    latitudeDirection  string
    longitude          string
    longitudeDirection string
    fixQuality         string
    satellites         string
}

func (g *GGA) GetLatitude() (float32, error) {
    return g.parseDegrees(g.latitude, g.latitudeDirection)
}

func (g *GGA) GetLongitude() (float32, error) {
    return g.parseDegrees(g.longitude, g.longitudeDirection)
}

func (g *GGA) GetLocationPoint() (*Location, error) {
    lat, err := g.GetLatitude()
    if err != nil {
        return nil, err
    }

    long, err := g.GetLongitude()
    if err != nil {
        return nil, err
    }

    return &Location{
        Latitude:  lat,
        Longitude: long,
    }, nil
}

func (g *GGA) parseDegrees(value string, direction string) (float32, error) {
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