# gotsx

A simple go interface for TheSkyX design to get status from camera or mount

gotsx is a simple CLI program using the JavaScript capability of Software Bisque's TheSkyX™ Professional.
It could be used to get usefull informations to take decision in imaging sequencing software.

At Multiscale Detection, we successfully use gotsx to get status of the mount and camera when managing emergency situation.
We mainly use it as an external program in our Voyager Astronomy™ dragscript sequencing software.

# usage

As far as we use cobra to code the CLI all command are self documented ;)
there are 2 major command to get informations about mount roof and camera from theSkyX

here's the help for the 'gotsx mount' command:

```
hugh⨕shula:gotsx [ master | ✔  ]
|● ./bin/gotsx mount --help
command to get mount status informations

Usage:
  gotsx mount [flags]
  gotsx mount [command]

Available Commands:
  atParkPosition    check that the AstroPhysics mount is at park1 position or parked
  getPosition       get current Alt/Az position of the mount
  isConnected       verify that mount is connected in theSkyX
  isParked          verify that the mount is parked
  isStopped         verify that the mount is stopped
  isTracking        verify that the mount is in sideral tracking
  parkAndDisconnect should diconnect the mount from theSkyX (removed implementation)

Flags:
  -h, --help   help for mount

Global Flags:
      --config string   config file (default is conf/gotsx.yaml)

Use "gotsx mount [command] --help" for more information about a command.
```

here's the help for the 'gotsx camera' command:

```
hugh⨕shula:gotsx [ master | ✔  ]
|● ./bin/gotsx camera --help
command to get camera status informations

Usage:
  gotsx camera [flags]
  gotsx camera [command]

Available Commands:
  getCoolerPower         get the cooler power in percentage
  getCoolerStatus        get information from the camera cooler
  getTemperature         current temperature of the ccd camera
  getTemperatureSetPoint target temperature for the camera
  isCoolerOff            verify if camera warmup end and regulation is off

Flags:
  -h, --help   help for camera

Global Flags:
      --config string   config file (default is conf/gotsx.yaml)

Use "gotsx camera [command] --help" for more information about a command.
```


