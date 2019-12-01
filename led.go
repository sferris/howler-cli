package main

import (
  "fmt"
  "strings"

  howler "github.com/sferris/howler-controller"
)

type LEDStruct struct {
  Name   string `yaml:"led"`
  Scope  string `yaml:"scope"`
  Color  string `yaml:"color"`
}

func (led LEDStruct) Process() error {
  var err error

  if len(led.Scope) <= 0 {
    led.Scope = "current"
  }

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  switch strings.ToLower(led.Scope) {
    case "current":
      err = led.Set();

    case "default":
      err = led.SetDefault();

    default:
      return fmt.Errorf("Invalid LED Scope: %s\n", led.Scope)
  }

  return err
}
