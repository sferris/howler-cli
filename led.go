package main

import (
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

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  switch strings.ToLower(led.Scope) {
    case "default":
      err = led.SetDefault();

    default:
      err = led.Set();
  }

  return err
}
