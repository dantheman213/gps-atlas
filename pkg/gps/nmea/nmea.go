package nmea

import (
    "errors"
    "math"
    "strconv"
    "strings"
)

type NMEA struct {
    GGALocationFixData *GGA
    RMCRecMinData *RMC
    GSAOverallSatelliteData *GSA
    GSVDetailedSatelliteData *GSV
    VTGCourseAndGroundSpeed *VTG
}

type GGA struct {
    Timestamp       string
    Latitude           string
    LatitudeDirection  string
    Longitude          string
    LongitudeDirection string
    FixQuality         string
    Satellites         string
    Checksum string
}

type RMC struct {
    Timestamp string
    LatitudeDMS string
    LongitudeDMS string
    SpeedOverGroundInKnots string
    TrackAngleInDegrees string
    Date string
    MagneticVariation string
    Checksum string
}

type GSA struct {
    Mode1 string
    Mode2 string
    PRN string
    PDOP string
    HDOP string
    VDOP string
    Checksum string
}

// SV = Satellite Vehicle
type GSV struct {
    VisibleSVCount string
    MessageNumber string
    MessageCountInCycle string
    SVPRN string
    ElevationDegrees string
    AzimuthDegreesFromTrueNorth string
    SNR string
    Checksum string
}

type VTG struct {
    TrackMadeGoodDegreesTrue string
    TrackMadeGoodDegreesMagnetic string
    SpeedInKnots string
    SpeedOverGroundKPH string
    Checksum string
}

func (g *GGA) GetLatitudeDD() (float32, error) {
    return ParseDMSToDD(g.Latitude, g.LatitudeDirection)
}

func (g *GGA) GetLongitudeDD() (float32, error) {
    return ParseDMSToDD(g.Longitude, g.LongitudeDirection)
}

func (g *GGA) GetLatitudeDMS() (string, string, error) {
    return g.Latitude, g.LatitudeDirection, nil
}

func (g *GGA) GetLongitudeDMS() (string, string, error) {
    return g.Longitude, g.LongitudeDirection, nil
}

// Parse DMS (Degrees Minutes Seconds to Decimal Degrees)
func ParseDMSToDD(value string, direction string) (float32, error) {
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

func ParseGGA(sentenceNMEA string) (*GGA, error) {
    tokens := strings.Split(sentenceNMEA, ",")
    if len(tokens) >= 8 {
        return &GGA{
            Timestamp:          tokens[1],
            Latitude:           tokens[2],
            LatitudeDirection:  tokens[3],
            Longitude:          tokens[4],
            LongitudeDirection: tokens[5],
            FixQuality:         tokens[6],
            Satellites:         tokens[7],
        }, nil
    }

    return nil, errors.New("malformed NMEA sentence")
}