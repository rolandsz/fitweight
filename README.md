<p align="center">

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/rolandsz/fitweight/build.yml?branch=master)](https://github.com/rolandsz/fitweight/actions/workflows/build.yml)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/rolandsz/fitweight)](https://github.com/rolandsz/fitweight/releases/latest)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/rolandsz/fitweight)](https://github.com/rolandsz/fitweight/blob/master/go.mod)
[![Go Report Card](https://goreportcard.com/badge/github.com/rolandsz/fitweight)](https://goreportcard.com/report/github.com/rolandsz/fitweight)
[![Go Reference](https://pkg.go.dev/badge/github.com/rolandsz/fitweight.svg)](https://pkg.go.dev/github.com/rolandsz/fitweight)
[![GitHub all releases](https://img.shields.io/github/downloads/rolandsz/fitweight/total?color=orange)](https://github.com/rolandsz/fitweight/releases)
[![GitHub license](https://img.shields.io/github/license/rolandsz/fitweight)](https://github.com/rolandsz/fitweight/blob/master/LICENSE)

</p>

# fitweight
Fitweight is a program to write weight measurements to .fit files based.


## Acknowledgement
This program is based on [bodycomposition](https://github.com/davidkroell/bodycomposition).

## Download
Releases can be found in [release](https://github.com/rolandsz/fitweight/releases) tab.


## Usage

Creating a .fit file containing a weight measurement which can be uploaded to Garmin connect.
```bash
$ ./fitweight create --weight 80 --bone 14 --fat 13 --hydration 58 --muscle 42 --output monday.fit
```

General usage
```bash
$ ./fitweight -h
Fitweight is a program to write weight measurements to .fit files.

Version v1.0.0

Usage:
  fitweight [command]

Available Commands:
  help        Help about any command
  write       Writes the specified weight measurement to a .fit file

Flags:
  -h, --help   help for fitweight

Use "fitweight [command] --help" for more information about a command.
```

#### Write command usage

```bash
$ ./fitweight write -h
Writes the specified weight measurement to a .fit file

Usage:
  fitweight write [flags]

Aliases:
  write, w

Flags:
      --bmi float               Set your BMI - body mass index
  -b, --bone float              Set your bone mass in percent (default -1)
      --bone-mass float         Set your bone mass in kilograms (default -1)
  -c, --calories float          Set your caloric intake
  -f, --fat float               Set your fat in percent
  -h, --help                    help for write
      --hydration float         Set your hydration in percent
      --metabolic-age float     Set your metabolic age
  -m, --muscle float            Set your muscle mass in percent (default -1)
      --muscle-mass float       Set your muscle mass in kilograms (default -1)
  -o, --output string           Path to the output .fit file (default "default.fit")
      --physique-rating float   Set your physique rating (valid values: 1-9)
  -t, --unix-timestamp int      Set the timestamp of the measurement (default -1)
      --visceral-fat float      Set your visceral fat rating (valid values: 1-60)
  -w, --weight float            Set your weight in kilograms (default -1)
```
