# gps-atlas

Auto-detect many popular GPS devices with no external drivers or setup necessary. Plot and map locations with ease. View or export GPS data in a variety of file formats that can be imported into popular applications such as Google Earth.

## Features

* Read NMEA data (United States' GPS, Russia's GLONASS, or Europe's GNSS)
* Calculate signal strength and quality
* Calculate bearing and speed
* Receive latitude and longitude in decimal degrees (DD)
* Receive satellite positional data
* Automatically create waypoint data (KML)

...and more!

This is a console application and is available to download as a binary for your operating system.

## Supported Devices

Here are some devices that are supported with no configuration at all. More devices are supported than this but these are some that I tested personally or read specifications and shoud be guaranteed to perform. Any device using NMEA protocol and connects through serial (via physical or virtual interface like USB) should work just fine.

Feel free to submit a pull request to add additional confirmed working devices.

### Devices

* BS-708
* VK-162
* VK-172

### Modules

* u-blox 6
* u-blox 7

## Getting Started

TODO

## Reference

* https://www.gpsinformation.org/dale/nmea.htm

* https://www.u-blox.com/sites/default/files/products/documents/u-blox7-V14_ReceiverDescriptionProtocolSpec_%28GPS.G7-SW-12001%29_Public.pdf

* https://m2msupport.net/m2msupport/tutorial-for-nmea-gps/

* https://www.sparkfun.com/datasheets/GPS/NMEA%20Reference%20Manual1.pdf

* https://brandidowns.com/2018/08/05/decoding-nmea-sentences/

* https://stackoverflow.com/questions/3159260/gps-signal-strength-calculation
