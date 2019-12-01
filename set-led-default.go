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

func setDefaultLED(c *cli.Context) error {
  led := LEDStruct{
    Name:  c.String("input"),
    Scope: "default",
    Color: c.String("color"),
  }

  return led.Process()
}

func (led LEDStruct) SetDefault() error {
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

  //func (howler *HowlerDevice) SetDefaultLEDColor(led LedInputs, value string) (HowlerLed, error)
  _, err = controller.SetDefaultLEDColor(name, led.Color)
  return err
}
