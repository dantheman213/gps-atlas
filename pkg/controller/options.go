package controller

import (
	"errors"
	"flag"
	"fmt"
	"runtime"
)

type Options struct {
	AutoDetect *bool
	BaudRate *int
	CalculateBearing *bool
	CalculateSignalStrength *bool
	CalculateSpeedInKPH *bool
	CalculateSpeedInMPH *bool
	Daemon *bool
	Help *bool
	PlotInterval *int
	PrintGPSCoordsToCLI *bool
	PrintNMEAToCLI *bool
	SerialPort *int
	Silent *bool
	Timeout *int
	TimezoneLocal *bool
	Verbose *bool
	WriteCSVFilePath *string
	WriteGPSCoordsFilePath *string
	WriteKMLFilePath *string
	WriteNMEAFilePath *string
}

var opts *Options

func GetOptions() *Options {
	return opts
}

func ParseOptions() {
	o := Options{}

	o.AutoDetect = flag.Bool("autodetect", true, "Auto detect the serial port and baud rate for the connected GPS device. Disabled if baud rate or port is manually set.")
	o.BaudRate = flag.Int("baudrate", -1, "Set the baud rate for the serial port.")
	o.CalculateBearing = flag.Bool("calculate-bearing", false, "Calculate directional bearing based on GPS position signals and print or write data, pair with other options.")
	o.CalculateSignalStrength = flag.Bool("calculate-signal", false, "Calculate signal strength and accuracy; print or write data, pair with other options.")
	o.CalculateSpeedInKPH = flag.Bool("calculate-speed-kph", false, "Calculate speed in kilometers per hour and print or write data, pair with other options.")
	o.CalculateSpeedInMPH = flag.Bool("calculate-speed-mph", false, "Calculate speed in miles per hour and print or write data, pair with other options.")
	o.Daemon = flag.Bool("daemon", false, "Run as a background task.")
	o.Help = flag.Bool ("help", false, "Print help sheet.")
	o.PlotInterval = flag.Int("interval", 30, "Set the plot interval (seconds) for returning a GPS location from device.")
	o.PrintGPSCoordsToCLI = flag.Bool("print-gps", false, "Print the GPS coordinates to standard out.")
	o.PrintNMEAToCLI = flag.Bool("print-nmea", false, "Print NMEA messages to standard out.")
	o.SerialPort = flag.Int("port", -1, "Set the serial port to connect.")
	o.Silent = flag.Bool("silent", false, "No output will be sent to standard out. Cannot be used with flags that write to standard out.")
	o.Timeout = flag.Int("timeout", 60, "Set the timeout (seconds) before disconnecting on error or inactivity.")
	o.TimezoneLocal = flag.Bool("timezone-local", false, "Use local timezone instead of default UTC.")
	o.Verbose = flag.Bool("verbose", false, "Extra information provided in standard out.")
	o.WriteCSVFilePath = flag.String("write-csv", "", "Write timestamp, GPS coordinates, and NMEA message(s) for location to CSV file at path provided.")
	o.WriteGPSCoordsFilePath = flag.String("write-gps", "", "Write raw GPS coordinates to file at path provided.")
	o.WriteKMLFilePath = flag.String("write-kml", "", "Write Google Maps / Earth KML format as a waypoint workflow to file at path provided.")
	o.WriteNMEAFilePath = flag.String("write-nmea", "", "Write raw NMEA messages to file at path provided.")

	flag.Parse()
	opts = &o
}

func PrintHelpSheet() {
	fmt.Println("GPS Atlas / gps-usb-serial-reader")
	fmt.Println("Auto-detect, plot, and map with common GPS USB serial devices")
	fmt.Print("\nARGUMENTS:\n\n")
	flag.PrintDefaults()
}

func SanitizeOptions(o *Options) {
	// Disable auto-detect if either baud rate or port is set
	if *o.AutoDetect && (*o.BaudRate > -1 || *o.SerialPort > -1) {
		*o.AutoDetect = false
	}

	// Set a default option if no option is set
	if !*o.PrintGPSCoordsToCLI && !*o.PrintNMEAToCLI && *o.WriteCSVFilePath == "" && *o.WriteGPSCoordsFilePath == "" && *o.WriteKMLFilePath == "" && *o.WriteNMEAFilePath == "" {
		fmt.Println("Warning: no print or write option has been set; printing GPS coordinates")
		*o.PrintGPSCoordsToCLI = true
	}
}

func ValidateOptions(o *Options) error {
	if runtime.GOOS == "windows" && !*o.AutoDetect && (*o.SerialPort <= 0 || *o.SerialPort > 256) {
		return errors.New("COM serial ports should be between 0-255")
	}

	if !*o.AutoDetect && *o.SerialPort < 0 {
		return errors.New("serial port must be valid")
	}

	if !*o.AutoDetect && *o.BaudRate <= 0 {
		return errors.New("baud rate must be valid")
	}

	if *o.Silent && *o.Verbose {
		return errors.New("cannot have both silent and verbose flags cannot both be set")
	}

	if *o.Silent && (*o.PrintNMEAToCLI || *o.PrintGPSCoordsToCLI) {
		return errors.New("can't be silent and paired with a flag that increases verbosity")
	}

	if *o.Timeout < 0 {
		return errors.New("timeout cannot be negative")
	}

	if *o.PlotInterval <= 0 {
		return errors.New("plot interval cannot be less than 0")
	}

	return nil
}