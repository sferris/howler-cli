package main

import (
  "fmt"

  "gopkg.in/urfave/cli.v2"

  howler "github.com/sferris/howler-controller"
)

/*
  Name: "input",
  --
  Name: "color",
*/

func setLED(c *cli.Context) error {
  led := LEDStruct{
    Name:  c.String("input"),
    Scope: "current",
    Color: c.String("color"),
  }

  return led.Process()
}

func (led LEDStruct) Set() error {
  var err error

  name := howler.ToLed(led.Name)
  if name == -1 {
    return fmt.Errorf(
      "Invalid LED Button reference: '%s': ",
      led.Name,
    )
  }

  if controller == nil {
    controller, err = howler.OpenDevice(device)
    if err != nil {
      return err
    }
  }

  return controller.SetLEDColor(name, led.Color)
}
