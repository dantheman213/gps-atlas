package main

import (
	"flag"
)

type Config struct {
	AutoDetect *bool
	BaudRate *int
	Help *bool
	PrintGPSCoordsToCLI *bool
	SerialPort *int
	Timeout *int
}

func parseFlags() *Config {
	c := Config{}

	c.AutoDetect = flag.Bool("autodetect", true, "Auto detect the serial port for the GPS device")
	c.BaudRate = flag.Int("baudrate", 9600, "Manually set the baud rate for the serial port")
	c.Help = flag.Bool ("help", false, "Print help sheet")
	c.PrintGPSCoordsToCLI = flag.Bool("print", false, "Print the GPS coordinates")
	c.SerialPort = flag.Int("port", -1, "Manually set the serial port to connect to")
	c.Timeout = flag.Int("timeout", 30, "Manually set the timeout before disconnecting on error or inactivity")

	flag.Parse()
	return &c
}

func printHelpSheet() {
	flag.PrintDefaults()
}