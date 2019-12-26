package main

import (
	"flag"
	"fmt"
)

type Options struct {
	AutoDetect *bool
	BaudRate *int
	Help *bool
	PlotInterval *int
	PrintGPSCoordsToCLI *bool
	PrintNMEAToCLI *bool
	SerialPort *int
	Timeout *int
	WriteCSVFilePath *string
	WriteGPSCoordsFilePath *string
	WriteKMLFilePath * string
	WriteNMEAFilePath *string
}

func parseFlags() *Options {
	o := Options{}

	o.AutoDetect = flag.Bool("autodetect", true, "Auto detect the serial port and baud rate for the connected GPS device. Partially or fully disabled if baud rate and/or port is manually set.")
	o.BaudRate = flag.Int("baudrate", 9600, "Set the baud rate for the serial port.")
	o.Help = flag.Bool ("help", false, "Print help sheet.")
	o.PlotInterval = flag.Int("interval", 30, "Set the plot interval (seconds) for returning a GPS location from device.")
	o.PrintGPSCoordsToCLI = flag.Bool("print-gps", false, "Print the GPS coordinates to standard out.")
	o.PrintNMEAToCLI = flag.Bool("print-nmea", false, "Print NMEA messages to standard out.")
	o.SerialPort = flag.Int("port", 0, "Set the serial port to connect.")
	o.Timeout = flag.Int("timeout", 60, "Set the timeout (seconds) before disconnecting on error or inactivity.")
	o.WriteCSVFilePath = flag.String("write-csv", "", "Write timestamp, GPS coordinates, and NMEA message(s) for location to CSV file at path provided.")
	o.WriteGPSCoordsFilePath = flag.String("write-gps", "", "Write raw GPS coordinates to file at path provided.")
	o.WriteKMLFilePath = flag.String("write-kml", "", "Write Google Maps / Earth KML format as a waypoint workflow to file at path provided.")
	o.WriteNMEAFilePath = flag.String("write-nmea", "", "Write raw NMEA messages to file at path provided.")

	flag.Parse()
	return &o
}

func printHelpSheet() {
	fmt.Println("GPS Atlas / gps-usb-serial-reader")
	fmt.Println("Auto-detect, plot, and map with common GPS USB serial devices")
	fmt.Print("\nARGUMENTS:\n\n")
	flag.PrintDefaults()
}
