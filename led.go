package main

import (
  "os"
  "fmt"
  _"log"
  _"strings"
  _"strconv"
  "flag"

  //howler "github.com/sferris/howler-controller"
)

type Led struct {
  Button    string
  Scope     string
  RGB       rgbFlags
}

var setLedCMD *flag.FlagSet

func init() {
  setLedCMD = flag.NewFlagSet("set-led", flag.ExitOnError)
}

func parseLed() error {
  LedButton := setLedCMD.String("led", "", "Led Button to change the color (Required)")
  LedScope   := setLedCMD.String("scope", "current", "The Led scope (default or current)")

  var LedRGB rgbFlags
  setLedCMD.Var(&LedRGB, "rgb", "The RGB value for the Led color (Required)")

  setLedCMD.Parse(os.Args[2:])

  if setLedCMD.Parsed() {
    led := Led{
      Button: *LedButton,
      Scope:  *LedScope,
      RGB:     LedRGB,
    }

    led.Set();
  }

  return nil
}

func (led *Led) Set() error {
  fmt.Println(led)
  return nil
}

