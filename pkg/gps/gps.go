package gps

import (
    "fmt"
    "github.com/dantheman213/gps-atlas/pkg/gps/nmea"
    "strings"
)

const (
    DirectionNorth     = "N"
    DirectionEast      = "E"
    DirectionSouth     = "S"
    DirectionWest      = "W"
    DirectionNorthEast = "NE"
    DirectionNorthWest = "NW"
    DirectionSouthEast = "SE"
    DirectionSouthWest = "SW"
)

// Decimal Degrees
type LocationDD struct {
    Latitude  float32
    Longitude float32
}

// Degrees Minutes Seconds
type LocationDMS struct {
    Latitude           float32
    LatitudeDirection  string
    Longitude          float32
    LongitudeDirection string
}

type GPS struct {
    nmea *nmea.NMEA
}

func NewGPS() *GPS {
    r := &GPS{}
    r.nmea = &nmea.NMEA{}
    return r
}

func (g *GPS) GetGPSLocation() (*LocationDD, error) {
    lat, err := g.nmea.GGALocationFixData.GetLatitudeDD()
    if err != nil {
        return nil, err
    }

    long, err := g.nmea.GGALocationFixData.GetLongitudeDD()
    if err != nil {
        return nil, err
    }

    return &LocationDD{
        Latitude:  lat,
        Longitude: long,
    }, nil
}

func (g *GPS) GetGPSLocationPretty() string {
    loc, err := g.GetGPSLocation()
    if err != nil {
        // TODO
    }

    str := ""
    if loc != nil {
        str = fmt.Sprintf("%f, %f\n", loc.Latitude, loc.Longitude)
    }

    return str
}

func (g *GPS) ingestSatelliteNetworkType(prefix string) {
    switch prefix {
    case "GP":
        g.nmea.GPCount = g.nmea.GPCount + 1
        break
    case "GL":
        g.nmea.GLCount += 1
        break
    case "GN":
        g.nmea.GNCount += 1
        break
    default:
        // TODO
    }
}

func (g *GPS) IngestNMEASentences(sentences string) {
    s := strings.ReplaceAll(sentences, "\r", "")
    items := strings.Split(s, "\n")

    for _, item := range items {
        if nmea.IsValidNMEASentence(item) {
            g.ingestSatelliteNetworkType(item[1:3])
            nmeaCode := item[3:strings.Index(item, ",")]
            switch nmeaCode {
            case "GGA":
                d, err := nmea.ParseGGA(item)
                if err != nil {
                    // TODO
                }
                g.nmea.GGALocationFixData = d
                break
            case "RMC":
                d, err := nmea.ParseRMC(item)
                if err != nil {
                    // TODO
                }
                g.nmea.RMCRecMinData = d
                break
            case "GSA":
                d, err := nmea.ParseGSA(item)
                if err != nil {
                    // TODO
                }
                g.nmea.GSAOverallSatelliteData = d
                break
            case "GSV":
                // TODO
                break
            case "VTG":
                d, err := nmea.ParseVTG(item)
                if err != nil {
                    // TODO
                }
                g.nmea.VTGCourseAndGroundSpeed = d
                break
            default:
                // TODO ?
            }
        } else {
            // invalid NMEA sentence
            // TODO
        }
    }
}
