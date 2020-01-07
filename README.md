# gps-atlas

Auto-detect many popular GPS devices with no external drivers or setup necessary. Plot and map locations with ease. View or export GPS data in a variety of file formats that can be imported into popular applications such as Google Earth.

## Features

* Read NMEA data (United States' GPS, Russia's GLONASS, or Europe's GNSS)
* Calculate signal strength and quality
* Calculate bearing and speed (KPH, MPH, knots)
* Receive latitude and longitude in decimal degrees (DD)
* Receive satellite positional data
* Automatically create waypoint data (KML)

...and more!

This is a console application and is available to download as a binary for your operating system.

## Supported Platforms

Binaries will be provided for:

* Linux (x86 & ARM)
  - Raspberry Pi
* MacOS
* Windows

Should run on just about any modern desktop or server OS. If your OS is not found here submit an issue and I'll see what I can do. You always have the option to compile it yourself, as well.

## Supported GPS Devices

Here are some devices that are supported with no configuration at all. More devices are supported than this but these are some that I tested personally or read specifications and shoud be guaranteed to perform. Any device using NMEA protocol and connects through serial (via physical or virtual interface like USB) should work just fine.

Feel free to submit a pull request to add additional confirmed working devices.

### USB Dongles

* BS-708
* VK-162
* VK-172

### Modules

* u-blox 6
* u-blox 7

## Getting Started

### Program Options

```
Usage: gps-atlas [OPTIONS]...

OPTIONS:

  -autodetect
        Auto detect the serial port and baud rate for the connected GPS device. Disabled if baud rate or port is manually set. (default true)
  -baudrate int
        Set the baud rate for the serial port. (default -1)
  -calculate-bearing
        Calculate directional bearing based on GPS position signals and print or write data, pair with other options.
  -calculate-signal
        Calculate signal strength and accuracy; print or write data, pair with other options.
  -calculate-speed-knots
        Calculate speed in knots (nautical miles per hour) and print or write data, pair with other options.
  -calculate-speed-kph
        Calculate speed in kilometers per hour and print or write data, pair with other options.
  -calculate-speed-mph
        Calculate speed in miles per hour and print or write data, pair with other options.
  -daemon
        Run as a background task.
  -help
        Print help sheet.
  -interval int
        Set the plot interval (seconds) for returning a GPS location from device. (default 10)
  -port int
        Set the serial port to connect. (default -1)
  -print-gps
        Print GPS coordinates to stdout.
  -print-gps-extra
        Print additional GPS info to stdout.
  -print-nmea
        Print NMEA messages to stdout.
  -silent
        No output will be sent to stdout. Cannot be used with flags that write to stdout.
  -timeout int
        Set the timeout (seconds) before disconnecting on error or inactivity. (default 60)
  -timezone-local
        Use local timezone instead of default UTC.
  -verbose
        Extra information provided in stdout.
  -version
        Get the application version.
  -write-csv string
        Write timestamp, GPS coordinates, NMEA message(s), and more to CSV file at path provided.
  -write-gps string
        Write raw GPS coordinates to file at path provided.
  -write-json string
        Write timestamp, GPS coordinates, NMEA message(s), and more to JSON file at path provided.
  -write-kml string
        Write Google Maps / Earth KML format as a waypoint workflow to file at path provided.
  -write-nmea string
        Write raw NMEA messages to file at path provided.
  -write-xml string
        Write timestamp, GPS coordinates, NMEA message(s), and more to XML file at path provided.
```

## Development

TODO

### Prerequisites

#### 1. Golang 1.13

Install Golang by downloading the installer from their website.

#### 2. Make

##### MacOS

The best way is to install the Apple Developer Tools.

##### Windows

It is recommended to install `choco` utility and run the following from command prompt:

    choco install make

### Building Binary

Download dependencies by running:

    make deps

Now you can build the binary.

#### Linux

    make build-linux

#### MacOS

    make build-macos

#### Windows

    make build-windows

## Libraries

### GPS

Cross-platform GPS NMEA library created to interpret data for this project:

https://github.com/dantheman213/gps

### Serial

Cross-platform serial library created to assist with connecting to serial ports for this project:

https://github.com/dantheman213/serial

## Reference

* https://www.gpsinformation.org/dale/nmea.htm

* https://www.u-blox.com/sites/default/files/products/documents/u-blox7-V14_ReceiverDescriptionProtocolSpec_%28GPS.G7-SW-12001%29_Public.pdf

* http://www.scarpaz.com/Attic/Documents/GPS-NMEA.pdf

* https://www.sparkfun.com/datasheets/GPS/NMEA%20Reference%20Manual1.pdf

* https://m2msupport.net/m2msupport/tutorial-for-nmea-gps/

* https://brandidowns.com/2018/08/05/decoding-nmea-sentences/

* https://stackoverflow.com/questions/3159260/gps-signal-strength-calculation

* https://stackoverflow.com/questions/47028071/calculating-speed-from-set-of-longitude-and-latitudes-values-obtained-in-one-min